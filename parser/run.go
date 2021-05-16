package parser

import (
	"io"
	"log"

	"github.com/udhos/gobule/bulexer"
)

func Run(input io.Reader) int {

	lex := &Lex{lex: bulexer.New(input), debug: true}

	status := yyParse(lex)

	return status
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
		log.Printf("parser.Lex: %s parserId=%d", token.String(), parserId)
	}

	if token.Type == bulexer.TkEOF {
		return 0 // real EOF for parser
	}

	return parserId
}

func parserToken(lexerId bulexer.TokenType) int {
	return 999 // FIXME WRITEME
}

func (l *Lex) Error(s string) {
	l.errors++
	log.Printf("parser.Lex.Error: errors=%d: %s", l.errors, s)
}
