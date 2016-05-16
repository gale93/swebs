package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

// wsEcho will echo everything it'll get via websocket
func wsEcho(ws *websocket.Conn) {
	defer ws.Close()

	for {
		var msg interface{}

		// Get message from client
		if err := websocket.JSON.Receive(ws, &msg); err != nil {
			fmt.Println("dced: " + err.Error())
			return
		}

		fmt.Println(msg)

		// Echo it.
		websocket.JSON.Send(ws, msg)
	}
}

func main() {

	// Retriving cmd flags
	port := flag.String("port", "9393", "Specify desiderated port [Default 9393]")
	wsActive := flag.Bool("ws", true, "Active an echo server to websockets")
	createDir := flag.Bool("checkdir", true, "It will check if resources dir exists and makes it if not.")
	flag.Parse()

	if *createDir {
		if err := os.MkdirAll("resources", 0711); err != nil {
			fmt.Println("Impossible create resources dir. Consider -checkdir=false.\n[error]: " + err.Error())
			os.Exit(1)
		}
		f, err := os.Create("resources/index.html")
		if err != nil {
			fmt.Println("Impossible create resources dir. Consider -checkdir=false.\n[error]: " + err.Error())
			os.Exit(1)
		}

		f.Write([]byte("<h1>Hello simple world!</h1>"))
	}

	// Redirecting requests to resources' folder
	http.Handle("/", http.FileServer(http.Dir("resources")))

	// Activate ws handler
	if *wsActive {
		http.Handle("/socket", websocket.Handler(wsEcho))
	}

	// Let's get this party started
	fmt.Println("Simple Server Started!")
	http.ListenAndServe(":"+*port, nil)
}
