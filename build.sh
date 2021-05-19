#!/bin/bash

install_goyacc() {
    echo installing goyacc
    pushd $PWD >/dev/null
    cd /tmp
    go get modernc.org/goyacc ;# supports %precedence
    popd >/dev/null
}

#
# Test scanner
#
go test ./bulexer

#
# Build standalone scanner
#
go install ./bulexer-run

#
# Build parser
#
which goyacc 2>/dev/null || install_goyacc
rm -f parser/parser.go
go generate -v -x ./parser ;# see ./parser/generate.go
#[ -f parser/parser.go ] || goyacc -o parser.go parser.y
go install ./parser-run

