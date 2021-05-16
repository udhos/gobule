package main

import (
	"log"
	"os"

	"github.com/udhos/gobule/lexer"
)

func main() {
	log.Printf("creating lexer\n")
	lexer := lexer.New(os.Stdin)
	log.Printf("reading from stdin")
	lexer.Run()
	log.Printf("done")
}
