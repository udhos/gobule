#!/bin/bash

install_goyacc() {
    echo installing goyacc
    pushd $PWD >/dev/null
    cd /tmp
    go install modernc.org/goyacc@v1.0.3 ;# supports %precedence
    popd >/dev/null
}

# build() {
# 	local pkg="$1"

# 	gofmt -s -w "$pkg"
# 	go fix "$pkg"
# 	go vet "$pkg"

# 	hash golint 2>/dev/null && golint "$pkg"
# 	hash staticcheck 2>/dev/null && staticcheck "$pkg"

# 	#go test -failfast "$pkg"

# 	go install -v "$pkg"
# }

#go test -race -covermode=atomic -coverprofile=coverage.txt ./...

#build ./conv
#build ./bulexer
#build ./bulexer-run

#
# Generate parser
#
which goyacc 2>/dev/null || install_goyacc
rm -f parser/parser.go
go generate -v -x ./parser ;# see ./parser/generate.go
#[ -f parser/parser.go ] || goyacc -o parser.go parser.y

gofmt -s -w .

revive ./...

go install ./...

go test -race ./...

go test -race -covermode=atomic -coverprofile=coverage.txt ./...

#build ./parser
#build ./parser-run
