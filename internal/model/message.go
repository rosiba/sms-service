package model

type Message struct {
	MessageID string
	Status    string
	Recipient string
	Content   string

	Model
}
