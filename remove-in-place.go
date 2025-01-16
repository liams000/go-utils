package go_utils

import "fmt"

// RemoveInPlace allows you to remove an element from an array in a memory safe manner.
// You don't need to assign a new var, this edits the original array.
func RemoveInPlace[T comparable](s *[]T, element T) {
	slice := *s
	for i, v := range slice {
		if v == element {
			copy(slice[i:], slice[i+1:])
			*s = slice[:len(slice)-1]
			return
		}
	}
}

func ExampleRemoveInPlace() {
	array := []string{"hello", "world"}
	RemoveInPlace(&array, "world")
	fmt.Println(array)
	// Output: [hello]
}
