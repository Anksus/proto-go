package main

import (
	"fmt"
	simplepb "proto-go/src/simple"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := simplepb.SimpleMessage{
		Id:         123,
		IsSimple:   true,
		Name:       "Ankit Susne",
		SampleList: []int32{1, 2, 3},
	}
	fmt.Println(sm)
}
