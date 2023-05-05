package errors

import (
	"fmt"
	"os"
)

// ParsingError represents an error that ocurred while parsing the config file.
type ParsingError struct {
	Err error
}

// Error represents an error that ocurred while parsing the config file
func (e ParsingError) Error() string {
	return fmt.Sprintf("failed : %v", e.Err)
}

// ExitOnError print error and exit
func ExitOnError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
