#!/bin/bash
if [ ! -f "3rd_party/go-test-report/go-test-report" ]; then
    pushd 3rd_party/go-test-report
    go build
    popd
fi

go test -json ./... | 3rd_party/go-test-report/go-test-report -g 10 -s 40 -t "Ferrormarket backend tests"

go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o test_coverage.html
