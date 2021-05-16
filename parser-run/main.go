package main

import (
	"log"
	"os"

	"github.com/udhos/gobule/parser"
)

func main() {
	result := parser.Run(os.Stdin)

	log.Printf("result: eval=%v status=%d errors=%d last_error: [%s]\n", result.Eval, result.Status, result.Errors, result.LastError)
}
