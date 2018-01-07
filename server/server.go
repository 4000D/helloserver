package server

import (
	"net/http"

	"helloserver/utils"
)

var (
	logger = utils.NewLogger()
)

type Server struct {
	config *utils.Config
}

func NewServer(config *utils.Config) *Server {
	server := &Server{config: config}
	return server
}

func (s *Server) RunServer() error {
	port := s.config.Port

	logger.Println("listening " + port + " port")

	http.HandleFunc("/", s.httpHandler)

	http.ListenAndServe(":"+port, nil) // TOOD: use config.port

	return nil
}

func (s *Server) httpHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(s.config.Greeting))
}
