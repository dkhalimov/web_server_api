package main

import (
	"log"

	"github.com/dkhalimov/web_server_api/internal/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err = nil {
		log.Fatal(err)
	}
}