package gogochat

import (
	"errors"
	"time"
)

type Message struct {
	Id int `json:"id"`
	ChannelId int `json:"channelId"`
	AuthorId int `json:"authorId"`

	CreatedAt time.Time `json:"createdAt"`
	Content string `json:"content"`
}

func (m *Message) Validate() error {
	if m.Content == "" {
		return errors.New("content is required")
	}
	return nil
}

type MessageService interface {
	CreateMessage(message *Message) error
	ListMessages(filter MessageFilter) ([]*Message, error)
}

type MessageFilter struct {
	Id        *int       `json:"id"`
	ChannelId *int       `json:"channelId"`
	From      *time.Time `json:"from"`
}
