package main

import (
	"log"

	"github.com/toxinu/dekiteru/cmd"
)

func main() {
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
