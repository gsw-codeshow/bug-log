package main

import (
	"fmt"

	flatbuffers "github.com/google/flatbuffers/go"
)

type Student_01 struct {
	Name string
}

var (
	handleMap = map[int32]interface{}{
		1: &Student_01{},
	}
	handleFuncMap = map[int32]func(b *flatbuffers.Builder, i interface{}) flatbuffers.UOffsetT{}
)

func main() {
	fmt.Println(handleMap[1])
	return
}
