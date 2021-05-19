package bulexer

import (
	"bytes"
	"strings"
	"testing"
)

type lexerTest struct {
	name           string
	input          string
	expectedResult string
}

var testTable = []lexerTest{
	{"empty", "", "EOF()"},
	{"newline", "\n", "EOF()"},
	{"keyword true newline", "true\n", "KW-TRUE(true) EOF()"},
	{"keyword true", "true", "KW-TRUE(true) EOF()"},
	{"keyword false", "false", "KW-FALSE(false) EOF()"},
	{"keyword AND", "AND", "KW-AND(AND) EOF()"},
	{"keyword OR", "OR", "KW-OR(OR) EOF()"},
	{"keyword NOT", "NOT", "KW-NOT(NOT) EOF()"},
	{"keyword CONTAINS", "CONTAINS", "KW-CONTAINS(CONTAINS) EOF()"},
	{"keyword CurrentTime", "CurrentTime", "KW-CURRENTTIME(CurrentTime) EOF()"},
	{"keyword Number", "Number", "KW-NUMBER(Number) EOF()"},
	{"keyword List", "List", "KW-LIST(List) EOF()"},
	{"variable", "myVar123", "IDENT(myVar123) EOF()"},
	{"number", "123", "NUMBER(123) EOF()"},
	{"number+ident", "123abc", "ERROR(letter-after-number)"},
	{"text", "'short text'", "TEXT(short text) EOF()"},
	{"braces1", "()[]", "LPAR(() RPAR()) LSBKT([) RSBKT(]) EOF()"},
	{"braces2", "  (  )  [  ]  ", "LPAR(() RPAR()) LSBKT([) RSBKT(]) EOF()"},
	{"compare1", "<>!==<=>=", "LT(<) GT(>) NE(!=) EQ(=) LE(<=) GE(>=) EOF()"},
	{"compare2", "  <  >  !=  =  <=  >=  ", "LT(<) GT(>) NE(!=) EQ(=) LE(<=) GE(>=) EOF()"},
	{"compare3 invalid unequal", "<>", "LT(<) GT(>) EOF()"},
	{"compare4 invalid unequal", "><", "GT(>) LT(<) EOF()"},
	{"compare5 double equals", "==", "EQ(=) EQ(=) EOF()"},
	{"exclamation is not valid symbol", "!", "ERROR(error-after-unexpected:!)"},
	{"all valid symbols", "\n true false AND OR NOT CONTAINS CurrentTime Number List myVar123 123 'not long text' ( ) [ ] < > != = <= >= ", "KW-TRUE(true) KW-FALSE(false) KW-AND(AND) KW-OR(OR) KW-NOT(NOT) KW-CONTAINS(CONTAINS) KW-CURRENTTIME(CurrentTime) KW-NUMBER(Number) KW-LIST(List) IDENT(myVar123) NUMBER(123) TEXT(not long text) LPAR(() RPAR()) LSBKT([) RSBKT(]) LT(<) GT(>) NE(!=) EQ(=) LE(<=) GE(>=) EOF()"},
}

func TestScanner(t *testing.T) {

	for _, data := range testTable {

		lexer := New(bytes.NewBufferString(data.input))

		var tokenList []string

	SCANNER:
		for {
			token := lexer.Next()
			//t.Logf("%s: %s\n", data.name, token.String())
			tokenList = append(tokenList, token.String())
			switch token.Type {
			case TkEOF, TkError:
				break SCANNER
			}
		}

		result := strings.Join(tokenList, " ")

		if result != data.expectedResult {
			t.Errorf("%s: input=[%s] expected=[%s] got=[%s]\n", data.name, data.input, data.expectedResult, result)
		}

	}
}
