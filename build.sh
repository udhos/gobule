#!/bin/bash

#install_nex() {
#    echo installing nex
#    pushd $PWD >/dev/null
#    cd /tmp
#    go get github.com/blynn/nex
#    popd >/dev/null
#}

install_goyacc() {
    echo installing goyacc
    pushd $PWD >/dev/null
    cd /tmp
    go get modernc.org/goyacc ;# supports %precedence
    popd >/dev/null
}

expression="false != [1 a b 'Teste' CurrentTime] CONTAINS Number(name)"

#which nex 2>/dev/null || install_nex
#
#echo generating lexer
#
#rm lexer/lexer.go
#go generate -v -x ./lexer ;# see ./lexer/generate.go
#[ -f lexer/lexer.go ] || nex -s < lexer/lexer.nex > lexer/lexer.go
#
#go install ./lexer-run
#
#echo running lexer test: $expression
#
#echo "$expression" | lexer-run

go install ./bulexer-run
echo "$expression" | bulexer-run

which goyacc 2>/dev/null || install_goyacc
rm parser/parser.go
go generate -v -x ./parser ;# see ./parser/generate.go
#[ -f parser/parser.go ] || goyacc -o parser.go parser.y
go install ./parser-run

