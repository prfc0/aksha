package deck

import (
	"math/rand"
	"time"

	"log"

	"github.com/prfc0/aksha/internal/card"
)

type Deck struct {
	Cards []*card.Card
}

func NewDeck() *Deck {
	deck := &Deck{}
	suits := []card.Suit{
		card.Spades,
		card.Hearts,
		card.Diamonds,
		card.Clubs,
	}
	ranks := []card.Rank{
		card.Two,
		card.Three,
		card.Four,
		card.Five,
		card.Six,
		card.Seven,
		card.Eight,
		card.Nine,
		card.Ten,
		card.Jack,
		card.Queen,
		card.King,
		card.Ace,
	}

	for _, suit := range suits {
		for _, rank := range ranks {
			deck.Cards = append(deck.Cards, card.NewCard(suit, rank))
		}
	}

	log.Println("Created a new deck of 52 cards.")
	return deck
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
	log.Println("Shuffled the deck.")
}

func (d *Deck) Draw() *card.Card {
	if len(d.Cards) == 0 {
		log.Println("No cards left in the deck.")
		return nil
	}
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	log.Printf("Drew a card: %s\n", card.String())
	return card
}

func (d *Deck) Reset() {
	d.Cards = NewDeck().Cards
	log.Println("Reset the deck to 52 cards.")
}
