package server

type Server struct {
	Name string
}

func NewServer(Name string) *Server {
	return &Server{Name}
}
