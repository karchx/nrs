package project

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
