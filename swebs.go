package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"swebs/filehandler"
	"swebs/ws"

	"golang.org/x/net/websocket"
)

func main() {

	// Retriving cmd flags
	port := flag.String("port", "9393", "Specify desiderated port [Default 9393]")
	wsActive := flag.Bool("ws", true, "Active an echo server to websockets")
	createDir := flag.Bool("checkdir", true, "It will check if resources dir exists and makes it if not.")
	flag.Parse()

	if *createDir {
		if err := filehandler.SetUp(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	// Redirecting requests to resources' folder
	http.Handle("/", http.FileServer(http.Dir("resources")))

	// Activate ws handler
	if *wsActive {
		http.Handle("/socket", websocket.Handler(ws.Handle))
	}

	// Let's get this party started
	fmt.Printf("swebs started [port %s] !\n", *port)
	http.ListenAndServe(":"+*port, nil)
}
