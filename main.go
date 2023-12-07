package main

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/lucasepe/runpad/cmd"
)

var Commit = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value
			}
		}
	}

	return ""
}()

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, cmd.BuildKey, Commit)

	if err := cmd.Run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
}
