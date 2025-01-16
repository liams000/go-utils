package go_utils

import (
	"fmt"
	"strings"
)

// Filter allows you to filter items from the array.
func Filter[T any](items []T, f func(T) bool) []T {
	count := 0
	for _, item := range items {
		if f(item) {
			items[count] = item
			count++
		}
	}
	return items[:count]
}

func ExampleFilter() {
	array := []string{"hello", "world"}
	result := Filter(array, func(s string) bool {
		return strings.HasPrefix(s, "world")
	})
	fmt.Println(result)
	// Output: [world]
}
