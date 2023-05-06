package project

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/karchx/nrs/pkg/todo"
)

// TransformRule defines a title transformation rule.
type TransformRule struct {
	Match   string
	Replace string
}

// TitleConfig contains project level configuration related to issues titles.
type TitleConfig struct {
	Transforms []*TransformRule
}

// Project contains project struct.
type Project struct {
	Title         *TitleConfig
	Keywords      []string
	BodySeparator string
	Remote        string
}

// WalkTodosOfDir visits all of the TODOs in a particular directory.
func (project Project) WalkTodosOfDir(dirpath string, visit func(todo.Todo) error) error {
	cmd := exec.Command("git", "ls-files", dirpath)
	var outBuffer bytes.Buffer
	cmd.Stdout = &outBuffer

	err := cmd.Run()
	if err != nil {
		return err
	}

	for scanner := bufio.NewScanner(&outBuffer); scanner.Scan(); {
		filepath := scanner.Text()
		stat, err := os.Stat(filepath)
		if err != nil {
			return err
		}

		if stat.IsDir() {
			fmt.Printf("[WARN] `%s` is probably a submodule. Skipping it for now...\n", filepath)
			continue
		}
	}

	return nil
}
