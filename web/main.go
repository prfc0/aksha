package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type GameState struct {
	Players        []PlayerState `json:"players"`
	CommunityCards []Card        `json:"communityCards"`
	Pot            int           `json:"pot"`
}

type PlayerState struct {
	Name  string `json:"name"`
	Stack int    `json:"stack"`
	Hand  []Card `json:"hand"`
}

type SuitType struct {
	LongName  string `json:"longname"`
	ShortName string `json:"shortname"`
}

type Card struct {
	Rank int      `json:"rank"`
	Suit SuitType `json:"suit"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (for development only)
	},
	HandshakeTimeout: 10 * time.Second,
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	log.Println("Game service connected")

	// Handle WebSocket connection
	for {
		// Read game state update from the game service
		var gameState GameState
		err := conn.ReadJSON(&gameState)
		if err != nil {
			log.Println("WebSocket read failed:", err)
			break
		}

		// Log the received game state
		log.Printf("Received game state: %+v\n", gameState)

		// Broadcast the game state to all connected UI clients
		broadcastGameState(gameState)
	}

	log.Println("Game service disconnected")
}

var uiClients = make(map[*websocket.Conn]bool) // Track connected UI clients
var clientMutex = sync.Mutex{}                 // Mutex to protect uiClients

func broadcastGameState(gameState GameState) {
	for client := range uiClients {
		err := client.WriteJSON(gameState)
		if err != nil {
			log.Println("WebSocket write failed:", err)
			client.Close()
			delete(uiClients, client)
		}
	}
}

func handleUIWebSocket(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming WebSocket connection request from UI client")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer func() {
		conn.Close()
		clientMutex.Lock()
		delete(uiClients, conn)
		clientMutex.Unlock()
		log.Println("UI client disconnected")
	}()

	// Register UI client
	clientMutex.Lock()
	uiClients[conn] = true
	clientMutex.Unlock()
	log.Println("UI client connected successfully")

	select {}
	// Handle WebSocket connection
	/*
		for {
			// Read message from UI client (if needed)
			_, _, err := conn.ReadMessage()
			if err != nil {
				log.Println("WebSocket read failed:", err)
				break
			}
		}
	*/
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Serve static files (for the web UI)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// WebSocket endpoint for game service
	http.HandleFunc("/ws", handleWebSocket)

	// WebSocket endpoint for UI clients
	http.HandleFunc("/ui", handleUIWebSocket)

	// Start the server
	log.Println("Web server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
