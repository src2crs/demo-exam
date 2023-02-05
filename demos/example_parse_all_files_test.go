package demos

import (
	"fmt"

	"github.com/src2crs/srcparser/sourcefile"
)

// ParserResult contains the submission name,
// file path, and the error produced by the parser.
type ParserResult struct {
	Submission string
	Path       string
	Error      string
}

// Example_parse_all_files uses SrcParser to open and parse all submissions.
// It produces a list of results
func Example_parse_all_files() {
	// Create a list of ParserResults and populate it with submission names and file paths.
	submission_names := []string{
		"submission1",
		"submission2",
		"submission3",
		"submission4",
		"submission5",
		"submission6",
	}

	parser_results := []*ParserResult{}
	for _, name := range submission_names {
		parser_results = append(
			parser_results,
			&ParserResult{Submission: name, Path: fmt.Sprintf("../task1_%s/task.go", name)})
	}

	// Use srcparser.sourcefile.SourceFile to parse
	// each of the files and fill in the Error field.
	for _, result := range parser_results {
		sourcefile := sourcefile.New(result.Path)
		result.Error = sourcefile.ParserErrors()
	}

	// Print the result list:
	for _, result := range parser_results {
		fmt.Printf("%s:\n%s\n\n", result.Submission, result.Error)
	}

	// Output:
	// submission1:
	// ../task1_submission1/task.go: no parser errors
	//
	// submission2:
	// ../task1_submission2/task.go: no parser errors
	//
	// submission3:
	// ../task1_submission3/task.go:8:14: expected 1 expression
	//
	// submission4:
	// ../task1_submission4/task.go: no parser errors
	//
	// submission5:
	// ../task1_submission5/task.go: no parser errors
	//
	// submission6:
	// ../task1_submission6/task.go: no parser errors
}
