package utils

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"

	"github.com/karchx/nrs/pkg/errors"
	"github.com/karchx/nrs/pkg/project"
	"github.com/karchx/nrs/pkg/todo"
)

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

func ListSubCommand(project project.Project, filter func(todoP todo.Todo) bool) error {
	todosToList := []*todo.Todo{}

	err := project.WalkTodosOfDir(".", func(todoP todo.Todo) error {
		if filter(todoP) {
			todosToList = append(todosToList, &todoP)
		}
		return nil
	})
	if err != nil {
		return err
	}

	sort.Slice(todosToList, func(i, j int) bool {
		return todosToList[i].Urgency > todosToList[j].Urgency
	})

	for _, todo := range todosToList {
		fmt.Println(todo.LogString())
	}

	return nil
}

func GetProject(directory string) *project.Project {
	projectPath, err := locateProject(directory)
	errors.ExitOnError(err)

	project, err := project.NewProject(projectPath)
	errors.ExitOnError(err)

	return project
}

//func GetRepo(directory string, remote string) (string, issue.IssuesAPI, error) {
//  credentials := getCredentials()
//}
