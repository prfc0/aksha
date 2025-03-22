package action

import (
	"testing"
)

func TestActionValidation(t *testing.T) {
	tests := []struct {
		actionType ActionType
		amount     int
		stack      int
		valid      bool
	}{
		{Bet, 100, 200, true},
		{Bet, 300, 200, false},
		{Call, 100, 200, true},
		{Call, 300, 200, false},
		{Raise, 100, 200, true},
		{Raise, 300, 200, false},
		{Fold, 0, 200, true},
	}

	for _, test := range tests {
		action := NewAction(test.actionType, test.amount)
		err := action.Validate(test.stack, 0)
		if (err == nil) != test.valid {
			t.Errorf("Expected valid=%v for action %v with amount %v and stack %v, got error: %v", test.valid, test.actionType, test.amount, test.stack, err)
		}
	}
}
