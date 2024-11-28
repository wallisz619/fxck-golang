package main

import "fmt"

func main() {
	s1 := make([]int, 5)                              // length 5, capacity 5
	fmt.Printf("The length of s1 is %d\n", len(s1))   // 5
	fmt.Printf("The capacity of s1 is %d\n", cap(s1)) // 5
	fmt.Printf("The value of s1 is %d\n", s1)         // [0 0 0 0 0]

	s2 := make([]int, 5, 10)                          // length 5, capacity 10
	fmt.Printf("The length of s2 is %d\n", len(s2))   // 5
	fmt.Printf("The capacity of s2 is %d\n", cap(s2)) // 10
	fmt.Printf("The value of s2 is %d\n", s2)         // [0 0 0 0 0]
}
