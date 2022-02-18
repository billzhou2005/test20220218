package main

import (
	"log"

	"github.com/segmentio/ksuid"
)

// test for Game ID
func main() {
	for i := 0; i < 2; i++ {
		id := ksuid.New()
		log.Println(id)
	}
}
