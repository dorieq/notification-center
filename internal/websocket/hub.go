package websocket

type Hub struct {
	Clients    map[string]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Message
}

type Message struct {
	TargetID string
	Data     []byte
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[string]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client.ID] = client

		case client := <-h.Unregister:
			if _, ok := h.Clients[client.ID]; ok {
				close(client.Send)
				delete(h.Clients, client.ID)
			}

		case msg := <-h.Broadcast:
			if client, ok := h.Clients[msg.TargetID]; ok {
				client.Send <- msg.Data
			}
		}
	}
}
