package task1

import "fmt"

func ExampleFilterGreaterValues_common_cases() {
	list := []int{15, 2, 3, 27, 42, 8, 7}

	fmt.Println(FilterGreaterValues(list, 0))
	fmt.Println(FilterGreaterValues(list, 10))
	fmt.Println(FilterGreaterValues(list, 20))
	fmt.Println(FilterGreaterValues(list, 30))
	fmt.Println(FilterGreaterValues(list, 40))
	fmt.Println(FilterGreaterValues(list, 50))

	// Output:
	// [15 2 3 27 42 8 7]
	// [15 27 42]
	// [27 42]
	// [42]
	// [42]
	// []
}

func ExampleFilterGreaterValues_empty_list() {
	list := []int{}

	fmt.Println(FilterGreaterValues(list, 0))
	fmt.Println(FilterGreaterValues(list, 10))
	fmt.Println(FilterGreaterValues(list, 20))
	fmt.Println(FilterGreaterValues(list, 30))
	fmt.Println(FilterGreaterValues(list, 40))
	fmt.Println(FilterGreaterValues(list, 50))

	// Output:
	// []
	// []
	// []
	// []
	// []
	// []
}
