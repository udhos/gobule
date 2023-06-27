// Package main implements parser-run.
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

	log.Printf("FIXME input vars: %v", vars)

	envDebug := os.Getenv("DEBUG")
	debug := envDebug != ""

	log.Printf("DEBUG=[%s] debug=%v", envDebug, debug)

	result := parser.Run(os.Stdin, vars, debug)

	if result.IsError() {
		log.Printf("ERROR status=%d errors=%d last_error: [%s]\n", result.Status, result.Errors, result.LastError)
		return
	}

	log.Printf("result: %v", result.Eval)
}
