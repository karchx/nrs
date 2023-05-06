package todo

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
