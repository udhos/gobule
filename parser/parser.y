
%{

package parser

// header of parser.y 

import (
        "fmt"
        //"encoding/json"
        //"log"
        "strconv"
        "time"

        "github.com/udhos/gobule/conv"
)

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.

%union{
    typeBool bool
    typeString string // holds: variable, number, or text
    typeScalar scalar
    typeList []scalar
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct

%type <typeBool> bool_exp
%type <typeScalar> scalar_exp
//%type <typeString> scalar_text
%type <typeList> list_exp
%type <typeList> list

// same for terminals

%token <typeBool> TkKeywordTrue
%token <typeBool> TkKeywordFalse
%precedence <typeBool> TkKeywordOr
%precedence <typeBool> TkKeywordAnd
%token <typeBool> TkKeywordNot
%token <typeBool> TkKeywordContains
%token <typeBool> TkKeywordCurrentTime
%token <typeBool> TkKeywordNumber
%token <typeBool> TkKeywordList
%token <typeBool> TkKeywordVersion
%token <typeString> TkNumber
%token <typeString> TkText
%token <typeString> TkIdent
%token <typeBool> TkParL
%token <typeBool> TkParR
%token <typeBool> TkSBktL
%token <typeBool> TkSBktR
%token <typeBool> TkEQ
%token <typeBool> TkLT
%token <typeBool> TkGT
%token <typeBool> TkNE
%token <typeBool> TkGE
%token <typeBool> TkLE
%token <typeBool> TkDot

%%

prog:
    bool_exp { yylex.(*Lex).result.Eval = $1 }

bool_exp:
    TkParL bool_exp TkParR { $$ = $2 }
    | bool_exp TkKeywordAnd bool_exp { $$ = $1 && $3 }
    | bool_exp TkKeywordOr bool_exp { $$ = $1 || $3 }
    | TkKeywordNot bool_exp { $$ = !$2 }
    | TkKeywordTrue { $$ = true }
    | TkKeywordFalse { $$ = false }
    | list_exp TkKeywordContains scalar_exp { $$ = contains($1, $3) }
    | list_exp TkKeywordNot TkKeywordContains scalar_exp { $$ = !contains($1, $4) }
    | scalar_exp TkEQ scalar_exp { $$ = $1.equals($3) }
    | scalar_exp TkNE scalar_exp { $$ = !$1.equals($3) }
    | scalar_exp TkGT scalar_exp
        {
            var eval bool
            if $1.scalarType != $3.scalarType {
                yylex.Error("greater-than operator for different types")
            } else {
                eval = $1.greaterThan($3)
            }
            $$ = eval
        }
    | scalar_exp TkGE scalar_exp
        {
            var eval bool
            if $1.scalarType != $3.scalarType {
                yylex.Error("greater-than-or-equal operator for different types")
            } else {
                eval = !$3.greaterThan($1)
            }
            $$ = eval
        }
    | scalar_exp TkLT scalar_exp
        {
            var eval bool
            if $1.scalarType != $3.scalarType {
                yylex.Error("less-than operator for different types")
            } else {
                eval = $3.greaterThan($1)
            }
            $$ = eval
        }
    | scalar_exp TkLE scalar_exp
        {
            var eval bool
            if $1.scalarType != $3.scalarType {
                yylex.Error("less-than-or-equal operator for different types")
            } else {
                eval = !$1.greaterThan($3)
            }
            $$ = eval
        }

list_exp:
    TkSBktL TkSBktR { $$ = []scalar{} }
    |
    TkSBktL list TkSBktR { $$ = $2 }
    |
    TkKeywordList TkParL TkText TkParR
        {
            var list []scalar
            {
                s := $3
                strList, errParse := parseList(s)
                if errParse != nil {
                    yylex.Error(fmt.Sprintf("List(%s): bad list: %v", s, errParse))
                }
                for i, elem := range strList {
                    switch v := elem.(type) {
                    //case int:
                    //    list = append(list, scalar{scalarType: scalarNumber, number: int64(v)})
                    case int64:
                        list = append(list, scalar{scalarType: scalarNumber, number: v})
                    case string:
                        list = append(list, scalar{scalarType: scalarText, text: v})
                    default:
                        yylex.Error(fmt.Sprintf("List(%s): invalid type for element %d: %v", s, i, elem))
                    }
                }
            }
            $$ = list
        }
    |
    TkKeywordList TkParL TkIdent TkParR
        {
            var list []scalar
            v := $3
            l := yylex.(*Lex)
            if varValue, found := l.vars[v]; found {
                // found variable

                switch vv := varValue.(type) {
                    case []interface{}:
                        for i, elem := range vv {
                            switch val := elem.(type) {
                            case float64:
                                list = append(list, scalar{scalarType: scalarNumber, number: int64(val)})
                            case int:
                                list = append(list, scalar{scalarType: scalarNumber, number: int64(val)})
                            case string:
                                list = append(list, scalar{scalarType: scalarText, text: val})
                            default:
                                yylex.Error(fmt.Sprintf("List(%s): invalid type for element %d: %v", v, i, elem))
                            }
                        }
                    case []int:
                        for _, elem := range vv {
                            list = append(list, scalar{scalarType: scalarNumber, number: int64(elem)})
                        }
                    case []string:
                        for _, elem := range vv {
                            list = append(list, scalar{scalarType: scalarText, text: elem})
                        }
                    default:
                        yylex.Error(fmt.Sprintf("List(%s): unexpected list type (%T): %v", v, varValue, varValue))
                }

            } else {
                yylex.Error(fmt.Sprintf("List(%s): variable undefined", v))
            }
            $$ = list
        }


list:
    scalar_exp
    {
        l := yylex.(*Lex)
        l.scalarList = []scalar{$1}
        $$ = l.scalarList
    }
    | list scalar_exp
    {
        l := yylex.(*Lex)
        l.scalarList = append(l.scalarList, $2)
        $$ = l.scalarList
    }

scalar_exp:
    TkText { $$ = scalar{scalarType: scalarText, text: $1} }
    | TkNumber
        {
            s := $1
            n, errConv := parseInt(s)
            if errConv != nil {
                yylex.Error(fmt.Sprintf("bad number conversion: '%s': %v", s, errConv))
            }
            $$ = scalar{scalarType: scalarNumber, number: n}
        }
    | TkIdent
        {
            v := $1
            l := yylex.(*Lex)
            value := scalar{scalarType: scalarText}
            if varValue, found := l.vars[v]; found {
                switch val := varValue.(type) {
                case string:
                    value.text = val
                case int:
                    value.text = strconv.Itoa(val)
                case int64:
                    value.text = strconv.FormatInt(val, 10)
                case float64:
                    value.text = strconv.FormatInt(int64(val), 10)
                default:
                    yylex.Error(fmt.Sprintf("unexpected type='%T' for variable='%s' value='%v'", varValue, v, varValue))
                }
            } else {
                value.text = fmt.Sprintf("variable undefined:'%s'", v)
                yylex.Error(value.text)
            }
            $$ = value
        }
    | TkKeywordVersion TkParL TkNumber TkDot TkNumber TkDot TkNumber TkParR
        {
            s1 := $3
            s2 := $5
            s3 := $7

            v1, errConv1 := parseInt(s1)
            if errConv1 != nil {
                yylex.Error(fmt.Sprintf("bad Version(version) number conversion 1: '%s': %v", s1, errConv1))
            }
            v2, errConv2 := parseInt(s2)
            if errConv2 != nil {
                yylex.Error(fmt.Sprintf("bad Version(version) number conversion 2: '%s': %v", s2, errConv2))
            }
            v3, errConv3 := parseInt(s3)
            if errConv3 != nil {
                yylex.Error(fmt.Sprintf("bad Version(version) number conversion 3: '%s': %v", s3, errConv3))
            }

            $$ = scalar{scalarType: scalarNumber, number: conv.VersionToNumber(v1, v2, v3) }
        }
    | TkKeywordNumber TkParL TkIdent TkParR
        {
            v := $3
            l := yylex.(*Lex)
            value := scalar{scalarType: scalarNumber}
            if varValue, found := l.vars[v]; found {
                // found variable
                switch val := varValue.(type) {
                case string:
                    n, errConv := parseInt(val)
                    if errConv != nil {
                        yylex.Error(fmt.Sprintf("bad Number(variable) conversion: %s='%s': %v", v, val, errConv))
                    }
                    value.number = n
                case int:
                    value.number = int64(val)
                case float64:
                    value.number = int64(val)
                default:
                    yylex.Error(fmt.Sprintf("unexpected Number(variable) var type: '%s': %q", v, varValue))
                }
            } else {
                value.text = fmt.Sprintf("Number() variable undefined:'%s'", v)
                yylex.Error(value.text)
            }
            $$ = value
        }
    | TkKeywordCurrentTime TkParL TkParR
        {
            now := time.Now()
            n := now.Hour() * 10000 + now.Minute() * 100 + now.Second()
            //fmt.Printf("currenttime: %d\n", n)
            $$ = scalar{scalarType: scalarNumber, number: int64(n)}
        }

