package demos

import (
	"fmt"

	"github.com/src2crs/demo-exam/task1"
	"github.com/src2crs/srcparser/sourcefile"
	"github.com/src2crs/srcrunner/interpreter"
)

// Example_run_all_files uses SrcParser to open and parse all submissions.
// Then it uses SrcRunner to run tests against the functions in the sibmissions.
// It produces a list of results
func Example_run_all_files() {
	// Create a list of ParserResults and populate it with submission names and file paths.
	submission_names := []string{
		"submission1",
		"submission2",
		"submission3",
		"submission4",
		"submission5",
		"submission6",
	}

	results := []*task1.RunnerResult{}
	for _, name := range submission_names {
		results = append(
			results,
			&task1.RunnerResult{Submission: name, Path: fmt.Sprintf("../task1_%s/task.go", name)})
	}

	// Use srcparser.sourcefile.SourceFile to parse
	// each of the files and fill in the Error field.
	for _, result := range results {
		sourcefile := sourcefile.New(result.Path)
		result.HasParserError = sourcefile.HasParserErrors()
		result.ParserError = sourcefile.ParserErrors()
	}

	// Print those elements from the result list, that are not compiling:
	fmt.Println("Not Compiling:")
	fmt.Println("=============")
	for _, result := range results {
		if result.HasParserError {
			fmt.Printf("%s:\n%s\n\n", result.Submission, result.ParserError)
		}
	}

	markingSource := sourcefile.New("../task1/marking.go")
	for _, result := range results {
		if !result.HasParserError {
			submissionfile := sourcefile.New(result.Path)
			source := markingSource.Source()
			source += "\n\n" + submissionfile.Function("FilterGreaterValues").Source()
			interp := interpreter.New()
			interp.Evaluate(source)
			interp.Evaluate("task1.CheckFilterGreaterValues")
			if interp.Ok() {
				f := interp.LastResult().(func(string))
				defer func() {
					if r := recover(); r != nil {
						fmt.Printf("%s: Panic: %s\n", result.Submission, r)
					}
				}()
				f(result.Submission)
			} else {
				fmt.Println(interp.LastError())
			}
		}
	}

	// Output:
	// Not Compiling:
	// =============
	// submission3:
	// ../task1_submission3/task.go:8:14: expected 1 expression
	//
	// submission1_FilterGreaterValues_0_common: OK
	// submission1_FilterGreaterValues_10_common: OK
	// submission1_FilterGreaterValues_20_common: OK
	// submission1_FilterGreaterValues_30_common: OK
	// submission1_FilterGreaterValues_40_common: OK
	// submission1_FilterGreaterValues_50_common: OK
	// submission1_FilterGreaterValues_0_empty: OK
	// submission1_FilterGreaterValues_10_empty: OK
	// submission1_FilterGreaterValues_20_empty: OK
	// submission1_FilterGreaterValues_30_empty: OK
	// submission1_FilterGreaterValues_40_empty: OK
	// submission1_FilterGreaterValues_50_empty: OK
	//
	// submission2_FilterGreaterValues_0_common: OK
	// submission2_FilterGreaterValues_10_common:
	//   function did not produce the expected result:
	//   input: [15 2 3 27 42 8 7]
	//   expected: [15 27 42]
	//   actual: [15 2 3 27 42 8 7]
	// submission2_FilterGreaterValues_20_common:
	//   function did not produce the expected result:
	//   input: [15 2 3 27 42 8 7]
	//   expected: [27 42]
	//   actual: [15 2 3 27 42 8 7]
	// submission2_FilterGreaterValues_30_common:
	//   function did not produce the expected result:
	//   input: [15 2 3 27 42 8 7]
	//   expected: [42]
	//   actual: [15 2 3 27 42 8 7]
	// submission2_FilterGreaterValues_40_common:
	//   function did not produce the expected result:
	//   input: [15 2 3 27 42 8 7]
	//   expected: [42]
	//   actual: [15 2 3 27 42 8 7]
	// submission2_FilterGreaterValues_50_common:
	//   function did not produce the expected result:
	//   input: [15 2 3 27 42 8 7]
	//   expected: []
	//   actual: [15 2 3 27 42 8 7]
	// submission2_FilterGreaterValues_0_empty: OK
	// submission2_FilterGreaterValues_10_empty: OK
	// submission2_FilterGreaterValues_20_empty: OK
	// submission2_FilterGreaterValues_30_empty: OK
	// submission2_FilterGreaterValues_40_empty: OK
	// submission2_FilterGreaterValues_50_empty: OK
	//
	// submission4_FilterGreaterValues_0_common: OK
	// submission4_FilterGreaterValues_10_common: OK
	// submission4_FilterGreaterValues_20_common: OK
	// submission4_FilterGreaterValues_30_common: OK
	// submission4_FilterGreaterValues_40_common: OK
	// submission4_FilterGreaterValues_50_common: OK
	// submission4_FilterGreaterValues_0_empty: OK
	// submission4_FilterGreaterValues_10_empty: OK
	// submission4_FilterGreaterValues_20_empty: OK
	// submission4_FilterGreaterValues_30_empty: OK
	// submission4_FilterGreaterValues_40_empty: OK
	// submission4_FilterGreaterValues_50_empty: OK
	//
	// submission5_FilterGreaterValues_0_common:
	//   function did not produce the expected result:
	//   input: [15 2 3 27 42 8 7]
	//   expected: [15 2 3 27 42 8 7]
	//   actual: [16 3 4 28 43 9 8]
	// submission5_FilterGreaterValues_10_common:
	//   function did not produce the expected result:
	//   input: [15 2 3 27 42 8 7]
	//   expected: [15 27 42]
	//   actual: [16 28 43]
	// submission5_FilterGreaterValues_20_common:
	//   function did not produce the expected result:
	//   input: [15 2 3 27 42 8 7]
	//   expected: [27 42]
	//   actual: [28 43]
	// submission5_FilterGreaterValues_30_common:
	//   function did not produce the expected result:
	//   input: [15 2 3 27 42 8 7]
	//   expected: [42]
	//   actual: [43]
	// submission5_FilterGreaterValues_40_common:
	//   function did not produce the expected result:
	//   input: [15 2 3 27 42 8 7]
	//   expected: [42]
	//   actual: [43]
	// submission5_FilterGreaterValues_50_common: OK
	// submission5_FilterGreaterValues_0_empty: OK
	// submission5_FilterGreaterValues_10_empty: OK
	// submission5_FilterGreaterValues_20_empty: OK
	// submission5_FilterGreaterValues_30_empty: OK
	// submission5_FilterGreaterValues_40_empty: OK
	// submission5_FilterGreaterValues_50_empty: OK
	//
	// submission6: Panic: reflect: slice index out of range
}
