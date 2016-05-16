package ws

import (
	"fmt"
	"swebs/filehandler"

	"golang.org/x/net/websocket"
)

// Handle will take care about websocket requests
func Handle(ws *websocket.Conn) {
	defer ws.Close()
	fmt.Println("New socket connected.")
	for {
		var msg string

		// Get message from client
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			fmt.Println("A socket has been disconnected.")
			return
		}

		websocket.Message.Send(ws, filehandler.GetJSON(msg))
	}
}
