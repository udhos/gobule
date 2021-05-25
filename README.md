[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/udhos/gobule/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/udhos/gobule)](https://goreportcard.com/report/github.com/udhos/gobule)
[![Go Reference](https://pkg.go.dev/badge/github.com/udhos/gobule.svg)](https://pkg.go.dev/github.com/udhos/gobule)

# gobule

Golang Go parser for the Bule Language

## Usage

```
package main

import (
	"log"
	"os"

	"github.com/udhos/gobule/parser"
)

func main() {
	vars := map[string]interface{}{
		"platform": "android",
	}

	envDebug := os.Getenv("DEBUG")
	debug := envDebug != ""

	log.Printf("DEBUG=[%s] debug=%v", envDebug, debug)

	result := parser.RunString("platform = 'android'", vars, debug)

	if result.IsError() {
		log.Printf("ERROR status=%d errors=%d last_error: [%s]\n", result.Status, result.Errors, result.LastError)
		return
	}

	log.Printf("result: %v", result.Eval)
}
```

## Build

Use this recipe if you need to build the parser for development.

```
git clone https://github.com/udhos/gobule
cd gobule
go generate ./parser ;# generate parser
go test ./parser     ;# run tests
go install ./parser  ;# build
```

## Bule Language

https://github.com/johnowl/owl-rules

## Tokens

https://github.com/johnowl/owl-rules/blob/master/src/main/kotlin/com/johnowl/rules/RulesEvaluator.kt

## Syntax Analyser Generators

### ModernC goyacc

    go get modernc.org/goyacc

https://gitlab.com/cznic/goyacc

### Golang goyacc

    go get golang.org/x/tools/cmd/goyacc

https://blog.golang.org/generate

