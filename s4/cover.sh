#!/bin/bash

cd ..
go test -coverprofile s4/cover.out ./...

gocov test ./... | gocov-html > s4/cover3.html

cd s4

go tool cover -html=cover.out -o cover.html

go tool cover -o cover2.html -html=cover.out; sed -i'*.bak' 's/black/whitesmoke/g' cover2.html; rm -fr cover2.html*.bak
