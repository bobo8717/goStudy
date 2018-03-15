package server

type Server struct {
	opts Options
	Host string
}

func NewServer(opts ...Option) *Server {
	return &Server{
		opts: newOptions(opts...),
	}
}

func (s *Server) GetName() string {
	return s.opts.Name
}

func (s *Server) GetTag() string {
	return s.opts.Tag
}
