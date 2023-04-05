package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

type Hub struct {
	clients    []*Client
	register   chan *Client
	unRegister chan *Client
	mutex      *sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make([]*Client, 0),
		register:   make(chan *Client),
		unRegister: make(chan *Client),
		mutex:      &sync.Mutex{},
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case c := <-hub.register:
			hub.connect(c)
		case c := <-hub.unRegister:
			hub.disconnect(c)
		}
	}
}

func (hub *Hub) BroadCast(msg interface{}, ignore *Client) {}

func (hub *Hub) connect(cl *Client) {
	log.Println("Client Connect : ", cl.socket.RemoteAddr())
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	cl.id = cl.socket.RemoteAddr().String()
	hub.clients = append(hub.clients, cl)
}

func (hub *Hub) disconnect(cl *Client) {
	log.Println("Client Disconnect: ", cl.socket.RemoteAddr())
	cl.close()
	hub.mutex.Lock()
	defer hub.mutex.Unlock()

	i := -1

	for j, c := range hub.clients {
		if c.id == cl.id {
			i = j
		}
		break
	}

	copy(hub.clients[i:], hub.clients[i+1:])
	hub.clients[len(hub.clients)-1] = nil
	hub.clients = hub.clients[:len(hub.clients)-1]
}
