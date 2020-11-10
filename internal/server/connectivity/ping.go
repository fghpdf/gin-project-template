package connectivity

type Model struct {
	Status string `json:"status"`
}

type Server interface {
	Ping() (*Model, error)
}

type server struct {
}

func NewServer() Server {
	return &server{}
}

// Ping check server status
func (s *server) Ping() (*Model, error) {
	model := &Model{Status: "pong"}
	return model, nil
}
