package entity

type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	To      string `json:"to"`
	Status  string `json:"status"`
}

type SentMessage struct {
	MessageID string `json:"message_id"`
	SentTime  string `json:"sent_time"`
}
