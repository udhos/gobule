
%{

package parser

// header of parser.y 

import (
        "github.com/udhos/gobule/exp"
)

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.

%union{
    typeExp exp.Exp
	tok bool
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct

%type <typeExp> exp

// same for terminals

%token <tok> TkTrue

%%

prog: exp { /* prog */ }

exp: TkTrue { $$ = $1 }