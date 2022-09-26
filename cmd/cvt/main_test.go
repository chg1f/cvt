package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	buf := map[string]*bytes.Buffer{
		JSON: bytes.NewBufferString("{\"a\":1,\"b\":{\"c\":\"d\"},\"e\":[\"f\",2]}"),
		YAML: bytes.NewBufferString("a: 1\nb:\n  c: d\ne:\n- f\n- 2\n"),
		TOML: bytes.NewBufferString("a = 1.0\ne = [\"f\", 2.0]\n\n[b]\n  c = \"d\"\n"),
	}
	for from := range buf {
		for to := range buf {
			assert.Equal(t, buf[to].String(), convert(buf[from], from, to, false), "from "+from+" to "+to)
		}
	}
}
func TestConvertNil(t *testing.T) {
	buf := map[string]*bytes.Buffer{
		JSON: bytes.NewBufferString("null"),
		YAML: bytes.NewBufferString("null"),
		TOML: bytes.NewBufferString("null"),
	}
	for from := range buf {
		for to := range buf {
			assert.Equal(t, buf[to].String(), convert(buf[from], from, to, false), "from "+from+" to "+to)
		}
	}
}
