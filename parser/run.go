package parser

import (
	"bytes"
	"io"
	"log"

	"github.com/udhos/gobule/bulexer"
)

// Result returns parser result.
type Result struct {
	Eval      bool
	Errors    int
	Status    int
	LastError string
}

// Run executes parser for input.
func Run(input io.Reader, vars map[string]interface{}, debug bool) Result {

	lex := &Lex{lex: bulexer.New(input), debug: debug, vars: vars}

	status := yyParse(lex)

	lex.result.Errors = lex.errors
	lex.result.Status = status

	return lex.result
}

// RunString executes parser for string.
func RunString(input string, vars map[string]interface{}, debug bool) Result {
	return Run(bytes.NewBufferString(input), vars, debug)
}

// Lex provides the lexical scanner interface required by the generated parser.
type Lex struct {
	lex    *bulexer.Lexer
	errors int
	debug  bool

	// context data for parser:
	vars       map[string]interface{} // input variables
	scalarList []scalar               // aux
	result     Result                 // output
}

// Lex is called by the syntatical parser to request the next token.
func (l *Lex) Lex(lval *yySymType) int {

	token := l.lex.Next()

	parserID := parserToken(token.Type)

	if l.debug {
		log.Printf("parser.Lex: %s lexerId=%d parserId=%d", token.String(), token.Type, parserID)
	}

	if token.Type == bulexer.TkEOF {
		return 0 // real EOF for parser
	}

	// need to store values only for some terminals
	// when a parser rule action needs to consume the value
	// for example: variable, literal (number, text)
	switch parserID {
	case TkText, TkNumber, TkIdent:
		lval.typeString = token.Value
	}

	return parserID
}

func parserToken(lexerID bulexer.TokenType) int {
	return int(lexerID) + parserTokenIDFirst
}

func (l *Lex) Error(s string) {
	l.errors++
	if l.debug {
		log.Printf("parser.Lex.Error: errors=%d: %s", l.errors, s)
	}
	l.result.LastError = s
}
