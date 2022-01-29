package main

import (
	"cubetiq/registry/cmd"
	"log"
)

func main() {
	if err := cmd.Start(); err != nil {
		log.Fatal("Failed to start ngati registry:", err)
	}
}
