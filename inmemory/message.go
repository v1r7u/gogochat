package inmemory

import (
	ggc "gogochat"
)

var maxMessageId = 0

type MessageService struct {
	db      map[int]*ggc.Message
	channel ggc.ChannelService
	user    ggc.UserService
}

func NewMessageService(channel ggc.ChannelService, user ggc.UserService) *MessageService {
	return &MessageService{
		db:      make(map[int]*ggc.Message),
		channel: channel,
		user:    user,
	}
}

func (s *MessageService) CreateMessage(message *ggc.Message) error {
	if err := message.Validate(); err != nil {
		return err
	}

	// Check if channel exists
	if _, err := s.channel.GetChannel(ggc.ChannelFilter{Id: &message.ChannelId}); err != nil {
		return err
	}

	// Check if user exists
	if _, err := s.user.FindUserById(message.AuthorId); err != nil {
		return err
	}

	maxMessageId++
	message.Id = maxMessageId + 1
	s.db[maxMessageId] = message

	return nil
}

func (s *MessageService) ListMessages(filter ggc.MessageFilter) ([]*ggc.Message, error) {
	if filter.Id != nil {
		message := s.db[*filter.Id]
		return []*ggc.Message{message}, nil
	}

	messages := make([]*ggc.Message, 0)
	if filter.ChannelId != nil {
		for _, m := range s.db {
			if m.ChannelId == *filter.ChannelId {
				if filter.From == nil || (*filter.From).Before(m.CreatedAt) {
					messages = append(messages, m)
				}
			}
		}
	} else {
		for _, m := range s.db {
			if filter.From == nil || (*filter.From).Before(m.CreatedAt) {
				messages = append(messages, m)
			}
		}
	}

	return messages, nil
}
