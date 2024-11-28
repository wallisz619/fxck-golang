package main

import "fmt"

var container = []string{"a", "b", "c"}

func main() {
	container := map[int]string{1: "a", 2: "b", 3: "c"}
	value, ok := interface{}(container).([]string)
	if !ok {
		fmt.Println("The container is not a slice.")
		return
	}
	fmt.Printf("The element is %q.\n", value)
}
