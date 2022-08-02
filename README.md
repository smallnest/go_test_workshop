# go_test_workshop

如何编写健壮的代码？测试或者单元测试是一个非常有效的手段。自1997年JUint的出世引领单元测试以来，单元测试曾经风靡一时。然后在最近的十年，或者说在最近的几年，互联网企业面临巨大的竞争压力，或者说代码规范不收重视的情况下，糙快猛的编程方式也收到吹捧，以至于单元测试没有收到足够的重视，因此留下来很多的技术债务和线上风险，也就是俗称的"雷"。

这个实战课程，全面介绍Go语言的单元测试的知识，通过大量的实战的代码，带你领略go test的便捷和强大，解答你在实际开发中的众多的疑惑，你在以后的Go应用程序的开发中，可以应用本文介绍的知识，把单元测试威力实际发挥出来。

如果你会写`func TestXXX(t *testing.T) {`这一行代码，恭喜你，你已经知道了Go的一半的单元测试的知识，你可以重点关注其它的一半的Go单元测试的知识。如果你还未写过Go的单元测试，没关系，本门课程手把手教你怎么轻轻松松编写Go的单元测试。


本课程的代码放在了 [Go Test Workshop](https://github.com/smallnest/go_test_workshop)

![](/s0/bfe.svg)

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
  - https://github.com/steinfletcher/apitest
  - https://github.com/carlmjohnson/be
  - https://github.com/arikama/go-mysql-test-container
  - https://github.com/jfilipczyk/gomatch

  - monkey
    - https://github.com/bouk/monkey
    - https://github.com/agiledragon/gomonkey
      - printf '\x07' | dd of=api.test bs=1 seek=160 count=1 conv=notrunc
      - https://github.com/eisenxp/macos-golink-wrapper
  - container
    - https://github.com/testcontainers/testcontainers-go

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
// https://go-cover-treemap.io/
```