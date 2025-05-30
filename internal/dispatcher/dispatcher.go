package dispatcher

import (
	"encoding/json"
	"log"

	"github.com/dorieq/notification-center/internal/websocket"
)

type Notification struct {
	UserID  string `json:"user_id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

func HandleNotification(n Notification) {
	payload, err := json.Marshal(n)
	if err != nil {
		log.Println("Failed to marshal notification:", err)
		return
	}

	websocket.HubInstance.Broadcast <- websocket.Message{
		TargetID: n.UserID,
		Data:     payload,
	}
}
