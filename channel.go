package gogochat

import (
	"errors"
	"fmt"
)

type Channel struct {
	Id int `json:"id"`

	Name  string `json:"name"`
	Members []*User `json:"members"`
}

func (c *Channel) Validate() error {
	if c.Name == "" {
		return errors.New("channel name is required")
	}
	return nil
}

func (c Channel) String() string {
	return fmt.Sprintf("%v channel", c.Name)
}

type ChannelService interface {
	CreateChannel(channel *Channel) error
	JoinUser(channel *Channel, user *User) error
	DropUser(channel *Channel, user *User) error
	GetChannel(filter ChannelFilter) (*Channel, error)
	ListChannels() ([]*Channel, error)
}

type ChannelFilter struct {
	Id   *int   `json:"id"`
	Name string `json:"name"`
}
