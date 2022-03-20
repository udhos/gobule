package parser

import (
	"bytes"
	"testing"

	"github.com/udhos/gobule/bulexer"
)

func TestParserInternal(t *testing.T) {

	if symName := yySymName(0x7f - 1); symName != "'~'" {
		t.Errorf("unexpected sym name: '%s'", symName)
	}

	if symName := yySymName(0x7f); symName != "127" {
		t.Errorf("unexpected sym name: '%s'", symName)
	}

	{
		saveDebug := yyDebug
		yyDebug = 4 /* max yyDebug */
		lex := &Lex{lex: bulexer.New(&bytes.Buffer{}), debug: true}
		var lval yySymType
		if yychar := yylex1(lex, &lval); yychar != yyEofCode {
			t.Errorf("unexpected non-eof from yylex1(): expected=%d got=%d", yyEofCode, yychar)
		}
		yyDebug = saveDebug
	}
}
