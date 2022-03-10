package conv

import (
	"strconv"
	"testing"
)

type numberTest struct {
	name           string
	input          []int64
	expectedResult int64
}

type stringTest struct {
	name           string
	input          string
	expectedResult string
}

var testTableNumber = []numberTest{
	{"0,0,0", []int64{0, 0, 0}, 0},
	{"1,1,1", []int64{1, 1, 1}, 100010001},
	{"1,2,3", []int64{1, 2, 3}, 100020003},
	{"3,2,1", []int64{3, 2, 1}, 300020001},
}

var testTableString = []stringTest{
	{"0.0.0", "0.0.0", "0"},
	{"1.1.1", "1.1.1", "100010001"},
	{"1.2.3", "1.2.3", "100020003"},
	{"3.2.1", "3.2.1", "300020001"},
}

func TestNumber(t *testing.T) {
	for _, data := range testTableNumber {
		result := VersionToNumber(data.input[0], data.input[1], data.input[2])
		if result != data.expectedResult {
			t.Errorf("%s: input=%v expected=%d got=%d", data.name, data.input, data.expectedResult, result)
		}
	}
}

func TestString(t *testing.T) {
	for _, data := range testTableString {
		result := VersionToString(data.input)
		if result != data.expectedResult {
			t.Errorf("%s: input=%s expected=%s got=%s", data.name, data.input, data.expectedResult, result)
		}
	}
}

func TestMatch(t *testing.T) {
	for i, number := range testTableNumber {
		str := testTableString[i]
		expectedNumber := strconv.FormatInt(number.expectedResult, 10)
		if str.expectedResult != expectedNumber {
			t.Errorf("%s: mismatch expectedNumber=%s expectedString=%s", number.name, expectedNumber, str.expectedResult)
		}
	}
}
