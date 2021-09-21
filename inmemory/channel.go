package inmemory

import (
	"errors"
	"fmt"
	ggc "gogochat"
)

var maxChannelId = 0

type ChannelService struct {
	db map[int]*ggc.Channel
}

func NewChannelService() *ChannelService {
	return &ChannelService{db: make(map[int]*ggc.Channel)}
}

func (s *ChannelService) GetChannel(filter ggc.ChannelFilter) (*ggc.Channel, error) {
	if filter.Id != nil {
		id := *filter.Id

		if len(s.db) <= id {
			return nil, errors.New(fmt.Sprintf("there is no channel with id %v", filter.Id))
		}
		return s.db[id], nil
	}

	if filter.Name != "" {
		for _, c := range s.db {
			if c.Name == filter.Name {
				return c, nil
			}
		}

		return nil, errors.New(fmt.Sprintf("there is no channel with name %v", filter.Name))
	}

	return nil, errors.New("both criteria are empty")

}

func (s *ChannelService) ListChannels() ([]*ggc.Channel, error) {
	v := make([]*ggc.Channel, 0, len(s.db))

	for  _, value := range s.db {
		v = append(v, value)
	}
	return v, nil
}

func (s *ChannelService) CreateChannel(channel *ggc.Channel) error {
	if err := channel.Validate(); err != nil {
		return err
	}

	for _, c := range s.db {
		if c.Name == channel.Name {
			return errors.New(fmt.Sprintf("%v already exists", c.Name))
		}
	}

	maxChannelId++
	channel.Id = maxChannelId
	s.db[maxChannelId] = channel

	return nil
}

func (s *ChannelService) JoinUser(channel *ggc.Channel, user *ggc.User) error  {
	if channel == nil {
		return errors.New("channel is required")
	}
	if user == nil {
		return errors.New("user is required")
	}

	if _, err := s.GetChannel(ggc.ChannelFilter{Id: &channel.Id}); err != nil {
		return err
	}

	channel.Members = append(channel.Members, user)

	return nil
}

func (s *ChannelService) DropUser(channel *ggc.Channel, user *ggc.User) error  {
	if channel == nil {
		return errors.New("channel is required")
	}
	if user == nil {
		return errors.New("user is required")
	}

	if _, err := s.GetChannel(ggc.ChannelFilter{Id: &channel.Id}); err != nil {
		return err
	}

	u := make([]*ggc.User, len(channel.Members))
	i := 0
	for  _, value := range channel.Members {
		if value.Username != user.Username {
			u[i] = value
			i++
		}
	}

	channel.Members = u

	return nil
}
