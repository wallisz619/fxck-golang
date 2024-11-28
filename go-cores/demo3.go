package main

import "fmt"

func main() {
	/**
	可以把切片看做是对数组的一层简单的封装，
	因为在每个切片的底层数据结构中，一定会包含一个数组。
	数组可以被叫做切片的底层数组，而切片也可以被看作是对数组的某个连续片段的引用。

	*/
	s1 := make([]int, 5)                              // 如果不指明其容量，那么capacity就会和长度一致。
	fmt.Printf("The length of s1 is %d\n", len(s1))   // 5
	fmt.Printf("The capacity of s1 is %d\n", cap(s1)) // 5
	fmt.Printf("The value of s1 is %d\n", s1)         // [0 0 0 0 0]

	s2 := make([]int, 5, 10)                          // length 5, capacity 10
	fmt.Printf("The length of s2 is %d\n", len(s2))   // 5
	fmt.Printf("The capacity of s2 is %d\n", cap(s2)) // 10
	fmt.Printf("The value of s2 is %d\n", s2)         // [0 0 0 0 0]

	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4 is %d\n", len(s4))   // 3
	fmt.Printf("The capacity of s4 is %d\n", cap(s4)) // 7
	fmt.Printf("The value of s4 is %d\n", s4)         // [4 5 6]

	// 缩容
	s5 := make([]int, 3, 10)
	s5[0], s5[1], s5[2] = 1, 2, 3
	fmt.Printf("原切片: len = %d, cap = %d, value = %v\n", len(s5), cap(s5), s5) // 原切片: len = 3, cap = 10, value = [1 2 3]
	// 创建一个新的切片，长度和容量与原切片匹配
	newS5 := make([]int, len(s5))
	copy(newS5, s5)
	fmt.Printf("缩容后: len = %d, cap = %d, value = %v\n", len(newS5), cap(newS5), newS5) // 缩容后: len = 3, cap = 3, value = [1 2 3]

	// 对旧切片的处理: 将其置为nil，自动内存管理进行垃圾回收
	s6 := make([]int, 1000)
	newS6 := make([]int, len(s6))
	copy(newS6, s6)
	s6 = nil // 释放旧切片的内存
}
