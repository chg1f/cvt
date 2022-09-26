CVT
---

# Intro

convert json, toml, yaml ...

# Build

```
go build -o bin/cvt cmd/cvt/main.go
```

# Usage

```bash
$ echo '{"key": "value"}' | cvt -f json -t yaml
key: value
```
