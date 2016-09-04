package router

import (
	"../client"
	"../notifications"
	"fmt"
	"github.com/fatih/structs"
)

var (
	clients       map[int64]*client.Client
	AddClientChan chan *client.Client
)

func init() {
	clients = make(map[int64]*client.Client)
	AddClientChan = make(chan *client.Client)

	go func() {
		for {
			select {
			case c := <-AddClientChan:
				fmt.Println("new client connected")
				clients[c.ID] = c
				break
			case n := <-notifications.Channel:
				c, ok := clients[n.UserID]
				if ok {
					c.OutputChan <- &client.Event{"notification", structs.Map(n)}
				}
				break
			}
		}
	}()
}
