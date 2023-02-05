package task1

import (
	"fmt"
)

// RunnerResult contains the submission name,
// file path, and the error produced by the parser.
type RunnerResult struct {
	Submission     string
	Path           string
	HasParserError bool
	ParserError    string
	TestResults    string
}

// Helper for comparing slices.
func SlicesEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, el := range s1 {
		if el != s2[i] {
			return false
		}
	}
	return true
}

func CheckSingleList(testcase string, input_list, expected_result []int, input_x int) string {
	actual_result := FilterGreaterValues(input_list, input_x)
	message := testcase + ":"
	if !SlicesEqual(expected_result, actual_result) {
		error := "\n  function did not produce the expected result:\n"
		error += fmt.Sprintf("  input: %v\n", input_list)
		error += fmt.Sprintf("  expected: %v\n", expected_result)
		error += fmt.Sprintf("  actual: %v\n", actual_result)
		return message + error
	}
	return message + " OK\n"
}

func CheckFilterGreaterValues(submission string) {
	common_list := []int{15, 2, 3, 27, 42, 8, 7}
	testResults := ""

	{
		expected := []int{15, 2, 3, 27, 42, 8, 7}
		testResults += CheckSingleList(
			submission+"_FilterGreaterValues_0_common",
			common_list,
			expected,
			0,
		)
	}
	{
		expected := []int{15, 27, 42}
		testResults += CheckSingleList(
			submission+"_FilterGreaterValues_10_common",
			common_list,
			expected,
			10,
		)
	}
	{
		expected := []int{27, 42}
		testResults += CheckSingleList(
			submission+"_FilterGreaterValues_20_common",
			common_list,
			expected,
			20,
		)
	}
	{
		expected := []int{42}
		testResults += CheckSingleList(
			submission+"_FilterGreaterValues_30_common",
			common_list,
			expected,
			30,
		)
	}
	{
		expected := []int{42}
		testResults += CheckSingleList(
			submission+"_FilterGreaterValues_40_common",
			common_list,
			expected,
			40,
		)
	}
	{
		expected := []int{}
		testResults += CheckSingleList(
			submission+"_FilterGreaterValues_50_common",
			common_list,
			expected,
			50,
		)
	}
	empty_list := []int{}
	{
		expected := []int{}
		testResults += CheckSingleList(
			submission+"_FilterGreaterValues_0_empty",
			empty_list,
			expected,
			0,
		)
	}
	{
		expected := []int{}
		testResults += CheckSingleList(
			submission+"_FilterGreaterValues_10_empty",
			empty_list,
			expected,
			10,
		)
	}
	{
		expected := []int{}
		testResults += CheckSingleList(
			submission+"_FilterGreaterValues_20_empty",
			empty_list,
			expected,
			20,
		)
	}
	{
		expected := []int{}
		testResults += CheckSingleList(
			submission+"_FilterGreaterValues_30_empty",
			empty_list,
			expected,
			30,
		)
	}
	{
		expected := []int{}
		testResults += CheckSingleList(
			submission+"_FilterGreaterValues_40_empty",
			empty_list,
			expected,
			40,
		)
	}
	{
		expected := []int{}
		testResults += CheckSingleList(
			submission+"_FilterGreaterValues_50_empty",
			empty_list,
			expected,
			50,
		)
	}
	fmt.Println(testResults)
}
