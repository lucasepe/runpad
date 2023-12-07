package cmd

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/lucasepe/runpad/internal/config"
	"github.com/lucasepe/runpad/internal/gui"
)

const (
	appName = "runpad"
)

type (
	buildKey struct{}
)

var (
	BuildKey = buildKey{}
)

func Run(ctx context.Context) error {
	bld := ctx.Value(BuildKey).(string)

	fs := flag.CommandLine
	fs.Init(appName, flag.ExitOnError)
	fs.Usage = func() {
		fs.SetOutput(os.Stderr)

		fmt.Fprintln(fs.Output(), "           ┓")
		fmt.Fprintln(fs.Output(), "┏┓┓┏┏┓┏┓┏┓┏┫")
		fmt.Fprintln(fs.Output(), "┛ ┗┻┛┗┣┛┗┻┗┻")
		fmt.Fprintf(fs.Output(), "      ┛  (%s)\n", bld)
		fmt.Fprintln(fs.Output())
		fmt.Fprintln(fs.Output(), "Create a menu to run commands.")
		fmt.Fprintln(fs.Output())
		fmt.Fprintln(fs.Output(), "» crafted with passion by Luca Sepe <https://github.com/lucasepe/runpad>")

		fs.SetOutput(io.Discard)
	}

	fs.Bool("help", false, "Print usage and this help message and exit")

	if err := fs.Parse(os.Args[1:]); err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	all, err := config.FromDir(cwd)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("missing configuration file: %s", config.RecommendedFileName)
		}
		return err
	}

	gui.Show(all, appName)

	return nil
}
