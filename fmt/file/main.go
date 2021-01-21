package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	files, _ := filepath.Glob("*")
	fmt.Println(files) // contains a list of all files in the current directory
}
