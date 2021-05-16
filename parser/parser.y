
%{

package parser

// header of parser.y 

import (
        "fmt"
        "strconv"
)

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.

%union{
    typeBool bool
    typeText string
    typeNumber string
    typeValue nodeValue
    typeList []nodeValue
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct

%type <typeBool> bool_exp
%type <typeValue> value
%type <typeList> list_exp
%type <typeList> list

// same for terminals

%token <typeBool> TkKeywordTrue
%token <typeBool> TkKeywordFalse
%token <typeBool> TkKeywordAnd
%token <typeBool> TkKeywordOr
%token <typeBool> TkKeywordNot
%token <typeBool> TkKeywordContains
%token <typeBool> TkKeywordCurrentTime
%token <typeBool> TkKeywordNumber
%token <typeBool> TkKeywordList
%token <typeNumber> TkNumber
%token <typeText> TkText
%token <typeBool> TkIdent
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

prog: bool_exp { result.Eval = $1 }

bool_exp: TkParL bool_exp TkParR { $$ = $2 }
    | bool_exp TkKeywordAnd bool_exp { $$ = $1 && $3 }
    | bool_exp TkKeywordOr bool_exp { $$ = $1 || $3 }
    | TkKeywordNot bool_exp { $$ = !$2 }
    | TkKeywordTrue { $$ = true }
    | TkKeywordFalse { $$ = false }
    | list_exp TkKeywordContains value
        {
            list := $1
            wanted := $3
            $$ = contains(list, wanted)
        }
    | list_exp TkKeywordNot TkKeywordContains value
        {
            list := $1
            wanted := $4
            $$ = !contains(list, wanted)
        }

list_exp: TkSBktL list TkSBktR { $$ = $2 }

value: TkText { $$ = nodeValue{nodeType: valueText, text: $1} }
    | TkNumber
        {
            s := $1
            n, errConv := strconv.Atoi(s)
            if errConv != nil {
                yylex.Error(fmt.Sprintf("bad number conversion: '%s': %v", s, errConv))
            }
            $$ = nodeValue{nodeType: valueNumber, number: n}
        }

list: /* empty */
    {
        valueList = []nodeValue{}
        $$ = valueList
    }
    | value
    {
        valueList = []nodeValue{$1}
        $$ = valueList
    }
    | list value
    {
        valueList = append(valueList, $2)
        $$ = valueList
    }
