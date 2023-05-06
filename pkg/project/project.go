package project

import (
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
  //cmd := exec.Command("git", "ls-files", dirpath)
  return nil
}
