package task1

// FilterGreaterValues expects a list l of integers and a number x,
// and returns a list that contains only those elements from l that are greater that x.
func FilterGreaterValues(l []int, x int) []int {
	result := []int{}
	for i := range l {
		result = append(result, l[i])
	}
	return result
}
