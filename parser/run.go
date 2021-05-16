package parser

import (
	"io"
	"log"

	"github.com/udhos/gobule/bulexer"
)

type Result struct {
	Eval      bool
	Errors    int
	Status    int
	LastError string
}

func Run(input io.Reader) Result {

	lex := &Lex{lex: bulexer.New(input), debug: true}

	status := yyParse(lex)

	result.Errors = lex.errors
	result.Status = status

	return result
}

type Lex struct {
	lex    *bulexer.Lexer
	errors int
	debug  bool
}

func (l *Lex) Lex(lval *yySymType) int {

	token := l.lex.Next()

	parserId := parserToken(token.Type)

	if l.debug {
		log.Printf("parser.Lex: %s lexerId=%d parserId=%d", token.String(), token.Type, parserId)
	}

	if token.Type == bulexer.TkEOF {
		return 0 // real EOF for parser
	}

	// need to store values only for some terminals
	// when a parser rule action needs to consume the value
	// for example: variable, literal (number, text)
	switch parserId {
	case TkText:
		lval.typeText = token.Value
	case TkNumber:
		lval.typeNumber = token.Value
	}

	return parserId
}

func parserToken(lexerId bulexer.TokenType) int {
	return int(lexerId) + parserTokenIDFirst
}

func (l *Lex) Error(s string) {
	l.errors++
	if l.debug {
		log.Printf("parser.Lex.Error: errors=%d: %s", l.errors, s)
	}
	result.LastError = s
}
