package main

import (
	"fmt"

	"github.com/ArchishmanSengupta/tinyid"
)

func main() {
	id, err := tinyid.NewTinyID()
	if err != nil {
		fmt.Print("some problem")
	}
	fmt.Println("Generated ID:", id)
}
