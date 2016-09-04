package client

type Message struct {
	UserID         int    `json:"user_id"`
	ConversationID int    `json:"conversation_id"`
	Text           string `json:"text"`
}
