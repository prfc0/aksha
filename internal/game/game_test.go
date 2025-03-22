package game

import (
	"testing"

	"github.com/prfc0/aksha/internal/card"
	"github.com/prfc0/aksha/internal/player"
)

func TestNewGame(t *testing.T) {
	players := []*player.Player{
		player.NewPlayer("1", "Alice", 1000),
		player.NewPlayer("2", "Bob", 1000),
	}
	game := NewGame(players, 10, 20)

	if len(game.Players) != 2 || game.SmallBlind != 10 || game.BigBlind != 20 {
		t.Error("NewGame did not initialize correctly")
	}
}

func TestStartHand(t *testing.T) {
	players := []*player.Player{
		player.NewPlayer("1", "Alice", 1000),
		player.NewPlayer("2", "Bob", 1000),
	}
	game := NewGame(players, 10, 20)
	game.StartHand()

	if game.Pot.Chips != 30 || len(game.CommunityCards) != 0 {
		t.Error("StartHand did not reset the game state correctly")
	}
}

func TestDealCards(t *testing.T) {
	players := []*player.Player{
		player.NewPlayer("1", "Alice", 1000),
		player.NewPlayer("2", "Bob", 1000),
	}
	game := NewGame(players, 10, 20)
	game.StartHand()

	for _, player := range game.Players {
		if len(player.Hand) != 2 {
			t.Error("DealCards did not deal two cards to each player")
		}
	}
}

func TestEndHand(t *testing.T) {
	players := []*player.Player{
		player.NewPlayer("1", "Alice", 1000),
		player.NewPlayer("2", "Bob", 1000),
	}
	game := NewGame(players, 10, 20)
	game.StartHand()
	game.EndHand()

	if game.DealerPosition != 1 || game.BettingRound != 0 {
		t.Error("EndHand did not reset the game state correctly")
	}
}

func TestDetermineWinner(t *testing.T) {
	players := []*player.Player{
		player.NewPlayer("1", "Alice", 1000),
		player.NewPlayer("2", "Bob", 1000),
	}
	game := NewGame(players, 10, 20)

	// Simulate a hand where Alice has a flush and Bob has a straight
	game.CommunityCards = []*card.Card{
		card.NewCard(card.Spades, card.Ace),
		card.NewCard(card.Spades, card.King),
		card.NewCard(card.Spades, card.Queen),
		card.NewCard(card.Spades, card.Jack),
		card.NewCard(card.Spades, card.Ten),
	}
	players[0].AddCard(card.NewCard(card.Spades, card.Nine)) // Alice has a flush
	players[1].AddCard(card.NewCard(card.Hearts, card.Nine)) // Bob has a straight

	winners := game.DetermineWinner()
	if len(winners) != 1 || winners[0].Name != "Alice" {
		t.Error("Expected Alice to win with a flush")
	}
}
