# go_test_workshop
Go Test Workshop


- s1: go test basic
  - [table driven tests ](https://github.com/golang/go/wiki/TableDrivenTests)
  - [color output](https://github.com/rakyll/gotest)
    - gotest -v ./...

- s2: testing.T
  - testing.TB interface
  - testing.T
    - Fail/FailNow/Failed/Fatal/Fatalf
    - Log/Logf/Error/Errorf
    - Skip/SkipNow/Skipf/Skipped
    - Helper


- s3: go test command
  - go help test
  - go help testflag
  - https://pkg.go.dev/testing

- s4: coverage
  - go test -coverprofile cover.out ./...
  - go-cover-treemap -coverprofile cover.out > out.svg
  - https://github.com/microsoft/vscode-go/issues/816
  - https://golang.org/doc/go1.10#test
  - go tool cover -html=cover.out -o cover.tml
  - go tool cover -o coverage.html -html=coverage.out; sed -i 's/black/whitesmoke/g' coverage.html; sensible-browser coverage.html
  - mac: go tool cover -o coverage.html -html=coverage.out; sed -i'*.bak' 's/black/whitesmoke/g' coverage.html; open coverage.html


- s4: tools
- assert/require
- gocovery

- s5: mock

- s6: cases
 - http
 - sql
 - integration test

- s7: benchmark

- s8: example




## install tools

```
go install github.com/rakyll/gotest@master
go install github.com/axw/gocov/gocov@latest
go install github.com/matm/gocov-html@latest
go install github.com/nikolaydubina/go-cover-treemap@latest
go get github.com/smartystreets/goconvey/convey
go get github.com/stretchr/testify
```