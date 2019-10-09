package chat

type Server struct {
	Messages []*Message `json:"messages"`
}

func NewServer() *Server {
	Messages := []*Message{}

	return &Server{
		Messages,
	}
}
