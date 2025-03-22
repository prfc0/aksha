package hand

import (
	"fmt"
	"sort"

	"github.com/prfc0/aksha/internal/card"
)

type HandRank int

const (
	HighCard HandRank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

type Hand struct {
	Cards []*card.Card
	// Rank of the hand (e.g., Flush, Straight)
	Rank HandRank
	// For comparison (e.g. Straight vs Straight, Flush vs Flush)
	Strength []int
}

func NewHand(cards []*card.Card) *Hand {
	hand := &Hand{
		Cards:    cards,
		Rank:     HighCard,
		Strength: make([]int, 0),
	}
	hand.evaluate()
	return hand
}

func (h *Hand) evaluate() {
	sort.Slice(h.Cards, func(i, j int) bool {
		return h.Cards[i].Value() > h.Cards[j].Value()
	})

	isFlush := h.isFlush()
	isStraight := h.isStraight()

	if isFlush && isStraight && h.Cards[0].Rank == card.Ace {
		if h.Cards[0].Rank == card.Ace {
			h.Rank = RoyalFlush
		} else {
			h.Rank = StraightFlush
		}
		h.Strength = []int{h.Cards[0].Value()}
		return
	}

	if h.hasNOfAKind(4) {
		h.Rank = FourOfAKind
		h.Strength = h.getNOfAKindStrength(4)
		return
	}

	if h.hasFullHouse() {
		h.Rank = FullHouse
		h.Strength = h.getFullHouseStrength()
		return
	}

	if isFlush {
		h.Rank = Flush
		h.Strength = h.getHighCardStrength()
		return
	}

	if isStraight {
		h.Rank = Straight
		h.Strength = []int{h.Cards[0].Value()}
		return
	}

	if h.hasNOfAKind(3) {
		h.Rank = ThreeOfAKind
		h.Strength = h.getNOfAKindStrength(3)
		return
	}

	if h.hasTwoPair() {
		h.Rank = TwoPair
		h.Strength = h.getTwoPairStrength()
		return
	}

	if h.hasNOfAKind(2) {
		h.Rank = OnePair
		h.Strength = h.getNOfAKindStrength(2)
		return
	}

	h.Rank = HighCard
	h.Strength = h.getHighCardStrength()
}

func (h *Hand) isFlush() bool {
	suit := h.Cards[0].Suit
	for _, card := range h.Cards {
		if card.Suit != suit {
			return false
		}
	}
	return true
}

func (h *Hand) isStraight() bool {
	isRegularStraight := true
	for i := 0; i < len(h.Cards)-1; i++ {
		if h.Cards[i].Value() != h.Cards[i+1].Value()+1 {
			isRegularStraight = false
			break
		}
	}
	if isRegularStraight {
		return true
	}

	isWheelStraight := true
	wheelRanks := []card.Rank{card.Ace, card.Two, card.Three, card.Four, card.Five}
	rankMap := make(map[card.Rank]bool)
	for _, card := range h.Cards {
		rankMap[card.Rank] = true
	}

	for _, rank := range wheelRanks {
		if !rankMap[rank] {
			isWheelStraight = false
			break
		}
	}

	return isWheelStraight
}

func (h *Hand) hasNOfAKind(n int) bool {
	rankCount := make(map[card.Rank]int)
	for _, card := range h.Cards {
		rankCount[card.Rank]++
	}
	for _, count := range rankCount {
		if count == n {
			return true
		}
	}
	return false
}

func (h *Hand) getNOfAKindStrength(n int) []int {
	rankCount := make(map[card.Rank]int)
	for _, card := range h.Cards {
		rankCount[card.Rank]++
	}

	strength := make([]int, 0)
	for rank, count := range rankCount {
		if count == n {
			strength = append(strength, int(rank))
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(strength)))
	return strength
}

func (h *Hand) hasFullHouse() bool {
	return h.hasNOfAKind(3) && h.hasNOfAKind(2)
}

func (h *Hand) getFullHouseStrength() []int {
	strength := make([]int, 0)
	strength = append(strength, h.getNOfAKindStrength(3)...)
	strength = append(strength, h.getNOfAKindStrength(2)...)
	return strength
}

func (h *Hand) hasTwoPair() bool {
	pairCount := 0
	rankCount := make(map[card.Rank]int)
	for _, card := range h.Cards {
		rankCount[card.Rank]++
	}
	for _, count := range rankCount {
		if count == 2 {
			pairCount++
		}
	}
	return pairCount >= 2
}

func (h *Hand) getTwoPairStrength() []int {
	strength := make([]int, 0)
	rankCount := make(map[card.Rank]int)
	for _, card := range h.Cards {
		rankCount[card.Rank]++
	}
	for rank, count := range rankCount {
		if count == 2 {
			strength = append(strength, int(rank))
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(strength)))
	return strength
}

func (h *Hand) getHighCardStrength() []int {
	strength := make([]int, 0)
	for _, card := range h.Cards {
		strength = append(strength, card.Value())
	}
	return strength
}

// Compare compares two hands and returns:
// -1 if h is weaker than other,
// 0 if h is equal to other,
// 1 if h is stronger than other.
func (h *Hand) Compare(other *Hand) int {
	if h.Rank > other.Rank {
		return 1
	} else if h.Rank < other.Rank {
		return -1
	}

	for i := 0; i < len(h.Strength); i++ {
		if h.Strength[i] > other.Strength[i] {
			return 1
		} else if h.Strength[i] < other.Strength[i] {
			return -1
		}
	}
	return 0
}

func (h *Hand) String() string {
	return fmt.Sprintf("Hand: %v, Rank: %v, Strength: %v", h.Cards, h.Rank, h.Strength)
}
