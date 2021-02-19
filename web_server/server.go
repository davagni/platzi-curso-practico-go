package main

import "net/http"

//Server struct
type Server struct {
	port   string
	router *Router
}

//NewServer a
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

//Handle a
func (s *Server) Handle(path string, handler http.HandlerFunc) {
	s.router.rules[path] = handler
}

//AddMiddleWare a
func (s *Server) AddMiddleWare(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

//Listen a
func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}
