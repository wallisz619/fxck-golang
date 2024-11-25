package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
	flag.Parse()
	/*
		%v 是最常用的通用占位符，用于打印任何类型的值。
		%d, %x, %f 等用于格式化输出整数、十六进制、浮点数等常见类型。
		%s, %q 常用于处理字符串类型。
		%p 用于打印指针类型的值。
	*/
	fmt.Printf("Hello, %v!\n", name)

}
