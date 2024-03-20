package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter, r *http.Request) {
  c, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    log.Printf("upgrade:", err)
    return
  }
  defer c.Close()

  for {
    mt, message, err := c.ReadMessage()
    if err != nil {
      log.Printf("read:", err)
      break
    }

    log.Printf("received message: %v", message)
    err = c.WriteMessage(mt, message)
    if err != nil {
      log.Println("write:", err)
      break
    }
  }
}

// The Black Swan
// You are right

func main() {
  http.HandleFunc("/", echo)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
