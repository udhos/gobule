package main

import (
	"log"
	"os"

	"github.com/udhos/gobule/parser"
)

func main() {
	status := parser.Run(os.Stdin)

	log.Printf("status: %d\n", status)
}
