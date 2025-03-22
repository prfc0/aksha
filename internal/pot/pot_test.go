package pot

import (
	"testing"

	"github.com/prfc0/aksha/internal/player"
)

func TestAddChips(t *testing.T) {
	pot := NewPot()
	pot.AddChips(500)

	if pot.Chips != 500 {
		t.Error("Failed to add chips to the pot")
	}
}

func TestAddEligiblePlayer(t *testing.T) {
	pot := NewPot()
	player1 := player.NewPlayer("1", "Alice", 1000)

	pot.AddEligiblePlayer(player1)
	if len(pot.Eligible) != 1 || pot.Eligible[0].ID != "1" {
		t.Error("Failed to add eligible player to the pot")
	}
}

func TestDistribute(t *testing.T) {
	pot := NewPot()
	player1 := player.NewPlayer("1", "Alice", 1000)
	player2 := player.NewPlayer("2", "Bob", 1000)

	pot.AddChips(1000)
	pot.AddEligiblePlayer(player1)
	pot.AddEligiblePlayer(player2)

	winners := []*player.Player{player1}
	pot.Distribute(winners)

	if player1.Stack != 2000 || pot.Chips != 0 {
		t.Error("Failed to distribute the pot to the winner")
	}
}
