# go-utils

A personal collection of utility functions and example usages.

## Filter
Allows you to filter with a comparison to any type that you would like.
```go
func ExampleFilter() {
    array := []string{"hello", "world"}
    result := Filter(array, func(s string) bool {
        return strings.HasPrefix(s, "world")
    })
    fmt.Println(result)
    // Output: [world]
}
```

## Remove In Place
Allows you to remove an item from an array without knowing the index.
```go
func ExampleRemoveInPlace() {
    array := []string{"hello", "world"}
    RemoveInPlace(&array, "world")
    fmt.Println(array)
    // Output: [hello]
}
```