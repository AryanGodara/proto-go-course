package main

import (
	"fmt"

	pb "github.com/AryanGodara/proto-go-course/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          42,
		IsSimple:    true,
		Name:        "A Name",
		SampleLists: []int32{1, 2, 3, 4, 5, 6},
	}
}

func main() {
	fmt.Println(doSimple())
}
