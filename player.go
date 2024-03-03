package poker

// Player stores all the player information.
type Player struct {
	Name       string
	Hand       Cards
	Coins      uint
	BetCoins   uint
	HasFolded  bool
	HasChecked bool
}

// NewPlayer returns a player with that name.
func NewPlayer(name string) *Player {
	return &Player{
		Name: name,
		Hand: NO_CARD,
	}
}

// AddCard returns an error if the user has reached max cards in hand,
// in other case adds the card to the player hand.
func (p *Player) AddCard(card Cards) error {
	if p.Hand.Count() >= MAX_CARDS_PER_HAND {
		return errMaxCardsInHand
	}

	p.Hand |= card
	return nil
}
