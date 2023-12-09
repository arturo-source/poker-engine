package poker

import "errors"

var (
	errNoCardsInDeck     = errors.New("no more cards in deck")
	errNoCardsToFlip     = errors.New("no more cards to flip")
	errMaxCardsInHand    = errors.New("max cards added to hand")
	errCardAlreadyInHand = errors.New("card already added to hand")
)

const (
	MAX_CARDS_IN_BOARD = 5
	MAX_BURNED_CARDS   = 3
	MAX_CARDS_PER_HAND = 2
)

const (
	NO_CARD   Cards = 0
	ONE_SUIT  Cards = 0b1111111111111
	ALL_CARDS Cards = 0b1111111111111111111111111111111111111111111111111111
)

const (
	CLUBS = ONE_SUIT << (iota * 13)
	DIAMONDS
	HEARTS
	SPADES
)

const (
	TWOS Cards = 0b0000000000001000000000000100000000000010000000000001 << iota
	THREES
	FOURS
	FIVES
	SIXS
	SEVENS
	EIGHTS
	NINES
	TENS
	JACKS
	QUEENS
	KINGS
	ACES
)

var (
	TOTAL_CARDS = len(SUIT_VALUES) * len(NUMBER_VALUES)

	SUIT_VALUES = map[Cards]string{
		// DIAMONDS: "♦",
		// CLUBS:    "♣",
		// HEARTS:   "♥",
		// SPADES:   "♠",
		DIAMONDS: "d",
		CLUBS:    "c",
		HEARTS:   "h",
		SPADES:   "s",
	}

	NUMBER_VALUES = map[Cards]string{
		ACES:   "A",
		KINGS:  "K",
		QUEENS: "Q",
		JACKS:  "J",
		TENS:   "T",
		NINES:  "9",
		EIGHTS: "8",
		SEVENS: "7",
		SIXS:   "6",
		FIVES:  "5",
		FOURS:  "4",
		THREES: "3",
		TWOS:   "2",
	}
)

type combinationFunc func(Cards) (Cards, bool)
type tieBreakerFunc func(p1, p2 *Player, p1WinningCards, p2WinningCards Cards, tableCards Cards) *Player
type HandKind int

func (hk HandKind) String() string {
	names := [...]string{
		"High card",
		"Pair",
		"Two pair",
		"Three of a kind",
		"Straight",
		"Flush",
		"Full house",
		"Four of a kind",
		"Straight flush",
		"Royal flush",
	}

	if hk < HIGHCARD || hk > ROYALFLUSH {
		return "Unknown HandKind"
	}

	return names[hk]
}

const (
	HIGHCARD HandKind = iota
	PAIR
	TWOPAIR
	THREEOFAKIND
	STRAIGHT
	FLUSH
	FULLHOUSE
	FOUROFAKIND
	STRAIGHTFLUSH
	ROYALFLUSH
)

var (
	combinationFuncs = map[HandKind]combinationFunc{
		ROYALFLUSH:    RoyalFlush,
		STRAIGHTFLUSH: StraightFlush,
		FOUROFAKIND:   FourOfAKind,
		FULLHOUSE:     FullHouse,
		FLUSH:         Flush,
		STRAIGHT:      Straight,
		THREEOFAKIND:  ThreeOfAKind,
		TWOPAIR:       TwoPair,
		PAIR:          Pair,
		HIGHCARD:      HighCard,
	}
	tieBreakerFuncs = map[HandKind]tieBreakerFunc{
		ROYALFLUSH:    tieBreakerRoyalFlush,
		STRAIGHTFLUSH: tieBreakerStraightFlush,
		FOUROFAKIND:   tieBreakerFourOfAKind,
		FULLHOUSE:     tieBreakerFullHouse,
		FLUSH:         tieBreakerFlush,
		STRAIGHT:      tieBreakerStraight,
		THREEOFAKIND:  tieBreakerThreeOfAKind,
		TWOPAIR:       tieBreakerTwoPair,
		PAIR:          tieBreakerPair,
		HIGHCARD:      tieBreakerHighCard,
	}
)

func clamp(v, min, max int) int {
	if v > max {
		return max
	}
	if v < min {
		return min
	}

	return v
}
