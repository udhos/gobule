package parser

import (
	"bytes"
	"encoding/json"
	"testing"
)

type expect int

const (
	expectError expect = iota
	expectTrue
	expectFalse
)

type parserTest struct {
	name           string
	input          string
	vars           string
	expectedResult expect
}

var testTable = []parserTest{
	{"empty", "", "{}", expectError},
	{"missing variable", "name", "{}", expectError},
	{"missing variable", "[name] CONTAINS 'John'", "{}", expectError},
	{"simple var 1", "[name] CONTAINS 'John'", `{"name":"Jane"}`, expectFalse},
	{"simple var 2", "[name] CONTAINS 'John'", `{"name":"John"}`, expectTrue},
	{"simple var 2", "platform = 'android'", `{"platform":"android"}`, expectTrue},
	{"simple operator 1", "1 = 1", `{}`, expectTrue},
	{"simple operator 2", "1 != 1", `{}`, expectFalse},
	{"simple operator 3", "1 > 1", `{}`, expectFalse},
	{"simple operator 4", "1 >= 1", `{}`, expectTrue},
	{"simple operator 5", "1 < 1", `{}`, expectFalse},
	{"simple operator 6", "1 <= 1", `{}`, expectTrue},
	{"simple operator text 1", "'1' = '1'", `{}`, expectTrue},
	{"simple operator text 2", "'1' != '1'", `{}`, expectFalse},
	{"simple operator text 3", "'1' > '1'", `{}`, expectFalse},
	{"simple operator text 4", "'1' >= '1'", `{}`, expectTrue},
	{"simple operator text 5", "'1' < '1'", `{}`, expectFalse},
	{"simple operator text 6", "'1' <= '1'", `{}`, expectTrue},
	{"simple operator type mismatch 1", "1 = '1'", `{}`, expectFalse},
	{"simple operator type mismatch 2", "1 != '1'", `{}`, expectTrue},
	{"simple operator type mismatch 3", "1 > '1'", `{}`, expectError},
	{"simple operator type mismatch 4", "1 >= '1'", `{}`, expectError},
	{"simple operator type mismatch 5", "1 < '1'", `{}`, expectError},
	{"simple operator type mismatch 6", "1 <= '1'", `{}`, expectError},
}

func TestParser(t *testing.T) {

	for _, data := range testTable {

		vars := map[string]string{}

		if errJSON := json.Unmarshal([]byte(data.vars), &vars); errJSON != nil {
			t.Errorf("%s: json: %v: vars=%s", data.name, errJSON, data.vars)
		}

		debug := false

		result := Run(bytes.NewBufferString(data.input), vars, debug)

		if data.expectedResult == expectError {
			// error expected
			if result.Status == 0 && result.Errors == 0 {
				// no error
				t.Errorf("%s: input=[%s] expected=ERROR got: status=%d errors=%d error:%s\n",
					data.name, data.input, result.Status, result.Errors, result.LastError)
			}
			continue
		}

		// error unexpected

		if result.Status != 0 || result.Errors != 0 {
			// error found
			t.Errorf("%s: input=[%s] expected=noerror got: status=%d errors=%d error:%s\n",
				data.name, data.input, result.Status, result.Errors, result.LastError)
			continue
		}

		// no error

		expectedEval := data.expectedResult == expectTrue
		if expectedEval != result.Eval {
			t.Errorf("%s: input=[%s] expected=%v got=%v\n",
				data.name, data.input, expectedEval, result.Eval)
		}
	}
}