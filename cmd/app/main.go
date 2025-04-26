package main

import (
	"proyecto/http"
)

func main() {
	s := http.NewServer()

	err := s.Serve()
	if err != nil {
		// NOTE: THIS GETS HANDLED LOL
		panic(err)
	}
}


