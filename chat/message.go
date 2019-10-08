package chat

import "time"

type Message struct {
	UserName  string    `json:"user_name"`
	Body      string    `json:"body"`
	Timestamp time.Time `json:"time_stamp"`
}

func (self *Message) String() string {
	return self.UserName + " at " + self.Timestamp.String() + " says " + self.Body
}
