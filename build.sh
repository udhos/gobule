#!/bin/bash

install_nex() {
    echo installing nex
    pushd $PWD >/dev/null
    cd /tmp
    go get github.com/blynn/nex
    popd >/dev/null
}

install_goyacc() {
    echo installing goyacc
    pushd $PWD >/dev/null
    cd /tmp
    go get modernc.org/goyacc ;# supports %precedence
    popd >/dev/null
}

which nex 2>/dev/null || install_nex

echo generating lexer

nex -s < lexer/lexer.nex > lexer/lexer.go

go install ./lexer-run

expression="false != [1 a b 'Teste' CurrentTime] CONTAINS Number(name)"

echo running lexer test: $expression

echo "$expression" | lexer-run

which goyacc 2>/dev/null || install_goyacc

go generate ./parser ;# see ./parser/generate.go

go install ./parser-run