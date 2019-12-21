package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type client struct {
	serv *Server
	conn *websocket.Conn
	send chan []byte
}

func (c *client) reader() {
	defer func() {
		c.serv.disconnected <- c
		c.conn.Close()
	}()
	for {
		_, p, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Print("client: ", string(p))
		c.serv.broadcast <- p
	}
}

func (c *client) writer() {
	defer c.conn.Close()
	for {
		select {
		case msg := <-c.send:
			if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				break
			}
		}
	}
}

// Hanlder -
func Hanlder(serv *Server, w http.ResponseWriter, r *http.Request) {
	// server side conn
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	c := &client{
		serv: serv,
		send: make(chan []byte),
		conn: conn,
	}
	c.serv.connected <- c

	go c.reader()
	go c.writer()
}
