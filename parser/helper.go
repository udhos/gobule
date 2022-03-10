package parser

import (
	"errors"
	"strconv"
	"strings"
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

func (s scalar) GreaterThan(ss scalar) bool {
	if s.scalarType == scalarText && ss.scalarType == scalarText {
		return s.text > ss.text
	}
	if s.scalarType == scalarNumber && ss.scalarType == scalarNumber {
		return s.number > ss.number
	}
	return false
}

func (s scalar) GreaterThanOrEqual(ss scalar) bool {
	if s.scalarType == scalarText && ss.scalarType == scalarText {
		return s.text >= ss.text
	}
	if s.scalarType == scalarNumber && ss.scalarType == scalarNumber {
		return s.number >= ss.number
	}
	return false
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

func parseList(s string) ([]interface{}, error) {

	var list []interface{}

	s = strings.TrimSpace(s)
	if len(s) < 2 {
		return list, errors.New("too short")
	}

	if s[0] != '[' {
		return list, errors.New("missing opening square bracket")
	}

	last := len(s) - 1
	if s[last] != ']' {
		return list, errors.New("missing closing square bracket")
	}

	s = s[1:last]

	fields := strings.Fields(s)
	for _, f := range fields {
		n, errConv := parseInt(f)
		if errConv == nil {
			list = append(list, n)
			continue
		}
		list = append(list, f)
	}

	return list, nil
}

func parseInt(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
