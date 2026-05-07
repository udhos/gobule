#!/bin/bash

install_goyacc() {
    echo installing goyacc
    pushd $PWD >/dev/null
    cd /tmp
    go install modernc.org/goyacc@v1.0.3 ;# supports %precedence
    popd >/dev/null
}

go install golang.org/x/vuln/cmd/govulncheck@latest
go install golang.org/x/tools/cmd/deadcode@latest
go install github.com/mgechev/revive@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
go install golang.org/x/tools/go/analysis/passes/modernize/cmd/modernize@latest
go install github.com/gordonklaus/ineffassign@latest
go install github.com/client9/misspell/cmd/misspell@latest
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest

gofmt -s -w .

revive ./...

staticcheck ./...

modernize -fix ./...

gocyclo -over 15 .

ineffassign ./...

misspell .

go mod tidy

govulncheck ./...

#deadcode ./cmd/*

#
# Generate parser
#
which goyacc 2>/dev/null || install_goyacc
rm -f parser/parser.go
go generate -v -x ./parser ;# see ./parser/generate.go
#[ -f parser/parser.go ] || goyacc -o parser.go parser.y

go env -w CGO_ENABLED=1

go test -race ./...

go test -race -covermode=atomic -coverprofile=coverage.txt ./...

go env -w CGO_ENABLED=0

go install ./...

go env -u CGO_ENABLED
