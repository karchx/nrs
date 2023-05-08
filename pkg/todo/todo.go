package todo

import (
	"fmt"
	"strings"
)

// Todo contains information about a TODO in repo.
type Todo struct {
	Prefix        string
	Suffix        string
	Keyword       string
	Urgency       int
	ID            *string
	Filename      string
	Line          int
	Title         string
	Body          []string
	BodySeparator string
}

// LogString formats TODO for compilation logging.
func (todo Todo) LogString() string {
	urgencySuffix := strings.Repeat(string(todo.Keyword[len(todo.Keyword)-1]), todo.Urgency)

	if todo.ID == nil {
		return fmt.Sprintf("%s:%d: %s%s%s: %s",
			todo.Filename, todo.Line,
			todo.Prefix, todo.Keyword, urgencySuffix,
			todo.Suffix)
	}

	return fmt.Sprintf("%s:%d: %s%s%s(%s): %s",
		todo.Filename, todo.Line,
		todo.Prefix, todo.Keyword, urgencySuffix,
		*todo.ID, todo.Suffix)
}
