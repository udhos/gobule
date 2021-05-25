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
	{"true", "true", "{}", expectTrue},
	{"false", "false", "{}", expectFalse},
	{"NOT true", "NOT true", "{}", expectFalse},
	{"NOT false", "NOT false", "{}", expectTrue},
	{"NOT NOT true", "NOT NOT true", "{}", expectTrue},
	{"NOT NOT false", "NOT NOT false", "{}", expectFalse},
	{"list 1", "[1 2 3 4] CONTAINS 4", "{}", expectTrue},
	{"list 1", "[1 2 3 4] NOT CONTAINS 4", "{}", expectFalse},
	{"list 1", "NOT [1 2 3 4] CONTAINS 4", "{}", expectFalse},
	{"list 2", "['blue' 'yellow' 'green'] CONTAINS 'pink'", "{}", expectFalse},
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
	{"simple or 1", "(1 > 2) OR (1 = 1)", "{}", expectTrue},
	{"simple or 2", "(1 > 2) OR (1 < 1)", "{}", expectFalse},
	{"simple or 3", "(1 < 2) OR (1 = 1)", "{}", expectTrue},
	{"simple and 1", "(1 > 2) AND (1 = 1)", "{}", expectFalse},
	{"simple and 2", "(1 > 2) AND (1 < 1)", "{}", expectFalse},
	{"simple and 3", "(1 < 2) AND (1 = 1)", "{}", expectTrue},
	{"number 1", "Number('11') >= 12", `{}`, expectFalse},
	{"number 2", "Number('12') >= 12", `{}`, expectTrue},
	{"number 3", "Number(version) >= 12", `{"version":"11"}`, expectFalse},
	{"number 4", "Number(version) >= 12", `{"version":"12"}`, expectTrue},
	{"number 5", "Number(11) >= 12", `{}`, expectError},
	{"number 6", "Number() >= 12", `{}`, expectError},
	{"number 7", "Number >= 12", `{}`, expectError},
	{"number 8", "Number('11')", `{}`, expectError},
	{"number 9", "Number('bob') < 20", `{}`, expectError},
	{"currenttime 1", "CurrentTime() > 0", `{}`, expectTrue},
	{"currenttime 2", "CurrentTime() < 250000", `{}`, expectTrue},
	{"currenttime 3", "CurrentTime() < 0", `{}`, expectFalse},
	{"list literal", "[1 2 3 4] CONTAINS Number(var1)", `{"var1":1}`, expectTrue},
	{"list function 1", "List('[1 2 3 4]') CONTAINS Number(var1)", `{"var1":1}`, expectTrue},
	{"list function 2", "List(var0) CONTAINS Number(var1)", `{"var0":[1,2,3,4],"var1":1}`, expectTrue},
	{"list function 3", "List(var0) CONTAINS Number(var1)", `{"var0":[1,2,3,4],"var1":"1"}`, expectTrue},
	{"list function 4", "List(var0) CONTAINS var1", `{"var0":["alpha","beta",1,2],"var1":"beta"}`, expectTrue},
}

func TestParser(t *testing.T) {

	for _, data := range testTable {

		vars := map[string]interface{}{}

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

		if result.IsError() {
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
