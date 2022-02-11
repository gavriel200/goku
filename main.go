package main

import (
	"github.com/gavriel200/goku/server"
)

func main() {
	server := server.NewServer()
	server.Start()
}
