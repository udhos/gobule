package parser

const (
	parserTokenIDFirst = TkKeywordTrue
)

var result Result

type valueType int

const (
	valueText valueType = iota
	valueNumber
)

type nodeValue struct {
	nodeType valueType
	text     string
	number   int
}

func (n nodeValue) Equals(m nodeValue) bool {
	if n.nodeType == valueText && m.nodeType == valueText {
		return n.text == m.text
	}
	if n.nodeType == valueNumber && m.nodeType == valueNumber {
		return n.number == m.number
	}
	return false
}

func contains(list []nodeValue, wanted nodeValue) bool {
	for _, v := range list {
		if wanted.Equals(v) {
			return true
		}
	}
	return false
}

var valueList = []nodeValue{}
