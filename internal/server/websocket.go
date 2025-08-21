package server

import (
	"log"
	"mousephoneserver/internal/command"
	"net/http"

	"github.com/gorilla/websocket"
)

// Promotes normal HTTP request to websocket
var upgrader = websocket.Upgrader{

	// Permit any origin connection
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Handler for a new connection
func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade connection to WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	log.Println("New client connected")

	for {
		_, message, err := ws.ReadMessage()

		if err != nil {
			log.Printf("Connection error: %v", err)
			break // loop break
		}

		// Print message received

		cmd, err := command.Parse(message)
		if err != nil {
			log.Printf("Error trying to parse command: %v", err)

			// Ignore invalid command and go for the next message
			continue
		}

		log.Printf("Command received: %+v", cmd)
	}
}

// Initialize server on port 8080
func Start() {
	http.HandleFunc("/ws", handleConnections)

	log.Println("WebSocket server initialized in ws://localhost:8080/ws")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Error trying to initialize server: ", err)
	}
}
