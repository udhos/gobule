package parser

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
	{"true with nil vars", "true", "", expectTrue},
	{"var with nil vars", "Number(version) >= 12", "", expectError},
	{"false", "false", "{}", expectFalse},
	{"false deeply enclosed", "(((((false)))))", "{}", expectFalse},
	{"unbalanced parenthesis", "(((((true))))", "{}", expectError},
	{"NOT true", "NOT true", "{}", expectFalse},
	{"NOT(true)", "NOT(true)", "{}", expectFalse},
	{"NOT false", "NOT false", "{}", expectTrue},
	{"NOT(false)", "NOT(false)", "{}", expectTrue},
	{"NOT NOT true", "NOT NOT true", "{}", expectTrue},
	{"NOT(NOT true)", "NOT(NOT true)", "{}", expectTrue},
	{"NOT NOT false", "NOT NOT false", "{}", expectFalse},
	{"NOT(NOT false)", "NOT(NOT false)", "{}", expectFalse},
	{"list empty", "[] CONTAINS 4", "{}", expectFalse},
	{"list 1", "[1 2 3 4] CONTAINS 4", "{}", expectTrue},
	{"list 2", "[1 2 3 4] NOT CONTAINS 4", "{}", expectFalse},
	{"list 3", "NOT [1 2 3 4] CONTAINS 4", "{}", expectFalse},
	{"list 4", "['blue' 'yellow' 'green'] CONTAINS 'pink'", "{}", expectFalse},
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
	{"simple operator 7", "2 = 3", `{}`, expectFalse},
	{"simple operator 8", "2 != 3", `{}`, expectTrue},
	{"simple operator 9", "2 > 3", `{}`, expectFalse},
	{"simple operator 10", "2 >= 3", `{}`, expectFalse},
	{"simple operator 11", "2 < 3", `{}`, expectTrue},
	{"simple operator 12", "2 <= 3", `{}`, expectTrue},
	{"simple operator 13", "5 = 4", `{}`, expectFalse},
	{"simple operator 14", "5 != 4", `{}`, expectTrue},
	{"simple operator 15", "5 > 4", `{}`, expectTrue},
	{"simple operator 16", "5 >= 4", `{}`, expectTrue},
	{"simple operator 17", "5 < 4", `{}`, expectFalse},
	{"simple operator 18", "5 <= 4", `{}`, expectFalse},
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
	{"simple or 4", "1 > 2 OR 1 = 1", "{}", expectTrue},
	{"simple or 5", "1 > 2 OR 1 < 1", "{}", expectFalse},
	{"simple or 6", "1 < 2 OR 1 = 1", "{}", expectTrue},
	{"simple or 7", "true OR false", "{}", expectTrue},
	{"simple or 8", "false OR false", "{}", expectFalse},
	{"simple or 9", "(true) OR (false)", "{}", expectTrue},
	{"simple or 10", "(false) OR (false)", "{}", expectFalse},
	{"simple and 1", "(1 > 2) AND (1 = 1)", "{}", expectFalse},
	{"simple and 2", "(1 > 2) AND (1 < 1)", "{}", expectFalse},
	{"simple and 3", "(1 < 2) AND (1 = 1)", "{}", expectTrue},
	{"simple and 4", "1 > 2 AND 1 = 1", "{}", expectFalse},
	{"simple and 5", "1 > 2 AND 1 < 1", "{}", expectFalse},
	{"simple and 6", "1 < 2 AND 1 = 1", "{}", expectTrue},
	{"simple and 7", "true AND true", "{}", expectTrue},
	{"simple and 8", "false AND true", "{}", expectFalse},
	{"simple and 9", "(true) AND (true)", "{}", expectTrue},
	{"simple and 10", "(false) AND (true)", "{}", expectFalse},
	{"number 1", "Number('11') >= 12", `{}`, expectError},
	{"number 2", "Number('12') >= 12", `{}`, expectError},
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
	{"list literal 1", "[1 2 3 4] CONTAINS Number('2')", `{}`, expectError},
	{"list literal 2", "[1 2 3 4] CONTAINS Number('5')", `{}`, expectError},
	{"list literal 3", "['one' 'two' 'three'] CONTAINS 'two'", `{}`, expectTrue},
	{"list literal 4", "['one' 'two' 'three'] CONTAINS 'four'", `{}`, expectFalse},
	{"list literal 5", "[1 2 3 4] CONTAINS Number(var1)", `{"var1":1}`, expectTrue},
	{"list literal 6", "[1 2 3 4] CONTAINS Number(var1)", `{"var1":5}`, expectFalse},
	{"list literal 7", "[ 'one' 'two' 'three' ] CONTAINS 'two'", `{}`, expectTrue},
	{"list literal 8", "[ 'one' 'two' 'three' ] CONTAINS 'four'", `{}`, expectFalse},
	{"list literal 9", "[ 'one''two' 'three' ] CONTAINS 'four'", `{}`, expectError},
	{"list function empty", "List(var0) CONTAINS 2", `{"var0":[],"var1":1}`, expectFalse},
	{"list function 1", "List('[1 2 3 4]') CONTAINS Number(var1)", `{"var1":1}`, expectError},
	{"list function 2", "List(var0) CONTAINS Number(var1)", `{"var0":[1,2,3,4],"var1":1}`, expectTrue},
	{"list function 3", "List(var0) CONTAINS Number(var1)", `{"var0":[1,2,3,4],"var1":"1"}`, expectTrue},
	{"list function 4", "List(var0) CONTAINS var1", `{"var0":["alpha","beta",1,2],"var1":"beta"}`, expectTrue},
	{"list literal not", "[1 2 3 4] NOT CONTAINS Number(var1)", `{"var1":1}`, expectFalse},
	{"list function 1 not", "List('[1 2 3 4]') NOT CONTAINS Number(var1)", `{"var1":1}`, expectError},
	{"list function 2 not", "List(var0) NOT CONTAINS Number(var1)", `{"var0":[1,2,3,4],"var1":1}`, expectFalse},
	{"list function 3 not", "List(var0) NOT CONTAINS Number(var1)", `{"var0":[1,2,3,4],"var1":"1"}`, expectFalse},
	{"list function 4 not", "List(var0) NOT CONTAINS var1", `{"var0":["alpha","beta",1,2],"var1":"beta"}`, expectFalse},
	{"list with variable true", "['beta' 'alpha' 'has_reward_program' var1] CONTAINS 'blue'", `{"var1":"blue"}`, expectTrue},
	{"list with variable false", "['beta' 'alpha' 'has_reward_program' var1] CONTAINS 'canary'", `{"var1":"blue"}`, expectFalse},
	{"list numeric with numeric variable true", "[10 20 Number(var1)] CONTAINS 1", `{"var1":1}`, expectTrue},
	{"list numeric with numeric variable false", "[10 20 Number(var1)] CONTAINS 2", `{"var1":1}`, expectFalse},
	{"list numeric with string variable true", "[10 20 Number(var1)] CONTAINS 1", `{"var1":"1"}`, expectTrue},
	{"list numeric with string variable false", "[10 20 Number(var1)] CONTAINS 2", `{"var1":"1"}`, expectFalse},
	{"list with numeric variable true", "['beta' 'alpha' 'has_reward_program' var1] CONTAINS '1'", `{"var1":1}`, expectTrue},
	{"list with numeric variable false", "['beta' 'alpha' 'has_reward_program' var1] CONTAINS '2'", `{"var1":1}`, expectFalse},
	{"list with string variable true", "['beta' 'alpha' 'has_reward_program' var1] CONTAINS '1'", `{"var1":"1"}`, expectTrue},
	{"list with string variable false", "['beta' 'alpha' 'has_reward_program' var1] CONTAINS '2'", `{"var1":"1"}`, expectFalse},
	{"list contains variable true", "['beta' 'alpha' 'has_reward_program'] CONTAINS var1", `{"var1":"alpha"}`, expectTrue},
	{"list contains variable false", "['beta' 'alpha' 'has_reward_program'] CONTAINS var1", `{"var1":"gamma"}`, expectFalse},
	{"list contains variable enclosed all true", "(['beta' 'alpha' 'has_reward_program'] CONTAINS var1)", `{"var1":"alpha"}`, expectTrue},
	{"list contains variable enclosed all false", "(['beta' 'alpha' 'has_reward_program'] CONTAINS var1)", `{"var1":"gamma"}`, expectFalse},
	{"list contains variable enclosed list true", "(['beta' 'alpha' 'has_reward_program']) CONTAINS var1", `{"var1":"alpha"}`, expectTrue},
	{"list contains variable enclosed list false", "(['beta' 'alpha' 'has_reward_program']) CONTAINS var1", `{"var1":"gamma"}`, expectFalse},
	{"list contains variable enclosed var true", "['beta' 'alpha' 'has_reward_program'] CONTAINS (var1)", `{"var1":"alpha"}`, expectTrue},
	{"list contains variable enclosed var false", "['beta' 'alpha' 'has_reward_program'] CONTAINS (var1)", `{"var1":"gamma"}`, expectFalse},
	{"list contains variable enclosed list+all true", "((['beta' 'alpha' 'has_reward_program']) CONTAINS var1)", `{"var1":"alpha"}`, expectTrue},
	{"list contains variable enclosed list+all false", "((['beta' 'alpha' 'has_reward_program']) CONTAINS var1)", `{"var1":"gamma"}`, expectFalse},
	{"list contains variable enclosed list+all+var true", "((['beta' 'alpha' 'has_reward_program']) CONTAINS (var1))", `{"var1":"alpha"}`, expectTrue},
	{"list contains variable enclosed list+all+var false", "((['beta' 'alpha' 'has_reward_program']) CONTAINS (var1))", `{"var1":"gamma"}`, expectFalse},
	{"version 1", "000100010128 = Version(1.1.128)", `{}`, expectTrue},
	{"version 2", "000100010129 > Version(1.1.128)", `{}`, expectTrue},
	{"version 3", "000100010127 < Version(1.1.128)", `{}`, expectTrue},
	{"version 4", "000100020000 > Version(1.1.128)", `{}`, expectTrue},
	{"version 5", "000100000999 < Version(1.1.128)", `{}`, expectTrue},
	{"version 6", "Number(appVersion) = Version(1.1.128)", `{"appVersion":"000100010128"}`, expectTrue},
	{"version 7", "Number(appVersion) > Version(1.1.128)", `{"appVersion":"000100010129"}`, expectTrue},
	{"version 8", "Number(appVersion) < Version(1.1.128)", `{"appVersion":"000100010127"}`, expectTrue},
	{"version 9", "Number(appVersion) > Version(1.1.128)", `{"appVersion":"000100020000"}`, expectTrue},
	{"version 10", "Number(appVersion) < Version(1.1.128)", `{"appVersion":"000100000999"}`, expectTrue},
	{"not 1", "NOT NOT NOT NOT 1=1", "{}", expectTrue},
	{"not 2", "NOT NOT NOT NOT 1!=1", "{}", expectFalse},
	{"not 3", "NOT NOT NOT NOT (false AND true)", "{}", expectFalse},
	{"not 4", "NOT NOT NOT NOT (false OR true)", "{}", expectTrue},
	{"not 5", "(NOT NOT NOT NOT 1=1) AND (NOT NOT NOT NOT 1>2)", "{}", expectFalse},
	{"not 6", "(NOT NOT NOT NOT 1=1) AND (NOT NOT NOT NOT 1<2)", "{}", expectTrue},
	{"not 7", "NOT NOT NOT NOT 1=1 AND NOT NOT NOT NOT 1>2", "{}", expectFalse},
	{"not 8", "NOT NOT NOT NOT 1=1 AND NOT NOT NOT NOT 1<2", "{}", expectTrue},
}

func TestParser(t *testing.T) {
	scanTable(t, testTable, "builtin")
}

type parserTestCase struct {
	Name           string `json:"name"`
	Rule           string `json:"rule"`
	Vars           string `json:"vars"`
	ExpectedResult string `json:"expected_result"`
}

func TestSave(t *testing.T) {
	output := os.Getenv("TEST_SAVE")
	if output == "" {
		return
	}
	var table []parserTestCase
	for _, data := range testTable {
		tt := parserTestCase{
			Name: data.name,
			Rule: data.input,
			Vars: data.vars,
		}
		switch data.expectedResult {
		case expectTrue:
			tt.ExpectedResult = "true"
		case expectFalse:
			tt.ExpectedResult = "false"
		case expectError:
			tt.ExpectedResult = "error"
		}
		table = append(table, tt)
	}
	buf, _ := json.Marshal(table)
	if err := os.WriteFile(output, buf, 0777); err != nil {
		t.Errorf("write: %s: %v", output, err)
	}
}

func TestParserFromFile(t *testing.T) {
	const testDir = "tests"

	files, errDir := os.ReadDir(testDir)
	if errDir != nil {
		t.Errorf("list files: %s: %v", testDir, errDir)
		return
	}

	for _, f := range files {
		if f.Type().IsDir() {
			continue
		}
		filename := f.Name()
		if !strings.HasSuffix(filename, ".json") {
			continue
		}
		testFromFile(t, filepath.Join(testDir, filename))
	}
}

func testFromFile(t *testing.T, filename string) {

	t.Logf("test file: %s", filename)

	buf, errLoad := ioutil.ReadFile(filename)
	if errLoad != nil {
		t.Errorf("load tests error: %s: %v", filename, errLoad)
		return
	}

	var tab []parserTestCase

	errJSON := json.Unmarshal(buf, &tab)
	if errJSON != nil {
		t.Errorf("json error: %s: %v", filename, errJSON)
		return
	}

	t.Logf("loaded %d tests from file %s", len(tab), filename)

	var table []parserTest

	for _, item := range tab {
		tt := parserTest{
			name:  item.Name,
			input: item.Rule,
			vars:  item.Vars,
		}

		switch item.ExpectedResult {
		case "true":
			tt.expectedResult = expectTrue
		case "false":
			tt.expectedResult = expectFalse
		case "error":
			tt.expectedResult = expectError
		default:
			t.Errorf("%s: bad expected result from file: %s", item.Name, item.ExpectedResult)
		}
		table = append(table, tt)
	}

	scanTable(t, table, "fromFile")
}

func scanTable(t *testing.T, table []parserTest, label string) {

	for _, data := range table {

		var vars map[string]interface{}

		if data.vars != "" {
			if errJSON := json.Unmarshal([]byte(data.vars), &vars); errJSON != nil {
				t.Errorf("%s: json: %v: vars=%s", data.name, errJSON, data.vars)
			}
		}

		debug := false

		result := Run(bytes.NewBufferString(data.input), vars, debug)

		result2 := RunString(data.input, vars, debug)

		if result.IsError() != result2.IsError() || result.Eval != result2.Eval {
			t.Errorf("%s: Run.IsError=%v <=> RunString.IsError=%v, Run.Eval=%v <=> RunString.Eval=%v",
				data.name, result.IsError(), result2.IsError(), result.Eval, result2.Eval)
		}

		//t.Logf("%s %d/%d %s: rule='%s' vars='%s' vars_map='%v' result=%v", label, i, len(table), data.name, data.input, data.vars, vars, result)

		if data.expectedResult == expectError {
			// error expected
			if !result.IsError() {
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

func TestVarsListStringOnly(t *testing.T) {

	input := "List(userRoles) CONTAINS 'role2'"

	vars := map[string]interface{}{
		"userRoles": []string{"role1", "role2", "role3"},
	}

	const debug = false

	result := Run(bytes.NewBufferString(input), vars, debug)

	if result.IsError() {
		t.Errorf("unexpected error: %v", result)
	}

	if !result.Eval {
		t.Errorf("unexpected false evaluation")
	}
}

func TestVarsListNumberOnly(t *testing.T) {

	input := "List(userRoles) CONTAINS 2"

	vars := map[string]interface{}{
		"userRoles": []int{1, 2, 3},
	}

	const debug = false

	result := Run(bytes.NewBufferString(input), vars, debug)

	if result.IsError() {
		t.Errorf("unexpected error: %v", result)
	}

	if !result.Eval {
		t.Errorf("unexpected false evaluation")
	}
}

func TestVarsListMixed(t *testing.T) {

	input := "List(userRoles) CONTAINS 2"

	var list []interface{}
	list = append(list, "1")
	list = append(list, 2)
	list = append(list, "3")

	vars := map[string]interface{}{
		"userRoles": list,
	}

	const debug = false

	result := Run(bytes.NewBufferString(input), vars, debug)

	if result.IsError() {
		t.Errorf("unexpected error: %v", result)
	}

	if !result.Eval {
		t.Errorf("unexpected false evaluation")
	}
}

func TestVarsListWrongType(t *testing.T) {

	input := "List(booleans) CONTAINS [2]"

	vars := map[string]interface{}{
		"booleans": []bool{true, false, true}, // should be []interface{}
	}

	const debug = false

	result := Run(bytes.NewBufferString(input), vars, debug)

	if !result.IsError() {
		t.Errorf("expecting error for vars list but got success: %v", result)
	}
}

func TestVarsListWrongElemType(t *testing.T) {

	input := "List(userRoles) CONTAINS 2"

	var list []interface{}
	list = append(list, "1")
	list = append(list, 2)
	list = append(list, "3")
	list = append(list, true) // bool type is not supported

	vars := map[string]interface{}{
		"userRoles": list,
	}

	const debug = false

	result := Run(bytes.NewBufferString(input), vars, debug)

	if !result.IsError() {
		t.Errorf("expecting error for unsupported elem type but got success: %v", result)
	}
}

func TestVarNotFound(t *testing.T) {

	input := "List(userRolesNotFound) CONTAINS 2"

	var list []interface{}
	list = append(list, "1")
	list = append(list, 2)
	list = append(list, "3")
	list = append(list, true) // bool type is not supported

	vars := map[string]interface{}{
		"userRoles": list,
	}

	const debug = false

	result := Run(bytes.NewBufferString(input), vars, debug)

	if !result.IsError() {
		t.Errorf("expecting error for var not found but got success: %v", result)
	}
}

func TestBadNumberConversion(t *testing.T) {

	input := "Number(age) > 2"

	vars := map[string]interface{}{
		"age": "33333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333",
	}

	const debug = false

	result := Run(bytes.NewBufferString(input), vars, debug)

	if !result.IsError() {
		t.Errorf("expecting bad number conversion but got success: %v", result)
	}
}

func TestVarType(t *testing.T) {

	input := "int = '2' AND int64 = '3'"

	vars := map[string]interface{}{
		"int":   int(7),
		"int64": int64(7),
	}

	const debug = false

	result := Run(bytes.NewBufferString(input), vars, debug)

	if result.IsError() {
		t.Errorf("bad var type: %v", result)
	}
}

func TestVarTypeBool(t *testing.T) {

	input := "bool = '4'"

	vars := map[string]interface{}{
		"bool": true,
	}

	const debug = false

	result := Run(bytes.NewBufferString(input), vars, debug)

	if !result.IsError() {
		t.Errorf("expecting bad var type but got success: %v", result)
	}
}

func TestBadVersion1(t *testing.T) {

	input := "Version(3333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333.3.3) = 2"

	vars := map[string]interface{}{}

	const debug = false

	result := Run(bytes.NewBufferString(input), vars, debug)

	if !result.IsError() {
		t.Errorf("expecting bad version conversion but got success: %v", result)
	}
}

func TestBadVersion2(t *testing.T) {

	input := "Version(3.3333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333.3) = 2"

	vars := map[string]interface{}{}

	const debug = false

	result := Run(bytes.NewBufferString(input), vars, debug)

	if !result.IsError() {
		t.Errorf("expecting bad version conversion but got success: %v", result)
	}
}

func TestBadVersion3(t *testing.T) {

	input := "Version(3.3.3333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333) = 2"

	vars := map[string]interface{}{}

	const debug = false

	result := Run(bytes.NewBufferString(input), vars, debug)

	if !result.IsError() {
		t.Errorf("expecting bad version conversion but got success: %v", result)
	}
}

func TestBadVarNumberType(t *testing.T) {

	input := "Number(value) = 3"

	vars := map[string]interface{}{
		"value": true,
	}

	const debug = false

	result := Run(bytes.NewBufferString(input), vars, debug)

	if !result.IsError() {
		t.Errorf("expecting bad variable number conversion but got success: %v", result)
	}
}

func TestVarNumberType(t *testing.T) {

	input := "Number(value) = 3 AND Number(age) = 2"

	vars := map[string]interface{}{
		"value": int(3),
		"age":   int64(4),
	}

	const debug = false

	result := Run(bytes.NewBufferString(input), vars, debug)

	if result.IsError() {
		t.Errorf("bad variable number type: %v", result)
	}
}

func TestBadNumber(t *testing.T) {

	input := "2 < 33333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333"

	vars := map[string]interface{}{}

	const debug = false

	result := Run(bytes.NewBufferString(input), vars, debug)

	if !result.IsError() {
		t.Errorf("expecting bad number conversion but got success: %v", result)
	}
}

func TestDebug(t *testing.T) {
	if result := RunString("true", nil, true); result.IsError() || !result.Eval {
		t.Errorf("unexpected false result")
	}
	if result := RunString("v>1", nil, true); !result.IsError() {
		t.Errorf("unexpected non-error")
	}
}
