package main

import (
	"fmt"
	"log"
	"os"

	"github.com/udhos/gobule/bulexer"
)

func main() {

	lexer := bulexer.New(os.Stdin)

	log.Printf("reading from stdin...")

LOOP:
	for {
		token := lexer.Next()
		fmt.Printf("%s\n", token.String())
		switch token.Type {
		case bulexer.TkEOF, bulexer.TkError:
			break LOOP
		}
	}

	log.Printf("reading from stdin...done")
}
