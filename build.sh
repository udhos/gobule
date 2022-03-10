#!/bin/bash

install_goyacc() {
    echo installing goyacc
    pushd $PWD >/dev/null
    cd /tmp
    go get modernc.org/goyacc ;# supports %precedence
    popd >/dev/null
}

build() {
	local pkg="$1"

	gofmt -s -w "$pkg"
	go fix "$pkg"
	go vet "$pkg"

	hash golint >/dev/null && golint "$pkg"
	hash staticcheck >/dev/null && staticcheck "$pkg"

	go test -failfast "$pkg"

	go install -v "$pkg"
}

build ./bulexer
build ./bulexer-run
build ./conv

#
# Generate parser
#
which goyacc 2>/dev/null || install_goyacc
rm -f parser/parser.go
go generate -v -x ./parser ;# see ./parser/generate.go
#[ -f parser/parser.go ] || goyacc -o parser.go parser.y

build ./parser
build ./parser-run
build ./conv
