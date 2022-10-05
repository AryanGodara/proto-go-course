package main

import (
	"fmt"
	"reflect"

	pb "github.com/AryanGodara/proto-go-course/proto"
	"google.golang.org/protobuf/proto"
)

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"myid1": {Id: 42},
			"myid2": {Id: 43},
			"myid3": {Id: 44},
		},
	}
}

func doOneOf(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		fmt.Println("Neither message nor id, type is: ", x)
	}
}

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		// EyeColor: pb.EyeColor_EYE_COLOR_GREEN,
		EyeColor: 1,
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummyy{Id: 42, Name: "My Name"},
		MultipleDummies: []*pb.Dummyy{
			{Id: 43, Name: "My Second Name"},
			{Id: 44, Name: "My Third Name"},
		},
	}
}

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          42,
		IsSimple:    true,
		Name:        "A Name",
		SampleLists: []int32{1, 2, 3, 4, 5, 6},
	}
}

func doFile(p proto.Message) {
	path := "simple.bin"

	writeToFile(path, p)
	sm2 := &pb.Simple{}
	readFromFile(path, sm2)
	fmt.Println("Read the content:", sm2)
}

func doFromJSON(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(jsonString, message)
	return message
}

func main() {
	// fmt.Println(doSimple())
	// fmt.Println(doComplex())
	// fmt.Println(doEnum())
	// doOneOf(&pb.Result_Id{Id: 42})
	// doOneOf(&pb.Result_Message{Message: "My name"})
	// fmt.Println(doMap())
	// doFile(doSimple())
	// fmt.Println(doFromJSON(toJSON(doSimple()), reflect.TypeOf(pb.Simple{})))
	// fmt.Println(doFromJSON(toJSON(doComplex()), reflect.TypeOf(pb.Complex{})))
	fmt.Println(doFromJSON(`{"id": 42, "unknown": "test"}`, reflect.TypeOf(pb.Simple{})))

	// fmt.Println(doSimple())
	// fmt.Println("###")
	// fmt.Println(doComplex())
	// fmt.Println("###")
	// fmt.Println(doEnum())
	// fmt.Println("###")
	// fmt.Println("###")
	// doOneOf(&pb.Result_Id{Id: 42})
	// fmt.Println("###")
	// doOneOf(&pb.Result_Message{Message: "A Message"})
	// fmt.Println("###")
	// doOneOf(&pb.Dummyy{Id: 11, Name: "Dummy Name"})
	// fmt.Println("###")
	// fmt.Println(doMap())
	// fmt.Println("Reached here")
	// fmt.Println("###")
	// fmt.Println("###")
	// fmt.Println("###")
	// doFile(pb.Dummyy{Id: 11, Name: "Dummy Name"})
}
