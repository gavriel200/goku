package main

import (
	"fmt"

	"github.com/gavriel200/goku/server"
)

func main() {
	server := server.NewServer("test")

	fmt.Println(server.Name)
}
