package player

import (
	"fmt"
	"log"

	"github.com/prfc0/aksha/internal/action"
	"github.com/prfc0/aksha/internal/card"
)

// Player represents a poker player.
type Player struct {
	ID     string       // Unique identifier for the player
	Name   string       // Name of the player
	Stack  int          // Chip stack
	Hand   []*card.Card // Player's hand of cards
	Bet    int          // Amount he has already bet
	Active bool         // Whether the player is still in the current hand
}

func NewPlayer(id, name string, stack int) *Player {
	return &Player{
		ID:     id,
		Name:   name,
		Stack:  stack,
		Hand:   make([]*card.Card, 0),
		Active: true,
	}
}

func (p *Player) AddCard(card *card.Card) {
	p.Hand = append(p.Hand, card)
	log.Printf("Player %s received a card: %s\n", p.Name, card.String())
}

// PerformAction performs an action (e.g., Bet, Call, Raise, Fold).
func (p *Player) PerformAction(actionObj *action.Action, currentBet int) error {
	newStack, _, err := actionObj.Execute(p.Stack, currentBet)
	if err != nil {
		return err
	}

	p.Bet += currentBet
	p.Stack = newStack
	return nil
}

func (p *Player) Fold() {
	p.Active = false
	log.Printf("Player %s folded.\n", p.Name)
}

func (p *Player) ResetHand() {
	p.Hand = make([]*card.Card, 0)
	p.Active = true
	log.Printf("Player %s's hand and status reset for a new round.\n", p.Name)
}

func (p *Player) String() string {
	return fmt.Sprintf("Player %s (Stack: %d, Active: %v)", p.Name, p.Stack, p.Active)
}
