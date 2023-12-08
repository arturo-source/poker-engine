package poker

type Player struct {
	Name string
	Hand Cards
	// HasFolded bool
}

func NewPlayer(name string) *Player {
	return &Player{
		Name: name,
		Hand: NO_CARD,
	}
}

func (p *Player) AddCard(card Cards) error {
	if p.Hand.Ones() >= MAX_CARDS_PER_HAND {
		return errMaxCardsInHand
	}
	if (p.Hand & card) != 0 {
		return errCardAlreadyInHand
	}

	p.Hand |= card
	return nil
}
