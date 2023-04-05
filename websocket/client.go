package websocket

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Hub     *Hub
	id      string
	socket  *websocket.Conn
	outBand chan []byte
}

func NewClient(h *Hub, s *websocket.Conn) *Client {
	return &Client{
		Hub:     h,
		socket:  s,
		outBand: make(chan []byte),
	}
}

func (c *Client) Write() {
	for {
		select {
		case msg, ok := <-c.outBand:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
			}
			c.socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

func (c *Client) close() {
	c.socket.Close()
	close(c.outBand)
}
