package card

import (
	"testing"
)

func TestCardComparison(t *testing.T) {
	ace := NewCard(Spades, Ace)
	king := NewCard(Hearts, King)
	jack := NewCard(Diamonds, Jack)
	two := NewCard(Clubs, Two)

	if ace.Value() <= king.Value() {
		t.Errorf("Expected Ace (14) > King (13)")
	}
	if king.Value() <= jack.Value() {
		t.Errorf("Expected King (13) > Jack (11)")
	}
	if jack.Value() <= two.Value() {
		t.Errorf("Expected Jack (11) > Two (2)")
	}
}

func TestRankToString(t *testing.T) {
	tests := []struct {
		rank     Rank
		expected string
	}{
		{Ace, "A"},
		{King, "K"},
		{Queen, "Q"},
		{Jack, "J"},
		{Ten, "T"},
		{Two, "2"},
	}

	for _, test := range tests {
		if test.rank.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.rank.String())
		}
	}
}
