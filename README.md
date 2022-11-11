Cvt
---

## Intro

convert json, toml, yaml ..., and quota/unquota it

## Build

```
go build -o bin/cvt cmd/cvt/main.go
```

## Install

```
go get github.com/chg1f/cvt
go install github.com/chg1f/cvt
```

## Usage

```bash
$ echo '{"a":1.0,"b":["c",2],"d":{"e":"f"}}' | bin/cvt -t yaml
a: 1
b:
    - c
    - 2
d:
    e: f
$ echo '{"a":1.0,"b":["c",2],"d":{"e":"f"}}' | bin/cvt -t yaml | bin/cvt -f yaml -t toml
a = 1
b = ["c", 2]

[d]
  e = "f"
$ echo '{"a":1.0,"b":["c",2],"d":{"e":"f"}}' | bin/cvt -t yaml | bin/cvt -f yaml -t toml | bin/cvt -f toml -t json | jq
{
  "a": 1,
  "b": [
    "c",
    2
  ],
  "d": {
    "e": "f"
  }
}
$ echo '{"a":1.0,"b":["c",2],"d":{"e":"f"}}' | bin/cvt -t yaml | bin/cvt -f yaml -t toml | bin/cvt -f toml -t json | jq -rc | bin/cvt -q
"{\"a\":1,\"b\":[\"c\",2],\"d\":{\"e\":\"f\"}}\n"
$ echo '{"a":1.0,"b":["c",2],"d":{"e":"f"}}' | bin/cvt -t yaml | bin/cvt -f yaml -t toml | bin/cvt -f toml -t json | jq -rc | bin/cvt -q | bin/cvt -u
{"a":1,"b":["c",2],"d":{"e":"f"}}
```

## Bugs

```
$ make test
go test -v ./...
=== RUN   TestConvert
    main_test.go:18:
                Error Trace:    github.com/chg1f/cvt/cmd/cvt/main_test.go:18
                Error:          Not equal:
                                expected: "a: 1.0\nb:\n    - c\n    - 2\nd:\n    e: f\n"
                                actual  : "a: 1\nb:\n    - c\n    - 2\nd:\n    e: f\n"

                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1,2 +1,2 @@
                                -a: 1.0
                                +a: 1
                                 b:
                Test:           TestConvert
                Messages:       from json to yaml
    main_test.go:18:
                Error Trace:    github.com/chg1f/cvt/cmd/cvt/main_test.go:18
                Error:          Not equal:
                                expected: "a = 1.0\nb = [\"c\", 2]\n\n[d]\n  e = \"f\"\n"
                                actual  : "a = 1.0\nb = [\"c\", 2.0]\n\n[d]\n  e = \"f\"\n"

                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1,3 +1,3 @@
                                 a = 1.0
                                -b = ["c", 2]
                                +b = ["c", 2.0]

                Test:           TestConvert
                Messages:       from json to toml
    main_test.go:18:
                Error Trace:    github.com/chg1f/cvt/cmd/cvt/main_test.go:18
                Error:          Not equal:
                                expected: "{\"a\":1.0,\"b\":[\"c\",2],\"d\":{\"e\":\"f\"}}"
                                actual  : "{\"a\":1,\"b\":[\"c\",2],\"d\":{\"e\":\"f\"}}"

                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1 +1 @@
                                -{"a":1.0,"b":["c",2],"d":{"e":"f"}}
                                +{"a":1,"b":["c",2],"d":{"e":"f"}}
                Test:           TestConvert
                Messages:       from yaml to json
    main_test.go:18:
                Error Trace:    github.com/chg1f/cvt/cmd/cvt/main_test.go:18
                Error:          Not equal:
                                expected: "{\"a\":1.0,\"b\":[\"c\",2],\"d\":{\"e\":\"f\"}}"
                                actual  : "{\"a\":1,\"b\":[\"c\",2],\"d\":{\"e\":\"f\"}}"

                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1 +1 @@
                                -{"a":1.0,"b":["c",2],"d":{"e":"f"}}
                                +{"a":1,"b":["c",2],"d":{"e":"f"}}
                Test:           TestConvert
                Messages:       from toml to json
    main_test.go:18:
                Error Trace:    github.com/chg1f/cvt/cmd/cvt/main_test.go:18
                Error:          Not equal:
                                expected: "a: 1.0\nb:\n    - c\n    - 2\nd:\n    e: f\n"
                                actual  : "a: 1\nb:\n    - c\n    - 2\nd:\n    e: f\n"

                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1,2 +1,2 @@
                                -a: 1.0
                                +a: 1
                                 b:
                Test:           TestConvert
                Messages:       from toml to yaml
--- FAIL: TestConvert (0.00s)
=== RUN   TestConvertNil
--- PASS: TestConvertNil (0.00s)
=== RUN   TestQuota
--- PASS: TestQuota (0.00s)
FAIL
FAIL    github.com/chg1f/cvt/cmd/cvt    0.002s
FAIL
```
