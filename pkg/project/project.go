package project

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"

	"github.com/karchx/nrs/pkg/todo"
)

const defaultBodySeparator = "---"

// TransformRule defines a title transformation rule.
type TransformRule struct {
	Match   string
	Replace string
}

// Transform applies a title transformation rule
func (transformRule *TransformRule) Transform(title string) string {
	matchRegexp := regexp.MustCompile(transformRule.Match)
	return string(matchRegexp.ReplaceAll(
		[]byte(title), []byte(transformRule.Replace)))
}

// TitleConfig contains project level configuration related to issues titles.
type TitleConfig struct {
	Transforms []*TransformRule
}

// Transform transform the suffix into the title.
func (titleConfig *TitleConfig) Transform(title string) string {
	for _, rule := range titleConfig.Transforms {
		title = rule.Transform(title)
	}
	return title
}

// Project contains project struct.
type Project struct {
	Title         *TitleConfig
	Keywords      []string
	BodySeparator string
	Remote        string
}

func reportedTodoRegexp(keyword string) string {
	return "^(.*)" + regexp.QuoteMeta(keyword) + "(" + regexp.QuoteMeta(string(keyword[len(keyword)-1])) + "*)" + "\\((.*)\\): (.*)$"
}

func (project Project) lineAsReportedTodo(line string) *todo.Todo {
	for _, keyword := range project.Keywords {
		unreportedTodo := regexp.MustCompile(reportedTodoRegexp(keyword))
		groups := unreportedTodo.FindStringSubmatch(line)

		if groups != nil {
			prefix := groups[1]
			urgency := groups[2]
			suffix := groups[3]
			title := project.Title.Transform(suffix)

			return &todo.Todo{
				Prefix:        prefix,
				Suffix:        suffix,
				Keyword:       keyword,
				Urgency:       len(urgency),
				Title:         title,
				BodySeparator: project.BodySeparator,
			}
		}
	}

	return nil
}

// LineAsTodo contructs a Todo from a string.
func (project Project) LineAsTodo(line string) *todo.Todo {
	if todo := project.lineAsReportedTodo(line); todo != nil {
		return todo
	}
	return nil
}

// WalkTodosOfFile visits all of the TODOs in a particular file.
func (project Project) WalkTodosOfFile(path string, visit func(todo.Todo) error) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var todo *todo.Todo

	text, _, err := reader.ReadLine()
	for line := 1; err == nil; line = line + 1 {
		if todo == nil { // LookingForTodo
			todo = project.LineAsTodo(string(text))

			if todo != nil {
				todo.Filename = path
				todo.Line = line
			}
		} else { // CollectingBody
			if possibleTodo := project.LineAsTodo(string(text)); possibleTodo != nil {
				if err := visit(*todo); err != nil {
					return err
				}

				todo = possibleTodo
				todo.Filename = path
				todo.Line = line
			}
		}

		text, _, err = reader.ReadLine()
	}

	if todo != nil {
		if err := visit(*todo); err != nil {
			return err
		}
		todo = nil
	}

	if err != io.EOF {
		return err
	}
	return nil
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

		err = project.WalkTodosOfFile(filepath, visit)
		if err != nil {
			return err
		}
	}

	return nil
}

// NewProject contructs the Project.
func NewProject(filePath string) (*Project, error) {
	project := &Project{
		Title: &TitleConfig{
			Transforms: []*TransformRule{},
		},
		Keywords:      []string{},
		BodySeparator: defaultBodySeparator,
	}

	if len(project.Keywords) == 0 {
		project.Keywords = []string{"TODO"}
	}

	return project, nil
}
