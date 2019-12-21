package ws

import "log"

// Server -
type Server struct {
	clients      map[*client]struct{}
	connected    chan *client
	disconnected chan *client
	broadcast    chan []byte
}

// NewServer -
func NewServer() *Server {
	s := &Server{
		clients:      make(map[*client]struct{}),
		connected:    make(chan *client),
		disconnected: make(chan *client),
		broadcast:    make(chan []byte),
	}
	return s
}

// Run -
func (s *Server) Run() {
	for {
		select {
		case client := <-s.connected:
			log.Println("A client connect.")
			s.clients[client] = struct{}{}
		case client := <-s.disconnected:
			log.Println("A client disconnect.")
			delete(s.clients, client)
		case message := <-s.broadcast:
			for client := range s.clients {
				client.send <- message
			}
		}
	}
}
