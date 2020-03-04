#!/bin/bash
FormatCheck=$(gofmt -l *.go | wc -l)
if [ $FormatCheck -gt 0 ]; then
    gofmt -l *.go
    echo "gofmt -w *.go your code please."
    exit 1
fi
go test -v
