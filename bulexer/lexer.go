package bulexer

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
)

type lexState int

const (
	stBlank lexState = iota
	stNumber
	stText
	stIdent
	stDot
	stEOF
)

// Lexer is a lexical scanner.
type Lexer struct {
	reader io.ByteScanner
	state  lexState
	buf    lexBuf // buf could be bytes.Buffer, but we use an interface for testing
	debug  bool
}

type lexBuf interface {
	fmt.Stringer
	io.ByteWriter
	Reset() // is this define anywhere?
}

// TokenType identifies the type for lexical tokens returned by scanner.
type TokenType int

// Name returns the name of the token type.
func (t TokenType) Name() string {
	return tokenName[t]
}

// Token holds a token matched by the scanner.
type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	return t.Type.Name() + "(" + t.Value + ")"
}

// Token types returned by the scanner.
const (
	TkKeywordTrue TokenType = iota
	TkKeywordFalse
	TkKeywordOr
	TkKeywordAnd
	TkKeywordNot
	TkKeywordContains
	TkKeywordCurrentTime
	TkKeywordNumber
	TkKeywordList
	TkKeywordVersion
	TkNumber
	TkText
	TkIdent
	TkParL
	TkParR
	TkSBktL
	TkSBktR
	TkEQ
	TkLT
	TkGT
	TkNE
	TkGE
	TkLE
	TkDot
	TkError
	TkEOF
)

var tokenName = []string{
	"KW-TRUE",
	"KW-FALSE",
	"KW-OR",
	"KW-AND",
	"KW-NOT",
	"KW-CONTAINS",
	"KW-CURRENTTIME",
	"KW-NUMBER",
	"KW-LIST",
	"KW-VERSION",
	"NUMBER",
	"TEXT",
	"IDENT",
	"LPAR",
	"RPAR",
	"LSBKT",
	"RSBKT",
	"EQ",
	"LT",
	"GT",
	"NE",
	"GE",
	"LE",
	"DOT",
	"ERROR",
	"EOF",
}

// New creates a lexical scanner.
func New(input io.Reader) *Lexer {
	return &Lexer{reader: bufio.NewReader(input), buf: &bytes.Buffer{}}
}

func isBlank(b byte) bool {
	switch b {
	case ' ', '\r', '\n', '\t':
		return true
	}
	return false
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

var eof = Token{Type: TkEOF}

// Next returns the next scanner token, or TkEOF, or TkError.
func (l *Lexer) Next() Token {

SCANNER:
	for l.state != stEOF {

		b, errByte := l.reader.ReadByte()

		if l.debug {
			log.Printf("state=%d: byte: %d '%c' err:%v", l.state, b, b, errByte)
		}

		switch l.state {

		case stBlank:
			switch errByte {
			case io.EOF:
				l.state = stEOF
				break SCANNER
			case nil:
			default:
				return Token{Type: TkError, Value: errByte.Error()}
			}
			switch {
			case isBlank(b):
			case isDigit(b):
				if errSave := l.buf.WriteByte(b); errSave != nil {
					return Token{Type: TkError, Value: errSave.Error()}
				}
				l.state = stNumber
			case isLetter(b):
				if errSave := l.buf.WriteByte(b); errSave != nil {
					return Token{Type: TkError, Value: errSave.Error()}
				}
				l.state = stIdent
			case b == '\'':
				l.state = stText
			case b == '(':
				return Token{Type: TkParL, Value: "("}
			case b == ')':
				return Token{Type: TkParR, Value: ")"}
			case b == '[':
				return Token{Type: TkSBktL, Value: "["}
			case b == ']':
				return Token{Type: TkSBktR, Value: "]"}
			case b == '=':
				return Token{Type: TkEQ, Value: "="}
			case b == '<':
				bb, errBB := l.reader.ReadByte()
				switch errBB {
				case io.EOF:
					l.state = stEOF
					return Token{Type: TkLT, Value: "<"}
				case nil:
				default:
					return Token{Type: TkError, Value: errBB.Error()}
				}
				if bb == '=' {
					return Token{Type: TkLE, Value: "<="}
				}
				if errUnread := l.reader.UnreadByte(); errUnread != nil {
					return Token{Type: TkError, Value: errUnread.Error()}
				}
				return Token{Type: TkLT, Value: "<"}
			case b == '>':
				bb, errBB := l.reader.ReadByte()
				switch errBB {
				case io.EOF:
					l.state = stEOF
					return Token{Type: TkGT, Value: ">"}
				case nil:
				default:
					return Token{Type: TkError, Value: errBB.Error()}
				}
				if bb == '=' {
					return Token{Type: TkGE, Value: ">="}
				}
				if errUnread := l.reader.UnreadByte(); errUnread != nil {
					return Token{Type: TkError, Value: errUnread.Error()}
				}
				return Token{Type: TkGT, Value: ">"}
			case b == '!':
				bb, errBB := l.reader.ReadByte()
				if errBB != nil {
					return Token{Type: TkError, Value: "error-after-unexpected:!"}
				}
				if bb == '=' {
					return Token{Type: TkNE, Value: "!="}
				}
				return Token{Type: TkError, Value: "unexpected:!"}
			default:
				return Token{Type: TkError, Value: fmt.Sprintf("unexpected byte: %d '%c'", b, b)}
			}

		case stIdent:
			switch errByte {
			case io.EOF:
				l.state = stEOF
				return l.consumeIdent()
			case nil:
			default:
				return Token{Type: TkError, Value: errByte.Error()}
			}
			switch {
			case isBlank(b):
				l.state = stBlank
				return l.consumeIdent()
			case isDigit(b):
				if errSave := l.buf.WriteByte(b); errSave != nil {
					return Token{Type: TkError, Value: errSave.Error()}
				}
			case isLetter(b):
				if errSave := l.buf.WriteByte(b); errSave != nil {
					return Token{Type: TkError, Value: errSave.Error()}
				}
			default:
				if errUnread := l.reader.UnreadByte(); errUnread != nil {
					return Token{Type: TkError, Value: errUnread.Error()}
				}
				l.state = stBlank
				return l.consumeIdent()
			}

		case stNumber:
			switch errByte {
			case io.EOF:
				l.state = stEOF
				return l.consumeNumber()
			case nil:
			default:
				return Token{Type: TkError, Value: errByte.Error()}
			}
			switch {
			case b == '.':
				l.state = stDot
				return l.consumeNumber()
			case isBlank(b):
				l.state = stBlank
				return l.consumeNumber()
			case isDigit(b):
				if errSave := l.buf.WriteByte(b); errSave != nil {
					return Token{Type: TkError, Value: errSave.Error()}
				}
			case isLetter(b):
				return Token{Type: TkError, Value: "letter-after-number"}
			default:
				if errUnread := l.reader.UnreadByte(); errUnread != nil {
					return Token{Type: TkError, Value: errUnread.Error()}
				}
				l.state = stBlank
				return l.consumeNumber()
			}

		case stDot:
			switch errByte {
			case io.EOF:
				l.state = stEOF
				return Token{Type: TkError, Value: "EOF-after-version-dot"}
			case nil:
			default:
				return Token{Type: TkError, Value: errByte.Error()}
			}
			switch {
			case isDigit(b):
				if errUnread := l.reader.UnreadByte(); errUnread != nil {
					return Token{Type: TkError, Value: errUnread.Error()}
				}
				l.state = stNumber
				return Token{Type: TkDot, Value: "."}
			default:
				return Token{Type: TkError, Value: fmt.Sprintf("non-digit byte after version dot: %d '%c'", b, b)}
			}

		case stText:
			switch errByte {
			case io.EOF:
				l.state = stEOF
				return Token{Type: TkError, Value: "EOF-after-unterminated-text"}
			case nil:
			default:
				return Token{Type: TkError, Value: errByte.Error()}
			}
			switch {
			case b == '\'':
				l.state = stBlank
				return l.consume(Token{Type: TkText})
			default:
				if errSave := l.buf.WriteByte(b); errSave != nil {
					return Token{Type: TkError, Value: errSave.Error()}
				}
			}

		default:
			return Token{Type: TkError, Value: fmt.Sprintf("unexpected state:%d", l.state)}
		}
	}

	return eof
}

var keywords = map[string]TokenType{
	"true":        TkKeywordTrue,
	"false":       TkKeywordFalse,
	"AND":         TkKeywordAnd,
	"OR":          TkKeywordOr,
	"NOT":         TkKeywordNot,
	"CONTAINS":    TkKeywordContains,
	"CurrentTime": TkKeywordCurrentTime,
	"Number":      TkKeywordNumber,
	"List":        TkKeywordList,
	"Version":     TkKeywordVersion,
}

func (l *Lexer) consumeNumber() Token {
	return l.consume(Token{Type: TkNumber})
}

func (l *Lexer) consumeIdent() Token {
	token := l.consume(Token{})
	if tt, found := keywords[token.Value]; found {
		// Identifier is a keyword
		token.Type = tt
	} else {
		// Identifier is a variable
		token.Type = TkIdent
	}
	return token
}

func (l *Lexer) consume(token Token) Token {
	token.Value = l.buf.String() // save value
	l.buf.Reset()
	return token
}
