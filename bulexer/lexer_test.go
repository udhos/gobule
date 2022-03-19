package bulexer

import (
	"bytes"
	"errors"
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
	{"dot is not a valid symbol", ".", "ERROR(unexpected byte: 46 '.')"},
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
	{"exclamation alone is not valid symbol", "!", "ERROR(error-after-unexpected:!)"},
	{"exclamation alone is not valid symbol", " ! ", "ERROR(unexpected:!)"},
	{"all valid symbols", "\n true false AND OR NOT CONTAINS CurrentTime Number List myVar123 123 'not long text' ( ) [ ] < > != = <= >= ",
		"KW-TRUE(true) KW-FALSE(false) KW-AND(AND) KW-OR(OR) KW-NOT(NOT) KW-CONTAINS(CONTAINS) " +
			"KW-CURRENTTIME(CurrentTime) KW-NUMBER(Number) KW-LIST(List) IDENT(myVar123) NUMBER(123) " +
			"TEXT(not long text) LPAR(() RPAR()) LSBKT([) RSBKT(]) LT(<) GT(>) NE(!=) EQ(=) LE(<=) GE(>=) EOF()"},
	{"basic expression 1", "a=b", "IDENT(a) EQ(=) IDENT(b) EOF()"},
	{"basic expression 2", "1=1", "NUMBER(1) EQ(=) NUMBER(1) EOF()"},
	{"basic expression 3", "'1'='a'", "TEXT(1) EQ(=) TEXT(a) EOF()"},
	{"basic expression 4", "true=false", "KW-TRUE(true) EQ(=) KW-FALSE(false) EOF()"},
	{"long expression", "NOT [ 1 b 'c' ] NOT CONTAINS Number('3')",
		"KW-NOT(NOT) LSBKT([) NUMBER(1) IDENT(b) TEXT(c) RSBKT(]) KW-NOT(NOT) KW-CONTAINS(CONTAINS) KW-NUMBER(Number) LPAR(() TEXT(3) RPAR()) EOF()"},
	{"can concat some symbols", "true'text'123(var'text2'false", "KW-TRUE(true) TEXT(text) NUMBER(123) LPAR(() IDENT(var) TEXT(text2) KW-FALSE(false) EOF()"},
	{"version number", "Version(1.2.3)", "KW-VERSION(Version) LPAR(() NUMBER(1) DOT(.) NUMBER(2) DOT(.) NUMBER(3) RPAR()) EOF()"},
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

type brokenInputMock struct {
	buf  []byte
	done bool
}

func (r *brokenInputMock) Read(p []byte) (n int, err error) {
	if r.done {
		return 0, errors.New("brokenInputMock ERROR")
	}
	r.done = true
	p = append(p, r.buf...)
	return len(r.buf), nil
}

func brokenInput(t *testing.T, r *brokenInputMock) {
	lexer := New(r)
	lexer.debug = true
SCANNER:
	for {
		token := lexer.Next()
		switch token.Type {
		case TkEOF:
			t.Errorf("unexpected EOF")
			break SCANNER
		case TkError:
			break SCANNER
		}
	}
}

func TestBrokenInput(t *testing.T) {
	brokenInput(t, &brokenInputMock{buf: []byte("true")})
	brokenInput(t, &brokenInputMock{buf: []byte("tr")})
	brokenInput(t, &brokenInputMock{buf: []byte("true"), done: true})
	brokenInput(t, &brokenInputMock{buf: []byte{}, done: true})
}
