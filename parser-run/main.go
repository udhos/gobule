package main

import (
	"log"
	"os"

	"github.com/udhos/gobule/parser"
)

func main() {
	var vars = map[string]interface{}{
		"name":   "John",
		"number": "123",
	}

	log.Printf("FIXME vars: %v", vars)

	debug := false

	result := parser.Run(os.Stdin, vars, debug)

	log.Printf("result: eval=%v status=%d errors=%d last_error: [%s]\n", result.Eval, result.Status, result.Errors, result.LastError)
}
