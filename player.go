package poker

type Player struct {
	Name string
	Hand Card
	// HasFolded bool
}

func NewPlayer(name string) *Player {
	return &Player{
		Name: name,
		Hand: NO_CARD,
	}
}

func (p *Player) AddCard(card Card) error {
	if p.Hand.Ones() >= MAX_CARDS_PER_HAND {
		return errMaxCardsInHand
	}
	if (p.Hand & card) != 0 {
		return errCardAlreadyInHand
	}

	p.Hand |= card
	return nil
}
