# gobule
Golang Go parser for the Bule Language

## Usage

```
package main

import (
	"log"

	"github.com/udhos/gobule/parser"
)

func main() {
	vars := map[string]interface{}{
		"platform": "android",
	}

	result := parser.RunString("platform = 'android'", vars, false)

	log.Printf("result: eval=%v status=%d errors=%d last_error: [%s]\n", result.Eval, result.Status, result.Errors, result.LastError)
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

### ModernC goyac

    go get modernc.org/goyacc

https://gitlab.com/cznic/goyacc

### Golang goyacc

    go get golang.org/x/tools/cmd/goyacc

https://blog.golang.org/generate

