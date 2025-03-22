package action

import (
	"fmt"
	"log"
)

type ActionType string

const (
	Bet   ActionType = "Bet"
	Call  ActionType = "Call"
	Raise ActionType = "Raise"
	Fold  ActionType = "Fold"
)

type Action struct {
	Type   ActionType
	Amount int
}

func NewAction(actionType ActionType, amount int) *Action {
	return &Action{
		Type:   actionType,
		Amount: amount,
	}
}

func (a *Action) Validate(playerStack, currentBet int) error {
	switch a.Type {
	case Bet, Raise:
		if a.Amount > playerStack {
			return fmt.Errorf("player does not have enough chips to %s %d", a.Type, a.Amount)
		}
	case Call:
		if a.Amount > playerStack {
			return fmt.Errorf("player does not have enough chips to call %d", a.Amount)
		}
	case Fold:
	default:
		return fmt.Errorf("invalid action type: %s", a.Type)
	}
	return nil
}

func (a *Action) Execute(playerStack, currentBet int) (newStack, newPot int, err error) {
	if err := a.Validate(playerStack, currentBet); err != nil {
		return playerStack, currentBet, err
	}

	switch a.Type {
	case Bet, Raise:
		newStack = playerStack - a.Amount
		newPot = currentBet + a.Amount
	case Call:
		newStack = playerStack - a.Amount
		newPot = currentBet + a.Amount
	case Fold:
		newStack = playerStack
		newPot = currentBet
	}

	log.Printf("Performed action: %s %d\n", a.Type, a.Amount)
	return newStack, newPot, nil
}
