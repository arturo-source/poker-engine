package poker

import "math/bits"

// Cards represents a set of cards joined into an uint64
// where each bit represent one card from the deck.
// From the 0 to 12 one suit, 13 to 25 other one, etc.
type Cards uint64

// NewCard reads the string and returns the card, only if string has valid number and suit.
//
// Valid numbers are A K Q J T 9 8 7 6 5 4 3 2.
//
// Valid suits are s c h d.
//
// An example of card in string is "Ah".
func NewCard(cardStr string) Cards {
	var numCard, suitCard Cards
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

// String transforms the set of cards in `Cards` into a readable string.
func (c Cards) String() string {
	var cardsStr string

	for num := ACES; num >= TWOS; num >>= 1 {
		for suit := FIRST_SUIT; suit < ALL_CARDS; suit <<= 13 {
			card := c & num & suit
			if card != NO_CARD {
				cardsStr += NUMBER_VALUES[num] + SUIT_VALUES[suit]
			}
		}

		if c&num != NO_CARD {
			cardsStr += " "
		}
	}

	return cardsStr
}

// extractSuits returns the card separated by suit.
func (c Cards) extractSuits() (clubs, diamonds, hearts, spades Cards) {
	return c & CLUBS, c & DIAMONDS, c & HEARTS, c & SPADES
}

// mergeSuits joins all cards to first suit.
func (c Cards) mergeSuits() Cards {
	clubs, diamonds, hearts, spades := c.extractSuits()

	// clubs >>= 13 * 0
	diamonds >>= 13 * 1
	hearts >>= 13 * 2
	spades >>= 13 * 3

	return clubs | diamonds | hearts | spades
}

// expandToAllSuits takes Cards in first suit and replicates to all suits
// func (c Cards) expandToAllSuits() Cards {
// 	return c | c<<(13*1) | c<<(13*2) | c<<(13*3)
// }

// valueWithoutSuit returns the first cards of the same suit found.
// Use only if your cards are all the same suit!!
func (c Cards) valueWithoutSuit() Cards {
	for ; c != 0; c >>= 13 {
		value := c & FIRST_SUIT
		if value != 0 {
			return value
		}
	}

	return NO_CARD
}

// reduceRepeatedNumber is usefull when you have one number repeated (Pair, Three of a kind, or High Card) and you want the exact number of ones.
func (c Cards) reduceRepeatedNumber(nCardsToLeft int) Cards {
	mask := ALL_CARDS

	for c.Count() > nCardsToLeft {
		mask >>= 13
		c &= mask
	}

	return c
}

// quitStraightRepeatedNumbers is usefull when you only want one suit of each number in your straight.
func (c Cards) quitStraightRepeatedNumbers() Cards {
	var newCard = NO_CARD

	for n := ACES; n >= TWOS; n >>= 1 {
		val := c & n
		if val.Count() > 1 {
			val = val.reduceRepeatedNumber(1)
		}

		newCard |= val
	}

	return newCard
}

// getJustOneFlush is usefull when you only want one suit in your flush.
// Use only if your flush has 5 cards!!
func (c Cards) getJustOneFlush() Cards {
	const nCardsToLeft = 5
	mask := FIRST_SUIT << (13 * 3)

	for ; mask > 0; mask >>= 13 {
		result := c & mask
		if result.Count() == nCardsToLeft {
			return result
		}
	}

	return NO_CARD
}

// getFlushHighestNumbers is usefull when you have a flush and you want only the five higher.
// Use only if you are sure your flush is only one suit!!
func (c Cards) getFlushHighestNumbers() Cards {
	const nCardsToLeft = 5

	for mask := TWOS; mask <= ACES; mask <<= 1 {
		if c.Count() == nCardsToLeft {
			return c
		}
		c = c.QuitCards(mask)
	}

	return NO_CARD
}

// Count returns the number of cards in c.
func (c Cards) Count() int {
	return bits.OnesCount64(uint64(c))
}

// AddCards receives cardsToAdd, and returns the original card with cardsToAdd.
func (c Cards) AddCards(cardsToAdd Cards) Cards {
	return c | cardsToAdd
}

// QuitCards receives cardsToQuit, and returns the original card without cardsToQuit.
func (c Cards) QuitCards(cardsToQuit Cards) Cards {
	return c &^ cardsToQuit
}

// CardsArePresent returns true if any of cards passed is present.
func (c Cards) CardsArePresent(cards Cards) bool {
	card := c & cards
	return card != NO_CARD
}

// SetBit sets the bit in that position.
//
// Warning!! Use only if you know how the `Cards uint64` is built.
func (c Cards) SetBit(pos int) Cards {
	return c | (1 << pos)
}

// ClearBit clears the bit in that position.
//
// Warning!! Use only if you know how the `Cards uint64` is built.
func (c Cards) ClearBit(pos int) Cards {
	return c &^ (1 << pos)
}

// BitToggle toggles the bit in that position.
//
// Warning!! Use only if you know how the `Cards uint64` is built.
func (c Cards) BitToggle(pos int) Cards {
	return c ^ (1 << pos)
}

// HasBit returns true if there is a bit marked as 1 in that position.
//
// Warning!! Use only if you know how the `Cards uint64` is built.
func (c Cards) HasBit(pos int) bool {
	val := c & (1 << pos)
	return val != 0
}

// Split divides the set of cards in `Cards`, and returns an array with each one separated.
func (c Cards) Split() []Cards {
	cards := make([]Cards, 0, c.Count())

	for num := ACES; num >= TWOS; num >>= 1 {
		for suit := FIRST_SUIT; suit < ALL_CARDS; suit <<= 13 {
			card := c & num & suit
			if card != NO_CARD {
				cards = append(cards, card)
			}
		}
	}

	return cards
}

// JoinCards gets a set of cards and joins them into `Cards`.
func JoinCards(cards ...Cards) Cards {
	var c Cards
	for _, card := range cards {
		c |= card
	}

	return c
}

// HighCard receives a set of cards,
// if it is possible to find a HighCard combination, returns the best combination and true,
// in another case, random cards, and false.
func HighCard(cards Cards) (winningCards Cards, found bool) {
	for n := ACES; n >= TWOS; n >>= 1 {
		val := cards & n
		if val != NO_CARD {
			return val.reduceRepeatedNumber(1), true
		}
	}

	return NO_CARD, false
}

// Pair receives a set of cards,
// if it is possible to find a Pair combination, returns the best combination and true,
// in another case, random cards, and false.
func Pair(cards Cards) (winningCards Cards, found bool) {
	for n := ACES; n >= TWOS; n >>= 1 {
		winningCards = cards & n
		if winningCards.Count() >= 2 {
			return winningCards.reduceRepeatedNumber(2), true
		}
	}

	return NO_CARD, false
}

// TwoPair receives a set of cards,
// if it is possible to find a TwoPair combination, returns the best combination and true,
// in another case, random cards, and false.
func TwoPair(cards Cards) (winningCards Cards, found bool) {
	firstPair, found := Pair(cards)
	if !found {
		return NO_CARD, false
	}

	cardsWithoutFirstPair := cards.QuitCards(firstPair)
	secondPair, found := Pair(cardsWithoutFirstPair)

	return firstPair | secondPair, found
}

// ThreeOfAKind receives a set of cards,
// if it is possible to find a ThreeOfAKind combination, returns the best combination and true,
// in another case, random cards, and false.
func ThreeOfAKind(cards Cards) (winningCards Cards, found bool) {
	for n := ACES; n >= TWOS; n >>= 1 {
		winningCards = cards & n
		if winningCards.Count() >= 3 {
			return winningCards.reduceRepeatedNumber(3), true
		}
	}

	return NO_CARD, false
}

// Straight receives a set of cards,
// if it is possible to find a Straight combination, returns the best combination and true,
// in another case, random cards, and false.
func Straight(cards Cards) (winningCards Cards, found bool) {
	cardsOneSuited := cards.mergeSuits()
	maskMakesStraight := func(mask Cards) bool {
		winningCardsOneSuited := cardsOneSuited & mask
		return winningCardsOneSuited.Count() == 5
	}

	// Straight to ace -> Straight to six
	const strToSix = SIXS | FIVES | FOURS | THREES | TWOS
	const strToAce = strToSix << 8
	for mask := strToAce; mask >= strToSix; mask >>= 1 {
		if maskMakesStraight(mask) {
			winningCards = cards & mask
			return winningCards.quitStraightRepeatedNumbers(), true
		}
	}

	// Straight to five
	const strToFive = FIVES | FOURS | THREES | TWOS | ACES
	winningCards = cards & strToFive
	return winningCards.quitStraightRepeatedNumbers(), maskMakesStraight(strToFive)
}

// Flush receives a set of cards,
// if it is possible to find a Flush combination, returns the best combination and true,
// in another case, random cards, and false.
func Flush(cards Cards) (winningCards Cards, found bool) {
	var flushes []Cards
	appendFlushIfMatches := func(cards Cards) {
		if cards.Count() >= 5 {
			flushes = append(flushes, cards)
		}
	}

	clubs, diamonds, hearts, spades := cards.extractSuits()
	appendFlushIfMatches(clubs)
	appendFlushIfMatches(diamonds)
	appendFlushIfMatches(hearts)
	appendFlushIfMatches(spades)

	if len(flushes) == 0 {
		return NO_CARD, false
	}

	var index int
	var maxVal Cards = NO_CARD
	for i := range flushes {
		currVal := flushes[i].valueWithoutSuit()
		if currVal > maxVal {
			maxVal = currVal
			index = i
		}
	}

	return flushes[index].getFlushHighestNumbers(), true
}

// FullHouse receives a set of cards,
// if it is possible to find a FullHouse combination, returns the best combination and true,
// in another case, random cards, and false.
func FullHouse(cards Cards) (winningCards Cards, found bool) {
	threeOfAKind, found := ThreeOfAKind(cards)
	if !found {
		return NO_CARD, false
	}

	cardsWithoutThreeOfAKind := cards.QuitCards(threeOfAKind)
	pair, found := Pair(cardsWithoutThreeOfAKind)
	return threeOfAKind | pair, found
}

// FourOfAKind receives a set of cards,
// if it is possible to find a FourOfAKind combination, returns the best combination and true,
// in another case, random cards, and false.
func FourOfAKind(cards Cards) (winningCards Cards, found bool) {
	for n := ACES; n >= TWOS; n >>= 1 {
		winningCards = cards & n
		if winningCards.Count() >= 4 {
			return winningCards, true
		}
	}

	return NO_CARD, false
}

// StraightFlush receives a set of cards,
// if it is possible to find a StraightFlush combination, returns the best combination and true,
// in another case, random cards, and false.
func StraightFlush(cards Cards) (winningCards Cards, found bool) {
	const highestStraightFlushMask = KINGS | QUEENS | JACKS | TENS | NINES
	const lowestStraightFlushMask = SIXS | FIVES | FOURS | THREES | TWOS

	for mask := highestStraightFlushMask; mask >= lowestStraightFlushMask; mask >>= 1 {
		cardsMasked := cards & mask
		cardsMasked = cardsMasked.getJustOneFlush()
		if cardsMasked.Count() >= 5 {
			return cardsMasked, true
		}
	}

	const strToFiveMask = FIVES | FOURS | THREES | TWOS | ACES
	cardsMasked := cards & strToFiveMask
	cardsMasked = cardsMasked.getJustOneFlush()
	return cardsMasked, cardsMasked.Count() >= 5
}

// RoyalFlush receives a set of cards,
// if it is possible to find a RoyalFlush combination, returns the best combination and true,
// in another case, random cards, and false.
func RoyalFlush(cards Cards) (winningCards Cards, found bool) {
	const royalFlushMask = ACES | KINGS | QUEENS | JACKS | TENS
	royalFlush := cards & royalFlushMask
	royalFlush = royalFlush.getJustOneFlush()
	return royalFlush, royalFlush.Count() >= 5
}
