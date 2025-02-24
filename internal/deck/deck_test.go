package deck

import (
	"testing"

	"github.com/prfc0/aksha/internal/card"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	if len(deck.Cards) != 52 {
		t.Errorf("Expected 52 cards, got %d", len(deck.Cards))
	}
}

func TestShuffle(t *testing.T) {
	deck := NewDeck()
	originalOrder := make([]*card.Card, len(deck.Cards))
	copy(originalOrder, deck.Cards)

	deck.Shuffle()

	sameOrder := true
	for i, card := range deck.Cards {
		if card != originalOrder[i] {
			sameOrder = false
			break
		}
	}
	if sameOrder {
		t.Error("Deck was not shuffled.")
	}
}

func TestDraw(t *testing.T) {
	deck := NewDeck()
	card := deck.Draw()
	if card == nil {
		t.Error("Expected a card, got nil")
	}
	if len(deck.Cards) != 51 {
		t.Errorf("Expected 51 cards, got %d", len(deck.Cards))
	}
}

func TestReset(t *testing.T) {
	deck := NewDeck()
	deck.Draw()
	deck.Reset()
	if len(deck.Cards) != 52 {
		t.Errorf("Expected 52 cards, got %d", len(deck.Cards))
	}
}
