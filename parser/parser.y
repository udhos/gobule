
%{

package parser

// header of parser.y 

import (
        "fmt"
        //"log"
        "strconv"
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

list_exp:
    TkSBktL TkSBktR { $$ = []scalar{} }
    |
    TkSBktL list TkSBktR { $$ = $2 }

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
                value.text = varValue
            } else {
                value.text = fmt.Sprintf("variable undefined:'%s'", v)
                yylex.Error(value.text)
            }
            $$ = value
        }

