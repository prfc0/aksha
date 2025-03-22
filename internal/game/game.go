package game

import (
	"log"

	"github.com/prfc0/aksha/internal/action"
	"github.com/prfc0/aksha/internal/card"
	"github.com/prfc0/aksha/internal/deck"
	"github.com/prfc0/aksha/internal/hand"
	"github.com/prfc0/aksha/internal/player"
	"github.com/prfc0/aksha/internal/pot"
)

// Game represents a single hand of Texas Hold'em poker.
type Game struct {
	Players        []*player.Player // List of players in the game
	Deck           *deck.Deck       // Deck of cards
	Pot            *pot.Pot         // Total chips in the pot
	CommunityCards []*card.Card     // Community cards on the table
	CurrentBet     int              // Current bet amount
	DealerPosition int              // Position of the dealer button
	SmallBlind     int              // Small blind amount
	BigBlind       int              // Big blind amount
	BettingRound   int              // Current betting round (0: pre-flop, 1: flop, 2: turn, 3: river)
}

// NewGame initializes a new game with the given players and blinds.
func NewGame(players []*player.Player, smallBlind, bigBlind int) *Game {
	deck := deck.NewDeck()
	deck.Shuffle()

	return &Game{
		Players:        players,
		Deck:           deck,
		Pot:            pot.NewPot(),
		CommunityCards: make([]*card.Card, 0),
		CurrentBet:     0,
		DealerPosition: len(players) - 1,
		SmallBlind:     smallBlind,
		BigBlind:       bigBlind,
		BettingRound:   0,
	}
}

// StartHand starts a new hand of poker.
func (g *Game) StartHand() {
	log.Println("Starting a new hand of Texas Hold'em.")

	// Reset player hands and status
	for _, player := range g.Players {
		player.ResetHand()
	}

	// Reset community cards and pot
	log.Println("Resetting Community Cards.")
	g.CommunityCards = make([]*card.Card, 0)
	log.Println("Resetting Pot to 0.")
	g.Pot.Chips = 0
	log.Println("Resetting Current Bet to 0.")
	g.CurrentBet = 0
}

// postBlinds posts the small and big blinds.
func (g *Game) PostBlinds() {
	smallBlindPlayer := g.Players[(g.DealerPosition+1)%len(g.Players)]
	bigBlindPlayer := g.Players[(g.DealerPosition+2)%len(g.Players)]

	smallBlindAction := action.NewAction(action.Bet, g.SmallBlind)
	smallBlindPlayer.PerformAction(smallBlindAction, g.SmallBlind)

	bigBlindAction := action.NewAction(action.Bet, g.BigBlind)
	bigBlindPlayer.PerformAction(bigBlindAction, g.BigBlind)

	g.Pot.AddChips(g.SmallBlind + g.BigBlind)
	g.CurrentBet = g.BigBlind

	log.Printf("Posted blinds: %s (small blind) and %s (big blind).\n", smallBlindPlayer.Name, bigBlindPlayer.Name)
}

// DealCards deals two cards to each player.
func (g *Game) DealCards() {
	for i := 0; i < 2; i++ {
		for _, player := range g.Players {
			card := g.Deck.Draw()
			player.AddCard(card)
		}
	}
	log.Println("Dealt cards to all players.")
}

/*
// ManageBettingRounds manages the betting rounds (pre-flop, flop, turn, river).
func (g *Game) ManageBettingRounds() {
	for g.BettingRound < 4 {
		switch g.BettingRound {
		case 0:
			log.Println("Pre-flop betting round.")
		case 1:
			log.Println("Flop betting round.")
			g.DealCommunityCards(3)
		case 2:
			log.Println("Turn betting round.")
			g.DealCommunityCards(1)
		case 3:
			log.Println("River betting round.")
			g.DealCommunityCards(1)
		}

		// Perform betting for the current round
		// g.performBettingRound()

		g.BettingRound++
	}
}
*/

// dealCommunityCards deals the specified number of community cards.
func (g *Game) DealCommunityCards(numCards int) {
	for i := 0; i < numCards; i++ {
		card := g.Deck.Draw()
		g.CommunityCards = append(g.CommunityCards, card)
		log.Printf("Dealt community card: %s\n", card.String())
	}
}

// performBettingRound performs a single betting round.
func (g *Game) PerformBettingRound() {
	// TODO: Implement betting logic (e.g., players take turns to act)
	log.Println("Performing betting round...")
	for _, player := range g.Players {
		if player.Active {
			// Simulate a player action (e.g., Call the current bet)
			action := action.NewAction(action.Call, g.CurrentBet)
			err := player.PerformAction(action, g.CurrentBet)
			g.Pot.Chips += g.CurrentBet
			if err != nil {
				log.Printf("Player %s could not perform action: %v\n", player.Name, err)
			}
		}
	}
}

// DetermineWinner determines the winner(s) of the hand.
func (g *Game) DetermineWinner() []*player.Player {
	// Evaluate each player's best 5-card hand
	bestHands := make([]*hand.Hand, 0)
	for _, player := range g.Players {
		if player.Active {
			// Combine player's hole cards with community cards
			cards := append(player.Hand, g.CommunityCards...)
			// Evaluate the best 5-card hand
			bestHand := hand.NewHand(cards)
			bestHands = append(bestHands, bestHand)
			log.Printf("Player %s has hand: %v\n", player.Name, bestHand)
		}
	}

	// Find the strongest hand(s)
	winners := make([]*player.Player, 0)
	if len(bestHands) > 0 {
		strongestHand := bestHands[0]
		winners = append(winners, g.Players[0])

		for i := 1; i < len(bestHands); i++ {
			comparison := bestHands[i].Compare(strongestHand)
			if comparison == 1 {
				// New strongest hand
				strongestHand = bestHands[i]
				winners = []*player.Player{g.Players[i]}
			} else if comparison == 0 {
				// Tie
				winners = append(winners, g.Players[i])
			}
		}
	}

	return winners
}

// EndHand ends the current hand and resets the game state.
func (g *Game) EndHand() {
	// Determine the winner(s)
	winners := g.DetermineWinner()

	// Distribute the pot to the winner(s)
	potPerWinner := g.Pot.Chips / len(winners)
	for _, winner := range winners {
		winner.Stack += potPerWinner
		log.Printf("Player %s wins %d chips!\n", winner.Name, potPerWinner)
	}

	// Reset game state for the next hand
	g.DealerPosition = (g.DealerPosition + 1) % len(g.Players)
	g.BettingRound = 0
	log.Println("Hand ended. Ready for the next hand.")
}
