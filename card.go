package poker

import (
	"math/bits"
)

type Card uint64

func NewCard(cardStr string) Card {
	var numCard, suitCard Card
	number, suit := string(cardStr[0]), string(cardStr[1])

	for key, val := range NUMBER_VALUES {
		if number == val {
			numCard = key
			break
		}
	}

	for key, val := range SUIT_VALUES {
		if suit == val {
			suitCard = key
			break
		}
	}

	return numCard & suitCard
}

func (c Card) String() string {
	var cardsStr string

	for nkey, nval := range NUMBER_VALUES {
		if c&nkey != NO_CARD {
			for skey, sval := range SUIT_VALUES {
				if c&nkey&skey != NO_CARD {
					cardsStr += nval + sval
				}
			}
			cardsStr += " "
		}
	}

	return cardsStr
}

// Ones returns the number of bits marked as 1 in Card.
func (c Card) Ones() int {
	return bits.OnesCount64(uint64(c))
}

// ExtractSuits returns the card separated by suit.
func (c Card) ExtractSuits() (clubs, diamonds, hearts, spades Card) {
	return c & CLUBS, c & DIAMONDS, c & HEARTS, c & SPADES
}

// AllSuitsToOneSuit joins all cards to one suit.
func (c Card) AllSuitsToOneSuit() Card {
	clubs, diamonds, hearts, spades := c.ExtractSuits()

	// clubs >>= 13 * 0
	diamonds >>= 13 * 1
	hearts >>= 13 * 2
	spades >>= 13 * 3

	return clubs | diamonds | hearts | spades
}

// OneSuitToAllSuits expands the maskOneSuited to all suits.
// Then, executes AND logic with the real card, to delete the cards that are not in the expanded mask.
func (c Card) OneSuitToAllSuits(maskOneSuited Card) Card {
	maskExpanded := maskOneSuited | maskOneSuited<<(13*1) | maskOneSuited<<(13*2) | maskOneSuited<<(13*3)
	return c & maskExpanded
}

// ValueWithoutSuit returns the first cards of the same suit found.
// Use only if c are all the same suit
func (c Card) ValueWithoutSuit() Card {
	for ; c != 0; c >>= 13 {
		value := c & ONE_SUIT
		if value != 0 {
			return value
		}
	}

	return 0
}

// QuitCards receives cardsToQuit, and return original card without cardsToQuit
func (c Card) QuitCards(cardsToQuit Card) Card {
	return c &^ cardsToQuit
}

func JoinCards(cards ...Card) Card {
	var c Card
	for _, card := range cards {
		c |= card
	}

	return c
}

func HighCard(cards Card) (winningCards Card, found bool) {
	for n := ACES; n >= TWOS; n >>= 1 {
		val := cards & n
		if val != NO_CARD {
			return val, true
		}
	}
	return NO_CARD, false
}

func Pair(cards Card) (winningCards Card, found bool) {
	for n := ACES; n >= TWOS; n >>= 1 {
		winningCards = cards & n
		if winningCards.Ones() >= 2 {
			return winningCards, true
		}
	}

	return NO_CARD, false
}

func TwoPair(cards Card) (winningCards Card, found bool) {
	firstPair, found := Pair(cards)
	if !found {
		return NO_CARD, false
	}

	cardsWithoutFirstPair := cards.QuitCards(firstPair)
	secondPair, found := Pair(cardsWithoutFirstPair)

	return firstPair | secondPair, found
}

func ThreeOfAKind(cards Card) (winningCards Card, found bool) {
	for n := ACES; n >= TWOS; n >>= 1 {
		winningCards = cards & n
		if winningCards.Ones() >= 3 {
			return winningCards, true
		}
	}

	return NO_CARD, false
}

func Straight(cards Card) (winningCards Card, found bool) {
	cardsOneSuited := cards.AllSuitsToOneSuit()
	maskMakesStraight := func(mask Card) bool {
		winningCardsOneSuited := cardsOneSuited & mask
		return winningCardsOneSuited.Ones() >= 5
	}

	// to ace Straight -> to six Straight
	const end Card = 0b11111
	const start Card = end << 8
	for mask := start; mask >= end; mask >>= 1 {
		if maskMakesStraight(mask) {
			winningCards = cards.OneSuitToAllSuits(mask)
			return winningCards, true
		}
	}

	// to five Straight
	const maskHighCardFive = ONE_SUIT & (FIVES | FOURS | THREES | TWOS | ACES)
	winningCards = cards.OneSuitToAllSuits(maskHighCardFive)
	return winningCards, maskMakesStraight(maskHighCardFive)
}

func Flush(cards Card) (winningCards Card, found bool) {
	var flushes []Card
	appendFlushIfMatches := func(cards Card) {
		if cards.Ones() >= 5 {
			flushes = append(flushes, cards)
		}
	}

	clubs, diamonds, hearts, spades := cards.ExtractSuits()
	appendFlushIfMatches(clubs)
	appendFlushIfMatches(diamonds)
	appendFlushIfMatches(hearts)
	appendFlushIfMatches(spades)

	if len(flushes) == 0 {
		return NO_CARD, false
	}

	var index int
	var maxVal Card = NO_CARD
	for i := range flushes {
		currVal := flushes[i].ValueWithoutSuit()
		if currVal > maxVal {
			currVal = maxVal
			index = i
		}
	}

	return flushes[index], true
}

func FullHouse(cards Card) (winningCards Card, found bool) {
	threeOfAKind, found := ThreeOfAKind(cards)
	if !found {
		return NO_CARD, false
	}

	cardsWithoutThreeOfAKind := cards.QuitCards(threeOfAKind)
	pair, found := Pair(cardsWithoutThreeOfAKind)
	return threeOfAKind | pair, found
}

func FourOfAKind(cards Card) (winningCards Card, found bool) {
	for n := ACES; n >= TWOS; n >>= 1 {
		winningCards = cards & n
		if winningCards.Ones() >= 4 {
			return winningCards, true
		}
	}

	return NO_CARD, false
}

func StraightFlush(cards Card) (winningCards Card, found bool) {
	const highestStraightFlushMask Card = 0b0111110000000011111000000001111100000000111110000000
	const lowestStraightFlushMask Card = 0b0000000011111000000001111100000000111110000000011111

	for mask := highestStraightFlushMask; mask >= lowestStraightFlushMask; mask >>= 1 {
		cardsMasked := cards & mask
		if cardsMasked.Ones() >= 5 {
			return cardsMasked, true
		}
	}

	return NO_CARD, false
}

func RoyalFlush(cards Card) (winningCards Card, found bool) {
	const royalFlushMask Card = 0b1111100000000111110000000011111000000001111100000000
	royalFlush := cards & royalFlushMask
	return royalFlush, royalFlush.Ones() >= 5
}
