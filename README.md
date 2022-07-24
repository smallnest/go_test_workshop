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
    - Parallel


- s3: go test command
  - go help test
    - go test [build/test flags] [packages] [build/test flags & test binary flags]
    - local directory mode: cache disabled
    - package list mode: cache enabled. `-count=1` to disable cache
    - test binary flags
      - -c:  Compile the test binary to pkg.test but do not run it
      - -exec xprog
      - -json: Convert test output to JSON suitable for automated processing.
      - -o file:Compile the test binary to the named file.

  - go help testflag
      - -count n: Run each test, benchmark, and fuzz seed n times (default 1).
      - -cpu 1,2,4: Specify a list of GOMAXPROCS values for which the tests, benchmarks or fuzz tests should be executed.
      - -failfast: Do not start new tests after the first test failure.
      - -json
      - -list regexp
      - -parallel n
      - -run regexp
      - -short
      - -shuffle off,on,N
      - -timeout d
      - -v
      - -vet list


  - go help testfunc
    - func TestXxx(t *testing.T) { ... }
    - func BenchmarkXxx(b *testing.B) { ... }
    - func FuzzXxx(f *testing.F) { ... }
    - func ExamplePrintln() { ... }

  - https://pkg.go.dev/testing

- s4: coverage
  - go test -coverprofile cover.out ./...
  - go-cover-treemap -coverprofile cover.out > out.svg
  - https://github.com/microsoft/vscode-go/issues/816
  - https://golang.org/doc/go1.10#test
  - go tool cover -html=cover.out -o cover.html
  - go tool cover -o cover2.html -html=cover.out; sed -i 's/black/whitesmoke/g' cover2.html; open cover2.html
  - mac: go tool cover -o cover2.html -html=cover.out; sed -i'*.bak' 's/black/whitesmoke/g' cover2.html; rm -fr cover2.html*.bak;open cover2.html


- s5: tools
  - assert/require
  - gocovery: Behavior-driven Development

- s6: mock
  - https://github.com/golang/mock
  - https://github.com/vektra/mockery
  - https://github.com/DATA-DOG/go-sqlmock
  - https://github.com/h2non/gock
  - https://github.com/gavv/httpexpect

  - monkey
    - https://github.com/bouk/monkey
    - https://github.com/agiledragon/gomonkey
      - printf '\x07' | dd of=api.test bs=1 seek=160 count=1 conv=notrunc
      - https://github.com/eisenxp/macos-golink-wrapper

- s7: web case
 - dao
 - service
 - http

- s8: benchmark

- s9: example

- s10: fuzzing test



## install tools

```
go install github.com/rakyll/gotest@master
go install github.com/axw/gocov/gocov@latest
go install github.com/matm/gocov-html@latest
go install github.com/nikolaydubina/go-cover-treemap@latest
go get github.com/smartystreets/goconvey/convey
go get github.com/stretchr/testify
```