package parser

import (
	"strconv"
)

const (
	parserTokenIDFirst = TkKeywordTrue
)

type sType int

const (
	scalarText sType = iota
	scalarNumber
)

type scalar struct {
	scalarType sType
	text       string
	number     int64
}

func (s scalar) greaterThan(ss scalar) bool {
	if s.scalarType == scalarText && ss.scalarType == scalarText {
		return s.text > ss.text
	}
	if s.scalarType == scalarNumber && ss.scalarType == scalarNumber {
		return s.number > ss.number
	}
	return false
}

func (s scalar) equals(ss scalar) bool {
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
		if wanted.equals(v) {
			return true
		}
	}
	return false
}

func parseInt(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
