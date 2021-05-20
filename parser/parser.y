
%{

package parser

// header of parser.y 

import (
        "fmt"
        //"encoding/json"
        //"log"
        "strconv"
        "time"
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
    | scalar_exp TkEQ scalar_exp { $$ = $1.Equals($3) }
    | scalar_exp TkNE scalar_exp { $$ = !$1.Equals($3) }
    | scalar_exp TkGT scalar_exp
        {
            var eval bool
            if $1.scalarType != $3.scalarType {
                yylex.Error("greater-than operator for different types")
            } else {
                eval = $1.GreaterThan($3)
            }
            $$ = eval
        }
    | scalar_exp TkGE scalar_exp
        {
            var eval bool
            if $1.scalarType != $3.scalarType {
                yylex.Error("greater-than-or-equal operator for different types")
            } else {
                eval = $1.GreaterThanOrEqual($3)
            }
            $$ = eval
        }
    | scalar_exp TkLT scalar_exp
        {
            var eval bool
            if $1.scalarType != $3.scalarType {
                yylex.Error("less-than operator for different types")
            } else {
                eval = $3.GreaterThan($1)
            }
            $$ = eval
        }
    | scalar_exp TkLE scalar_exp
        {
            var eval bool
            if $1.scalarType != $3.scalarType {
                yylex.Error("less-than-or-equal operator for different types")
            } else {
                eval = $3.GreaterThanOrEqual($1)
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
                    case int:
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

                if vv, ok := varValue.([]interface{}); ok {

                    for i, elem := range vv {
                        switch val := elem.(type) {
                        case float64:
                            list = append(list, scalar{scalarType: scalarNumber, number: int(val)})
                        case int:
                            list = append(list, scalar{scalarType: scalarNumber, number: val})
                        case string:
                            list = append(list, scalar{scalarType: scalarText, text: val})
                        default:
                            yylex.Error(fmt.Sprintf("List(%s): invalid type for element %d: %v", v, i, elem))
                    }
                }

                } else {
                    yylex.Error(fmt.Sprintf("List(%s): unexpected list type: %v", v, varValue))
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
            n, errConv := strconv.Atoi(s)
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
                default:
                    yylex.Error(fmt.Sprintf("unexpected variable type: '%s'", v))
                }
            } else {
                value.text = fmt.Sprintf("variable undefined:'%s'", v)
                yylex.Error(value.text)
            }
            $$ = value
        }
    | TkKeywordNumber TkParL TkText TkParR
        {
            s := $3
            n, errConv := strconv.Atoi(s)
            if errConv != nil {
                yylex.Error(fmt.Sprintf("bad Number(text) conversion: '%s': %v", s, errConv))
            }
            $$ = scalar{scalarType: scalarNumber, number: n}
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
                    n, errConv := strconv.Atoi(val)
                    if errConv != nil {
                        yylex.Error(fmt.Sprintf("bad Number(variable) conversion: %s='%s': %v", v, val, errConv))
                    }
                    value.number = n
                case int:
                    value.number = val
                case float64:
                    value.number = int(val)
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
            $$ = scalar{scalarType: scalarNumber, number: n}
        }

