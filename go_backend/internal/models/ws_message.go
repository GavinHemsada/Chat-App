package models

type WSMessage struct {
	Type    string      `json:"type"`    // "message", "join", "leave", "typing"
	RoomID  string      `json:"room_id"`
	UserID  string      `json:"user_id,omitempty"`
	Content string      `json:"content,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

type WSMessageResponse struct {
	Type      string    `json:"type"`
	Message   *Message  `json:"message,omitempty"`
	UserID    string    `json:"user_id,omitempty"`
	Username  string    `json:"username,omitempty"`
	RoomID    string    `json:"room_id,omitempty"`
	Timestamp string    `json:"timestamp,omitempty"`
}
