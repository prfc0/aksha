package pot

import (
	"log"

	"github.com/prfc0/aksha/internal/player"
)

// Pot represents a poker pot (main pot or side pot).
type Pot struct {
	Chips    int              // Total chips in the pot
	Eligible []*player.Player // Players eligible to win the pot
}

func NewPot() *Pot {
	return &Pot{
		Chips:    0,
		Eligible: make([]*player.Player, 0),
	}
}

func (p *Pot) AddChips(amount int) {
	p.Chips += amount
	log.Printf("Added %d chips to the pot. Total: %d\n", amount, p.Chips)
}

func (p *Pot) AddEligiblePlayer(player *player.Player) {
	p.Eligible = append(p.Eligible, player)
	log.Printf("Player %s is eligible to win the pot.\n", player.Name)
}

func (p *Pot) Distribute(winners []*player.Player) {
	if len(winners) == 0 {
		log.Println("No winners to distribute the pot.")
		return
	}

	// Split the pot equally among winners
	chipsPerWinner := p.Chips / len(winners)
	for _, winner := range winners {
		winner.Stack += chipsPerWinner
		log.Printf("Player %s wins %d chips from the pot.\n", winner.Name, chipsPerWinner)
	}

	// Reset the pot
	p.Chips = 0
	p.Eligible = make([]*player.Player, 0)
}
