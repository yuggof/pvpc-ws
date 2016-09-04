package main

import (
	"./authentication"
	"./client"
	"./router"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func main() {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		id, err := authentication.AuthenticateRequest(r.FormValue("AccessToken"))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if id == -1 {
			http.Error(w, "", 401)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}

		c := client.New(id, conn)
		router.AddClientChan <- c
		go c.Run()
	})

	fmt.Println("listening on 8080")
	http.ListenAndServe(":8080", nil)
}
