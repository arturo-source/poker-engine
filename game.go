package poker

import (
	"fmt"
	"sort"
)

func Play() {
	game := NewGame()

	game.Deck.Shuffle()

	game.DealCards()
	game.Board.NextBoardState() // Flop
	game.Board.NextBoardState() // Turn
	game.Board.NextBoardState() // River

	tableCards := JoinCards(game.Board.TableCards...)
	winners := GetWinners(tableCards, game.Players)
	if winners == nil {
		fmt.Println("It's a tie!")
	} else {
		fmt.Println("The winners:", winners)
	}
}

type Game struct {
	Players []*Player
	Board   *Board
	Deck    *Deck
}

func NewGame() *Game {
	d := NewDeck()
	b := NewBoard(d)

	return &Game{
		Players: make([]*Player, 0),
		Board:   b,
		Deck:    d,
	}
}

func (g *Game) DealCards() error {
	for i := 0; i < MAX_CARDS_PER_HAND; i++ {
		for j := range g.Players {
			card := g.Deck.GetNextCard()
			if card == NO_CARD {
				return errNoCardsInDeck
			}

			g.Players[j].AddCard(card)
		}
	}

	return nil
}

type PlayerHandValue struct {
	Player   *Player
	BestHand Cards
	HandKind HandKind
}

// GetWinners returns an array with the players with the best hand (it can be one or more than one)
func GetWinners(tableCards Cards, players []*Player) []PlayerHandValue {
	// Get best hand from each player
	handValues := make([]PlayerHandValue, 0, len(players))
	for _, player := range players {
		pBestHand, handKind := BestHand(player, tableCards)
		handValues = append(handValues, PlayerHandValue{
			Player:   player,
			BestHand: pBestHand,
			HandKind: handKind,
		})
	}

	// Sort best hands to get first the bests
	sort.Slice(handValues, func(i, j int) bool {
		return handValues[i].HandKind > handValues[j].HandKind
	})

	bestHandValues := []PlayerHandValue{
		{handValues[0].Player, handValues[0].BestHand, handValues[0].HandKind},
	}

	// Solve ties
	for _, opponent := range handValues[1:] {
		best := bestHandValues[0]
		if opponent.HandKind != best.HandKind {
			break
		}

		winner := tieBreakerFuncs[best.HandKind](best.Player, opponent.Player, best.BestHand, opponent.BestHand, tableCards)
		switch winner {
		case nil: // tie
			bestHandValues = append(bestHandValues, PlayerHandValue{opponent.Player, opponent.BestHand, opponent.HandKind})
		case opponent.Player: // new best
			bestHandValues = []PlayerHandValue{{opponent.Player, opponent.BestHand, opponent.HandKind}}
		}
	}

	return bestHandValues
}

// BestHand calculates the best combination of cards and what kind of hand it is
func BestHand(p *Player, tableCards Cards) (Cards, HandKind) {
	pCards := JoinCards(p.Hand, tableCards)

	// sort hand kinds because map[] is not sorted
	handKindsSorted := make([]HandKind, 0, len(combinationFuncs))
	for key := range combinationFuncs {
		handKindsSorted = append(handKindsSorted, key)
	}
	sort.Slice(handKindsSorted, func(i, j int) bool {
		return handKindsSorted[i] > handKindsSorted[j]
	})

	for _, handKind := range handKindsSorted {
		winningCards, found := combinationFuncs[handKind](pCards)
		if found {
			return winningCards, handKind
		}
	}

	return NO_CARD, -1
}

func tieBreakerHighCard(p1, p2 *Player, p1WinningCards, p2WinningCards Cards, tableCards Cards) *Player {
	p1WCOnesuited := p1WinningCards.mergeSuits()
	p2WCOnesuited := p2WinningCards.mergeSuits()

	for n := ACES; n >= TWOS; n >>= 1 {
		p1Cards := p1WCOnesuited & n
		p2Cards := p2WCOnesuited & n

		p1CardsCount := p1Cards.Count()
		p2CardsCount := p2Cards.Count()

		if p1CardsCount > p2CardsCount {
			return p1
		}
		if p2CardsCount > p1CardsCount {
			return p2
		}
	}

	return nil
}

func commonTieBreaker(p1, p2 *Player, p1WinningCards, p2WinningCards Cards, tableCards Cards) *Player {
	p1WCOnesuited := p1WinningCards.mergeSuits()
	p2WCOnesuited := p2WinningCards.mergeSuits()

	if p1WCOnesuited > p2WCOnesuited {
		return p1
	} else if p2WCOnesuited > p1WCOnesuited {
		return p2
	}

	// Quit winning combination from each hand
	// And decide who wins using the rest of the cards
	restOfTheCards := func(hand Cards, winningCards Cards) Cards {
		allCards := JoinCards(hand, tableCards)
		return allCards.QuitCards(winningCards)
	}

	p1RestOfTheCards := restOfTheCards(p1.Hand, p1WinningCards)
	p2RestOfTheCards := restOfTheCards(p2.Hand, p2WinningCards)

	// In Poker you only use 5 best card combination to choose the winner
	// If both have a three of a kind, both can use only the 2 best cards of their rest cards
	// If both have a flush, its definitely a tie
	validCardsN := 5 - p1WinningCards.Count()
	for n := ACES; n >= TWOS && validCardsN > 0; n >>= 1 {
		p1Cards := p1RestOfTheCards & n
		p2Cards := p2RestOfTheCards & n

		p1CardsCount := p1Cards.Count()
		p2CardsCount := p2Cards.Count()
		p1CardsCount = clamp(p1CardsCount, 0, validCardsN)
		p2CardsCount = clamp(p2CardsCount, 0, validCardsN)

		if p1CardsCount > p2CardsCount {
			return p1
		}
		if p2CardsCount > p1CardsCount {
			return p2
		}

		validCardsN -= p1CardsCount
	}

	return nil
}

func tieBreakerPair(p1, p2 *Player, p1WinningCards, p2WinningCards Cards, tableCards Cards) *Player {
	return commonTieBreaker(p1, p2, p1WinningCards, p2WinningCards, tableCards)
}

func tieBreakerTwoPair(p1, p2 *Player, p1WinningCards, p2WinningCards Cards, tableCards Cards) *Player {
	return commonTieBreaker(p1, p2, p1WinningCards, p2WinningCards, tableCards)
}

func tieBreakerThreeOfAKind(p1, p2 *Player, p1WinningCards, p2WinningCards Cards, tableCards Cards) *Player {
	return commonTieBreaker(p1, p2, p1WinningCards, p2WinningCards, tableCards)
}

func tieBreakerStraight(p1, p2 *Player, p1WinningCards, p2WinningCards Cards, tableCards Cards) *Player {
	const fiveHighMask = 0b1000000001111

	p1WCOnesuited := p1WinningCards.mergeSuits()
	p2WCOnesuited := p2WinningCards.mergeSuits()

	// Quit Ace as high card, if highest card is really five
	if (p1WCOnesuited & fiveHighMask) == p1WCOnesuited {
		p1WCOnesuited &= 0b1111
	}
	if (p2WCOnesuited & fiveHighMask) == p2WCOnesuited {
		p2WCOnesuited &= 0b1111
	}

	if p1WCOnesuited > p2WCOnesuited {
		return p1
	} else if p2WCOnesuited > p1WCOnesuited {
		return p2
	}

	return nil
}

func tieBreakerFlush(p1, p2 *Player, p1WinningCards, p2WinningCards Cards, tableCards Cards) *Player {
	return commonTieBreaker(p1, p2, p1WinningCards, p2WinningCards, tableCards)
}

func tieBreakerFullHouse(p1, p2 *Player, p1WinningCards, p2WinningCards Cards, tableCards Cards) *Player {
	p1Three, _ := ThreeOfAKind(p1WinningCards)
	p2Three, _ := ThreeOfAKind(p2WinningCards)
	p1ThreeRealValue := p1Three.mergeSuits()
	p2ThreeRealValue := p2Three.mergeSuits()
	if p1ThreeRealValue > p2ThreeRealValue {
		return p1
	}
	if p2ThreeRealValue > p1ThreeRealValue {
		return p2
	}

	p1Pair := p1WinningCards.QuitCards(p1Three)
	p2Pair := p2WinningCards.QuitCards(p2Three)
	p1PairRealValue := p1Pair.mergeSuits()
	p2PairRealValue := p2Pair.mergeSuits()
	if p1PairRealValue > p2PairRealValue {
		return p1
	}
	if p2PairRealValue > p1PairRealValue {
		return p2
	}

	return nil
}

func tieBreakerFourOfAKind(p1, p2 *Player, p1WinningCards, p2WinningCards Cards, tableCards Cards) *Player {
	return commonTieBreaker(p1, p2, p1WinningCards, p2WinningCards, tableCards)
}

func tieBreakerStraightFlush(p1, p2 *Player, p1WinningCards, p2WinningCards Cards, tableCards Cards) *Player {
	return tieBreakerStraight(p1, p2, p1WinningCards, p2WinningCards, tableCards)
}

func tieBreakerRoyalFlush(p1, p2 *Player, p1WinningCards, p2WinningCards Cards, tableCards Cards) *Player {
	// Always tie
	return nil
}
