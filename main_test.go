package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	buf := map[string]string{
		JSON: `{"a":1.0,"b":["c",2],"d":{"e":"f"}}`,
		YAML: "a: 1.0\nb:\n    - c\n    - 2\nd:\n    e: f\n",
		TOML: "a = 1.0\nb = [\"c\", 2]\n\n[d]\n  e = \"f\"\n",
	}
	for from := range buf {
		for to := range buf {
			assert.Equal(t, buf[to], convert(bytes.NewBufferString(buf[from]), from, to, false), "from "+from+" to "+to)
		}
	}
}
func TestConvertNil(t *testing.T) {
	buf := map[string]string{
		JSON: "",
		YAML: "",
		TOML: "",
	}
	for from := range buf {
		for to := range buf {
			assert.Equal(t, buf[to], convert(bytes.NewBufferString(buf[from]), from, to, false), "from "+from+" to "+to)
		}
	}
}
func TestQuota(t *testing.T) {
	assert.Equal(t,
		"{\"a\":1,\"b\":[\"c\",\"2\"],\"d\":{\"e\":\"f\"}}",
		convert(bytes.NewBufferString(
			`"{\"a\":1,\"b\":[\"c\",\"2\"],\"d\":{\"e\":\"f\"}}"`,
		), JSON, JSON, true),
	)
	assert.Equal(t,
		"a: 1.0\nb:\n    - c\n    - 2\nd:\n    e: f",
		convert(bytes.NewBufferString(
			`"a: 1.0\nb:\n    - c\n    - 2\nd:\n    e: f"`,
		), YAML, YAML, true),
	)
	assert.Equal(t,
		"a = 1.0\nb = [\"c\", 2]\n\n[d]\n  e = \"f\"",
		convert(bytes.NewBufferString(
			`"a = 1.0\nb = [\"c\", 2]\n\n[d]\n  e = \"f\""`,
		), TOML, TOML, true),
	)
}
