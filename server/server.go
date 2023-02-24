package server

import (
	"log"
	"net/http"

	"github.com/ebizno/Ula/server/handle"
)

type Server struct {
	Port string
}

func NewServer(port string) *Server {
	return &Server{
		Port: port,
	}
}

func (s *Server) Start() {

	upload := handle.NewUpload(s.Port)

	http.HandleFunc("/upload", upload.UploadFile)
	http.HandleFunc("/", upload.PageUpload)

	if err := http.ListenAndServe(":"+s.Port, nil); err != nil {
		log.Fatal(err)
	}
}
