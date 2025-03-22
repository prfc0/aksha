package hand

import (
	"testing"

	"github.com/prfc0/aksha/internal/card"
)

func TestHandEvaluation(t *testing.T) {
	tests := []struct {
		cards    []*card.Card
		expected HandRank
	}{
		{
			cards: []*card.Card{
				card.NewCard(card.Spades, card.Ace),
				card.NewCard(card.Spades, card.King),
				card.NewCard(card.Spades, card.Queen),
				card.NewCard(card.Spades, card.Jack),
				card.NewCard(card.Spades, card.Ten),
			},
			expected: RoyalFlush,
		},
		{
			cards: []*card.Card{
				card.NewCard(card.Spades, card.Nine),
				card.NewCard(card.Spades, card.Eight),
				card.NewCard(card.Spades, card.Seven),
				card.NewCard(card.Spades, card.Six),
				card.NewCard(card.Spades, card.Five),
			},
			expected: StraightFlush,
		},
		{
			cards: []*card.Card{
				card.NewCard(card.Spades, card.Ace),
				card.NewCard(card.Hearts, card.Ace),
				card.NewCard(card.Diamonds, card.Ace),
				card.NewCard(card.Clubs, card.Ace),
				card.NewCard(card.Spades, card.King),
			},
			expected: FourOfAKind,
		},
		{
			cards: []*card.Card{
				card.NewCard(card.Spades, card.Ace),
				card.NewCard(card.Hearts, card.Ace),
				card.NewCard(card.Diamonds, card.Ace),
				card.NewCard(card.Clubs, card.King),
				card.NewCard(card.Spades, card.King),
			},
			expected: FullHouse,
		},
		{
			cards: []*card.Card{
				card.NewCard(card.Spades, card.Ace),
				card.NewCard(card.Spades, card.King),
				card.NewCard(card.Spades, card.Queen),
				card.NewCard(card.Spades, card.Jack),
				card.NewCard(card.Spades, card.Nine),
			},
			expected: Flush,
		},
		{
			cards: []*card.Card{
				card.NewCard(card.Spades, card.Ten),
				card.NewCard(card.Hearts, card.Nine),
				card.NewCard(card.Diamonds, card.Eight),
				card.NewCard(card.Clubs, card.Seven),
				card.NewCard(card.Spades, card.Six),
			},
			expected: Straight,
		},
		{
			cards: []*card.Card{
				card.NewCard(card.Spades, card.Ace),
				card.NewCard(card.Hearts, card.Ace),
				card.NewCard(card.Diamonds, card.Ace),
				card.NewCard(card.Clubs, card.King),
				card.NewCard(card.Spades, card.Queen),
			},
			expected: ThreeOfAKind,
		},
		{
			cards: []*card.Card{
				card.NewCard(card.Spades, card.Ace),
				card.NewCard(card.Hearts, card.Ace),
				card.NewCard(card.Diamonds, card.King),
				card.NewCard(card.Clubs, card.King),
				card.NewCard(card.Spades, card.Queen),
			},
			expected: TwoPair,
		},
		{
			cards: []*card.Card{
				card.NewCard(card.Spades, card.Ace),
				card.NewCard(card.Hearts, card.Ace),
				card.NewCard(card.Diamonds, card.King),
				card.NewCard(card.Clubs, card.Queen),
				card.NewCard(card.Spades, card.Jack),
			},
			expected: OnePair,
		},
		{
			cards: []*card.Card{
				card.NewCard(card.Spades, card.Ace),
				card.NewCard(card.Hearts, card.King),
				card.NewCard(card.Diamonds, card.Queen),
				card.NewCard(card.Clubs, card.Jack),
				card.NewCard(card.Spades, card.Nine),
			},
			expected: HighCard,
		},
	}

	for _, test := range tests {
		hand := NewHand(test.cards)
		if hand.Rank != test.expected {
			t.Errorf("Expected %v, got %v for hand: %v", test.expected, hand.Rank, hand.Cards)
		}
	}
}
