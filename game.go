package poker

import "fmt"

func Play() {
	game := NewGame()

	game.Deck.Shuffle()

	game.DealCards()
	game.Board.NextBoardState() // Flop
	game.Board.NextBoardState() // Turn
	game.Board.NextBoardState() // River

	winners := game.GetWinners()
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

type CombinationFunc func(Card) (Card, bool)
type TieBreakerFunc func(p1, p2 *Player, p1WinningCards, p2WinningCards Card, tableCards Card) *Player

var (
	combinationFuncs = []CombinationFunc{
		RoyalFlush,
		StraightFlush,
		FourOfAKind,
		FullHouse,
		Flush,
		Straight,
		ThreeOfAKind,
		TwoPair,
		Pair,
		HighCard,
	}
	tieBreakerFuncs = []TieBreakerFunc{
		TieBreakerRoyalFlush,
		TieBreakerStraightFlush,
		TieBreakerFourOfAKind,
		TieBreakerFullHouse,
		TieBreakerFlush,
		TieBreakerStraight,
		TieBreakerThreeOfAKind,
		TieBreakerTwoPair,
		TieBreakerPair,
		TieBreakerHighCard,
	}
)

func (g *Game) GetWinners() []*Player {
	type HandValue struct {
		Player    *Player
		BestHand  Card
		FuncIndex int
	}

	var handValues []HandValue
	tableCards := JoinCards(g.Board.TableCards...)

	for _, player := range g.Players {
		pBestHand, pFuncIndex := g.BestHand(player, tableCards)
		handValues = append(handValues, HandValue{
			Player:    player,
			BestHand:  pBestHand,
			FuncIndex: pFuncIndex,
		})
	}

	var winners []*Player
	var winnerBestHand = NO_CARD
	var smallestIndex = len(tieBreakerFuncs)

	for _, opponent := range handValues {
		if opponent.FuncIndex < smallestIndex {
			winners = []*Player{opponent.Player}
			smallestIndex = opponent.FuncIndex
			winnerBestHand = opponent.BestHand
			continue
		}

		if opponent.FuncIndex == smallestIndex {
			currWinner := winners[0]
			winner := tieBreakerFuncs[smallestIndex](currWinner, opponent.Player, winnerBestHand, opponent.BestHand, tableCards)
			if winner == nil {
				winners = append(winners, opponent.Player)
			}
			if winner == opponent.Player {
				winners = []*Player{opponent.Player}
				smallestIndex = opponent.FuncIndex
				winnerBestHand = opponent.BestHand
			}
		}
	}

	fmt.Println(smallestIndex)

	return winners
}

func (g *Game) BestHand(p *Player, tableCards Card) (Card, int) {
	pCards := JoinCards(p.Hand, tableCards)

	for i := range combinationFuncs {
		winningCards, found := combinationFuncs[i](pCards)
		if found {
			return winningCards, i
		}
	}

	return NO_CARD, len(tieBreakerFuncs)
}

func TieBreakerHighCard(p1, p2 *Player, p1WinningCards, p2WinningCards Card, tableCards Card) *Player {
	p1WCOnesuited := p1WinningCards.MergeSuits()
	p2WCOnesuited := p2WinningCards.MergeSuits()

	for n := ACES; n >= TWOS; n >>= 1 {
		p1Cards := p1WCOnesuited & n
		p2Cards := p2WCOnesuited & n

		p1Ones := p1Cards.Ones()
		p2Ones := p2Cards.Ones()

		if p1Ones > p2Ones {
			return p1
		}
		if p2Ones > p1Ones {
			return p2
		}
	}

	return nil
}

func CommonTieBreaker(p1, p2 *Player, p1WinningCards, p2WinningCards Card, tableCards Card) *Player {
	p1WCOnesuited := p1WinningCards.MergeSuits()
	p2WCOnesuited := p2WinningCards.MergeSuits()

	if p1WCOnesuited > p2WCOnesuited {
		return p1
	} else if p2WCOnesuited > p1WCOnesuited {
		return p2
	}

	// Quit winning combination from each hand
	// And decide who wins using the rest of the cards
	restOfTheCards := func(hand Card, winningCards Card) Card {
		allCards := JoinCards(hand, tableCards)
		return allCards.QuitCards(winningCards)
	}

	p1RestOfTheCards := restOfTheCards(p1.Hand, p1WinningCards)
	p2RestOfTheCards := restOfTheCards(p2.Hand, p2WinningCards)

	// In Poker you only use 5 best card combination to choose the winner
	// If both have a three of a kind, both can use only the 2 best cards of their rest cards
	// If both have a flush, its definitely a tie
	validCardsN := 5 - p1WinningCards.Ones()
	for n := ACES; n >= TWOS && validCardsN > 0; n >>= 1 {
		p1Cards := p1RestOfTheCards & n
		p2Cards := p2RestOfTheCards & n

		p1Ones := p1Cards.Ones()
		p2Ones := p2Cards.Ones()

		if p1Ones > p2Ones {
			return p1
		}
		if p2Ones > p1Ones {
			return p2
		}

		validCardsN -= p1Ones
	}

	return nil
}

func TieBreakerPair(p1, p2 *Player, p1WinningCards, p2WinningCards Card, tableCards Card) *Player {
	return CommonTieBreaker(p1, p2, p1WinningCards, p2WinningCards, tableCards)
}

func TieBreakerTwoPair(p1, p2 *Player, p1WinningCards, p2WinningCards Card, tableCards Card) *Player {
	return CommonTieBreaker(p1, p2, p1WinningCards, p2WinningCards, tableCards)
}

func TieBreakerThreeOfAKind(p1, p2 *Player, p1WinningCards, p2WinningCards Card, tableCards Card) *Player {
	return CommonTieBreaker(p1, p2, p1WinningCards, p2WinningCards, tableCards)
}

func TieBreakerStraight(p1, p2 *Player, p1WinningCards, p2WinningCards Card, tableCards Card) *Player {
	const fiveHighMask = 0b1000000001111

	p1WCOnesuited := p1WinningCards.MergeSuits()
	p2WCOnesuited := p2WinningCards.MergeSuits()

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

func TieBreakerFlush(p1, p2 *Player, p1WinningCards, p2WinningCards Card, tableCards Card) *Player {
	return CommonTieBreaker(p1, p2, p1WinningCards, p2WinningCards, tableCards)
}

func TieBreakerFullHouse(p1, p2 *Player, p1WinningCards, p2WinningCards Card, tableCards Card) *Player {
	p1Three, _ := ThreeOfAKind(p1WinningCards)
	p2Three, _ := ThreeOfAKind(p2WinningCards)
	p1ThreeRealValue := p1Three.MergeSuits()
	p2ThreeRealValue := p2Three.MergeSuits()
	if p1ThreeRealValue > p2ThreeRealValue {
		return p1
	}
	if p2ThreeRealValue > p1ThreeRealValue {
		return p2
	}

	p1Pair := p1WinningCards.QuitCards(p1Three)
	p2Pair := p2WinningCards.QuitCards(p2Three)
	p1PairRealValue := p1Pair.MergeSuits()
	p2PairRealValue := p2Pair.MergeSuits()
	if p1PairRealValue > p2PairRealValue {
		return p1
	}
	if p2PairRealValue > p1PairRealValue {
		return p2
	}

	return nil
}

func TieBreakerFourOfAKind(p1, p2 *Player, p1WinningCards, p2WinningCards Card, tableCards Card) *Player {
	return CommonTieBreaker(p1, p2, p1WinningCards, p2WinningCards, tableCards)
}

func TieBreakerStraightFlush(p1, p2 *Player, p1WinningCards, p2WinningCards Card, tableCards Card) *Player {
	return TieBreakerStraight(p1, p2, p1WinningCards, p2WinningCards, tableCards)
}

func TieBreakerRoyalFlush(p1, p2 *Player, p1WinningCards, p2WinningCards Card, tableCards Card) *Player {
	// Always tie
	return nil
}
