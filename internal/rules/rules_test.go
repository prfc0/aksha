package rules

import (
	"testing"

	"github.com/prfc0/aksha/internal/action"
	"github.com/prfc0/aksha/internal/player"
)

func TestValidateAction(t *testing.T) {
	player := player.NewPlayer("1", "Alice", 1000)

	tests := []struct {
		actionType  action.ActionType
		amount      int
		currentBet  int
		expectError bool
	}{
		{action.Bet, 500, 0, false},     // Valid bet
		{action.Bet, 1500, 0, true},     // Invalid bet (not enough chips)
		{action.Bet, 200, 300, true},    // Invalid bet (less than current bet)
		{action.Call, 200, 200, false},  // Valid call
		{action.Call, 1200, 200, true},  // Invalid call (not enough chips)
		{action.Raise, 300, 200, false}, // Valid raise
		{action.Raise, 100, 200, true},  // Invalid raise (less than current bet)
		{action.Fold, 0, 200, false},    // Valid fold
	}

	for _, test := range tests {
		action := action.NewAction(test.actionType, test.amount)
		err := ValidateAction(action, player, test.currentBet)

		if test.expectError {
			if err == nil {
				t.Errorf("Expected error for action %v with amount %v and current bet %v, but got none", test.actionType, test.amount, test.currentBet)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for action %v with amount %v and current bet %v: %v", test.actionType, test.amount, test.currentBet, err)
			}
		}
	}
}

func TestDetermineNextPlayer(t *testing.T) {
	players := []*player.Player{
		player.NewPlayer("1", "Alice", 1000),
		player.NewPlayer("2", "Bob", 1000),
		player.NewPlayer("3", "Charlie", 1000),
	}

	// Alice is the current player
	currentPosition := 0

	// Bob is the next active player
	nextPosition := DetermineNextPlayer(players, currentPosition)
	if nextPosition != 1 {
		t.Errorf("Expected next player position to be 1 (Bob), got %v", nextPosition)
	}

	// Bob folds, Charlie is the next active player
	players[1].Fold()
	nextPosition = DetermineNextPlayer(players, currentPosition)
	if nextPosition != 2 {
		t.Errorf("Expected next player position to be 2 (Charlie), got %v", nextPosition)
	}

	// Charlie folds, no active players left
	players[2].Fold()
	nextPosition = DetermineNextPlayer(players, currentPosition)
	if nextPosition != 0 {
		t.Errorf("Expected no active players, got %v", nextPosition)
	}
}
