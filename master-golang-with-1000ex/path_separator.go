package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string

	dir, file = path.Split("css/main.css")

	fmt.Println("d ir: ", dir)
	fmt.Println("file: ", file)
}
