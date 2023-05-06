package utils

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/karchx/nrs/pkg/errors"
	"github.com/karchx/nrs/pkg/project"
	"github.com/karchx/nrs/pkg/todo"
)

func listSubCommand(project project.Project, filter func(todo todo.Todo) bool) error {
  todosToList := []*todo.Todo{}

  // err := project.

  return nil
}

func locateDotGit(directory string) (string, error) {
	absDir, err := filepath.Abs(directory)
	rooted := ""
	if err != nil {
		return "", err
	}

	for absDir != rooted {
		dotGit := path.Join(absDir, ".git")

		if stat, err := os.Stat(dotGit); !os.IsNotExist(err) && stat.IsDir() {
			return dotGit, nil
		}
		rooted = absDir
		absDir = filepath.Dir(absDir)
	}

	return "", fmt.Errorf("Couldn't find .git. Maybe you are not inside of a git repo")
}

func locateProject(directory string) (string, error) {
	dotGit, err := locateDotGit(directory)
	if err != nil {
		return "", errors.ParsingError{Err: err}
	}

	return filepath.Dir(dotGit), nil
}

func GetProject(directory string) *project.Project {
	projectPath, err := locateProject(directory)
	errors.ExitOnError(err)

	project := &project.Project{
		BodySeparator: projectPath,
	}

	return project
}
