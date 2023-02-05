package task1

// FilterGreaterValues expects a list l of integers and a number x,
// and returns a list that contains only those elements from l that are greater that x.
func FilterGreaterValues(l []int, x int) []int {
	result := []int{}
	// SOLUTION-START
	for _, el := range l {
		if el > x {
			result = append(result, el)
		}
	}
	// SOLUTION-END
	return result
}
