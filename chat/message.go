package chat

import "time"

type Message struct {
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func (self *Message) String() string {
	return "Anonymous at " + self.CreatedAt.String() + " says " + self.Body
}
