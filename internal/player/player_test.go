package player

import (
	"testing"

	"github.com/prfc0/aksha/internal/action"
	"github.com/prfc0/aksha/internal/card"
)

func TestPerformAction(t *testing.T) {
	player := NewPlayer("1", "Alice", 1000)

	tests := []struct {
		actionType    action.ActionType
		amount        int
		expectedStack int
		expectError   bool
	}{
		{action.Bet, 500, 500, false},    // Valid bet
		{action.Bet, 1500, 1000, true},   // Invalid bet (not enough chips)
		{action.Call, 200, 300, false},   // Valid call
		{action.Call, 1200, 1000, true},  // Invalid call (not enough chips)
		{action.Raise, 300, 0, false},    // Valid raise
		{action.Raise, 1500, 1000, true}, // Invalid raise (not enough chips)
		{action.Fold, 0, 0, false},       // Valid fold
	}

	for _, test := range tests {
		action := action.NewAction(test.actionType, test.amount)
		err := player.PerformAction(action, test.amount)

		if test.expectError {
			if err == nil {
				t.Errorf("Expected error for action %v with amount %v, but got none", test.actionType, test.amount)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for action %v with amount %v: %v", test.actionType, test.amount, err)
			}
			if player.Stack != test.expectedStack {
				t.Errorf("Expected stack %v after action %v, got %v", test.expectedStack, test.actionType, player.Stack)
			}
		}
	}
}

func TestFold(t *testing.T) {
	player := NewPlayer("1", "Alice", 1000)
	player.Fold()

	if player.Active {
		t.Error("Expected player to be inactive after folding")
	}
}

func TestResetHand(t *testing.T) {
	player := NewPlayer("1", "Alice", 1000)
	player.AddCard(card.NewCard(card.Spades, card.Ace))
	player.Fold()
	player.ResetHand()

	if len(player.Hand) != 0 || !player.Active {
		t.Error("ResetHand did not reset the player correctly")
	}
}
