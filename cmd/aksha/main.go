package main

import (
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/prfc0/aksha/internal/action"
	"github.com/prfc0/aksha/internal/card"
	"github.com/prfc0/aksha/internal/game"
	"github.com/prfc0/aksha/internal/player"
	"github.com/prfc0/aksha/internal/table"
)

type GameState struct {
	// Players        []PlayerState `json:"players"`
	Players        []*player.Player `json:"players"`
	CommunityCards []*card.Card     `json:"communityCards"`
	Pot            int              `json:"pot"`
}

func sendGameState(conn *websocket.Conn, game *game.Game) {
	gameState := GameState{
		Players:        game.Players,
		CommunityCards: game.CommunityCards,
		Pot:            game.Pot.Chips,
	}

	// Broadcast game state
	err := conn.WriteJSON(gameState)
	if err != nil {
		log.Fatal("Failed to send game state:", err)
	}

	// Simulate delay (if needed)
	time.Sleep(3 * time.Second)
}

func main() {
	// Connect to the web server's WebSocket endpoint
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Failed to connect to web server:", err)
	}
	defer conn.Close()

	log.Println("Connected to web server")

	// Initialize players
	players := []*player.Player{
		player.NewPlayer("1", "Alice", 100),
		player.NewPlayer("2", "Bob", 100),
		player.NewPlayer("3", "Charlie", 100),
		player.NewPlayer("4", "Dave", 100),
		player.NewPlayer("5", "Eve", 100),
	}

	// Initialize the table
	log.Println("START: Adding players to table.")
	table := table.NewTable(5)
	for _, player := range players {
		table.AddPlayer(player)
	}
	log.Println("FINISH: Adding players to table.")
	log.Println("--------------------------------")

	// Initialize the game
	log.Println("START: Initializing first trial game.")
	game := game.NewGame(players, 1, 2)
	log.Println("FINISH: Initializing first trial game.")
	log.Println("--------------------------------")

	// Start a new hand
	log.Println("START: Initialization.")
	game.StartHand()
	log.Println("FINISH: Initialization.")
	sendGameState(conn, game)
	log.Println("--------------------------------")

	log.Println("START: Post small and big blind.")
	// Post small and big blinds
	game.PostBlinds()
	log.Println("FINISH: Post small and big blind.")
	sendGameState(conn, game)
	log.Println("--------------------------------")

	// Deal cards to players
	log.Println("START: Deal cards.")
	game.DealCards()
	sendGameState(conn, game)
	log.Println("FINISH: Deal cards.")
	log.Println("--------------------------------")

	// Pre-flop: Everyone calls the big blind
	activePlayers := table.ActivePlayers()
	activePlayers = append(activePlayers[2:], activePlayers[0:2]...)
	log.Println("START: Pre-flop betting round:")
	for _, player := range activePlayers {
		playerBet := player.Bet
		currentBet := game.CurrentBet
		toBet := currentBet - playerBet
		if toBet > 0 {
			actionObj := action.NewAction(action.Call, toBet)
			err := player.PerformAction(actionObj, game.CurrentBet)
			if err != nil {
				log.Printf("Player %s could not perform action: %v\n", player.Name, err)
			}
			log.Printf("Player %s calls %d chips.\n", player.Name, toBet)
			game.Pot.Chips += toBet
		} else {
			log.Printf("Player %s checks.\n", player.Name)
		}
	}
	log.Printf("Pot: %d\n", game.Pot.Chips)
	log.Println("FINISH: Pre-flop betting round:")
	sendGameState(conn, game)
	log.Println("--------------------------------")

	// Flop: Deal 3 community cards and everyone checks
	log.Println("START: Flop betting round:")
	game.DealCommunityCards(3)
	game.PerformBettingRound()
	log.Printf("Pot: %d\n", game.Pot.Chips)
	log.Println("FINISH: Flop betting round:")
	sendGameState(conn, game)
	log.Println("--------------------------------")

	// Turn: Deal 1 community card and everyone checks
	log.Println("START: Turn betting round:")
	game.DealCommunityCards(1)
	game.PerformBettingRound()
	log.Printf("Pot: %d\n", game.Pot.Chips)
	log.Println("FINISH: Turn betting round:")
	sendGameState(conn, game)
	log.Println("--------------------------------")

	// River: Deal 1 community card and everyone checks
	log.Println("START: River betting round:")
	game.DealCommunityCards(1)
	game.PerformBettingRound()
	log.Printf("Pot: %d\n", game.Pot.Chips)
	log.Println("FINISH: River betting round:")
	sendGameState(conn, game)
	log.Println("--------------------------------")

	// Reveal everyone's cards
	log.Println("START: Revealing hands:")
	for _, player := range table.ActivePlayers() {
		log.Printf("Player %s has: %v\n", player.Name, player.Hand)
	}
	log.Println("FINISH: Revealing hands:")
	sendGameState(conn, game)
	log.Println("--------------------------------")

	// Determine the winner
	log.Println("START: Determining Winners:")
	winners := game.DetermineWinner()
	if len(winners) > 0 {
		log.Printf("Player %s wins the pot of %d chips!\n", winners[0].Name, game.Pot.Chips)
	} else {
		log.Println("No winners.")
	}
	log.Println("FINISH: Determining Winners:")
	sendGameState(conn, game)
	log.Println("--------------------------------")

	// Distribute the pot to the winner(s)
	game.Pot.Distribute(winners)

	// Display everyone's stack
	log.Println("Final stacks:")
	for _, player := range table.Players {
		log.Printf("Player %s has %d chips.\n", player.Name, player.Stack)
	}
	sendGameState(conn, game)
}
