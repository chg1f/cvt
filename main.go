package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/pelletier/go-toml"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

const (
	JSON = "json"
	YAML = "yaml"
	TOML = "toml"
)

var (
	from, to              string
	quote, unquote        bool
	help, verbose, silent bool
)

func init() {
	pflag.CommandLine.SortFlags = false
	pflag.BoolVarP(&unquote, "unquote", "u", false, "unquote")
	pflag.StringVarP(&from, "from", "f", JSON, "from")
	pflag.StringVarP(&to, "to", "t", JSON, "to")
	pflag.BoolVarP(&quote, "quote", "q", false, "quote")
	pflag.BoolVarP(&verbose, "verbose", "v", false, "verbose")
	pflag.BoolVarP(&silent, "silent", "s", false, "silent")
	pflag.BoolVarP(&help, "help", "h", false, "help")
	// TODO:
	zap.ReplaceGlobals(zap.Must(
		zap.NewDevelopment(
			zap.ErrorOutput(os.Stderr),
			zap.IncreaseLevel(zap.FatalLevel),
		),
	))
}
func main() {
	pflag.Parse()
	if help {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		pflag.PrintDefaults()
		return
	}
	if !silent {
		if !verbose {
		}
	}
	defer zap.L().Sync()
	args := pflag.Args()
	if len(args) == 0 {
		print(convert(os.Stdin, from, to, unquote), quote)
		return
	}
	for ix := range args {
		f, err := os.Open(args[ix])
		if err != nil {
			zap.L().Fatal("", zap.Error(err))
		}
		print(convert(f, from, to, unquote), quote)
		if err := f.Close(); err != nil {
			zap.L().Fatal("", zap.Error(err))
		}
	}
}
func convert(reader io.Reader, from, to string, unquote bool) (s string) {
	bs, err := io.ReadAll(reader)
	if err != nil {
		zap.L().Fatal("", zap.Error(err))
	}

	if unquote {
		defer func() {
			s, err = strconv.Unquote(s)
			if err != nil {
				zap.L().Fatal("", zap.Error(err))
			}
		}()
	}
	if from == to {
		return string(bs)
	}

	var t interface{}

	if bs != nil && len(bs) > 0 {
		switch from {
		case JSON:
			if err := json.Unmarshal(bs, &t); err != nil {
				zap.L().Fatal("", zap.Error(err))
			}
		case YAML:
			if err := yaml.Unmarshal(bs, &t); err != nil {
				zap.L().Fatal("", zap.Error(err))
			}
		case TOML:
			if err := toml.Unmarshal(bs, &t); err != nil {
				zap.L().Fatal("", zap.Error(err))
			}
		default:
			zap.L().Fatal("", zap.Error(err))
		}
	}

	switch to {
	case JSON:
		if t == nil {
			return ""
		}
		bs, err := json.Marshal(t)
		if err != nil {
			zap.L().Fatal("", zap.Error(err))
		}
		return string(bs)
	case YAML:
		if t == nil {
			return ""
		}
		bs, err := yaml.Marshal(t)
		if err != nil {
			zap.L().Fatal("", zap.Error(err))
		}
		return string(bs)
	case TOML:
		if t == nil {
			return ""
		}
		bs, err := toml.Marshal(t)
		if err != nil {
			zap.L().Fatal("", zap.Error(err))
		}
		return string(bs)
	default:
		zap.L().Fatal("", zap.Error(err))
		return ""
	}
}
func print(s string, quote bool) {
	if quote {
		s = strconv.Quote(s)
	}
	fmt.Print(s)
}
