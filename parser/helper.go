package parser

const (
	parserTokenIDFirst = TkKeywordTrue
)

var result Result

type sType int

const (
	scalarText sType = iota
	scalarNumber
)

type scalar struct {
	scalarType sType
	text       string
	number     int
}

func (s scalar) Equals(ss scalar) bool {
	if s.scalarType == scalarText && ss.scalarType == scalarText {
		return s.text == ss.text
	}
	if s.scalarType == scalarNumber && ss.scalarType == scalarNumber {
		return s.number == ss.number
	}
	return false
}

func contains(list []scalar, wanted scalar) bool {
	for _, v := range list {
		if wanted.Equals(v) {
			return true
		}
	}
	return false
}

var scalarList = []scalar{}
