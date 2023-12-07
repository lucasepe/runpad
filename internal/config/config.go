package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	RecommendedFileName = "runpad.rc"
)

type Item interface {
	Label() string
	Args() []string
}

func FromDir(dir string) (all []Item, err error) {
	configFile := filepath.Join(dir, RecommendedFileName)

	rc, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	return FromReader(rc)
}

func FromReader(in io.Reader) (all []Item, err error) {
	lineno := 0
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		lineno++
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue

		}
		if line[0] == '#' || line[0] == ';' {
			continue
		}

		idx := strings.IndexByte(line, ':')
		if idx == -1 {
			return all, fmt.Errorf(
				"failed parsing line %d; expected key=value pair; got: '%s'",
				lineno, line)
		}

		all = append(all, &task{
			label: line[0:idx],
			args:  strings.Fields(line[idx+1:]),
		})
	}

	return
}

var _ Item = (*task)(nil)

type task struct {
	label string
	args  []string
}

func (t *task) Label() string {
	return t.label
}

func (t *task) Args() []string {
	return t.args
}
