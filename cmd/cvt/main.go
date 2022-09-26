package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/bytedance/sonic"
	"github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	cmd = &cobra.Command{
		RunE: run,
	}

	from, to string
)

var ErrNotImplement = errors.New("not implement")

func init() {
	cmd.PersistentFlags().StringVarP(&from, "from", "f", "", "")
	cmd.PersistentFlags().StringVarP(&to, "to", "t", "", "")
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGTERM,
		os.Interrupt,
	)
	defer cancel()
	if err := cmd.ExecuteContext(ctx); err != nil {
		os.Exit(-1)
	}
	os.Exit(0)
}
func run(c *cobra.Command, args []string) error {
	bs, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	var t interface{}
	switch from {
	case "json":
		if err := sonic.Unmarshal(bs, &t); err != nil {
			return err
		}
	case "yaml":
		if err := yaml.Unmarshal(bs, &t); err != nil {
			return err
		}
	case "toml":
		if err := toml.Unmarshal(bs, &t); err != nil {
			return err
		}
	default:
		return ErrNotImplement
	}
	switch to {
	case "json":
		s, err := sonic.MarshalString(t)
		if err != nil {
			return err
		}
		fmt.Print(s)
	case "yaml":
		bs, err := yaml.Marshal(t)
		if err != nil {
			return err
		}
		fmt.Print(string(bs))
	case "toml":
		bs, err := toml.Marshal(t)
		if err != nil {
			return err
		}
		fmt.Print(string(bs))
	default:
		fmt.Print(t)
	}
	return nil
}
