package main

import (
	"fmt"

	"github.com/ArturDV-main/Web-App-Advanced/internal/nserver"
)

func main() {
	var s nserver.Nserver
	s.StartServer("80")
	fmt.Println("Grate.")
}
