package parser

import (
	"io"
	"log"

	"github.com/udhos/gobule/lexer"
)

func Run(input io.Reader) int {

	lex := &Lex{lex: lexer.New(input)}

	status := yyParse(lex)

	return status
}

type Lex struct {
	lex *lexer.GobuleLexer
}

func (l *Lex) Lex(lval *yySymType) int {
	log.Printf("parser.Lex")
	return 0
}

func (l *Lex) Error(s string) {
	log.Printf("parser.Lex.Error: %s", s)
}

func (l *Lex) Errors() int {
	return 0
}
