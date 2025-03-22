package table

import (
	"testing"

	"github.com/prfc0/aksha/internal/player"
)

func TestAddPlayer(t *testing.T) {
	table := NewTable(6)
	player1 := player.NewPlayer("1", "Alice", 1000)
	player2 := player.NewPlayer("2", "Bob", 1000)

	err := table.AddPlayer(player1)
	if err != nil || len(table.Players) != 1 {
		t.Error("Failed to add player to the table")
	}

	err = table.AddPlayer(player2)
	if err != nil || len(table.Players) != 2 {
		t.Error("Failed to add player to the table")
	}
}

func TestRemovePlayer(t *testing.T) {
	table := NewTable(6)
	player1 := player.NewPlayer("1", "Alice", 1000)
	player2 := player.NewPlayer("2", "Bob", 1000)

	table.AddPlayer(player1)
	table.AddPlayer(player2)

	table.RemovePlayer("1")
	if len(table.Players) != 1 || table.Players[0].ID != "2" {
		t.Error("Failed to remove player from the table")
	}
}

func TestRotateDealer(t *testing.T) {
	table := NewTable(6)
	player1 := player.NewPlayer("1", "Alice", 1000)
	player2 := player.NewPlayer("2", "Bob", 1000)

	table.AddPlayer(player1)
	table.AddPlayer(player2)

	table.RotateDealer()
	if table.DealerPosition != 1 {
		t.Error("Failed to rotate dealer button")
	}

	table.RotateDealer()
	if table.DealerPosition != 0 {
		t.Error("Failed to rotate dealer button")
	}
}

func TestActivePlayers(t *testing.T) {
	table := NewTable(6)
	player1 := player.NewPlayer("1", "Alice", 1000)
	player2 := player.NewPlayer("2", "Bob", 1000)

	table.AddPlayer(player1)
	table.AddPlayer(player2)

	player2.Fold()

	activePlayers := table.ActivePlayers()
	if len(activePlayers) != 1 || activePlayers[0].ID != "1" {
		t.Error("Failed to get active players")
	}
}
