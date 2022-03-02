package main

import "fmt"

type Student struct {
	id   int
	name string
}

//struct long struct
type class struct {
	id   int
	name string
}

type school struct {
	class int
	email string
	info  class
}

func main() {

	st1 := Student{
		id:   123,
		name: "Hieu",
	}
	fmt.Println(st1)
	fmt.Println(st1.id)
	fmt.Println(st1.name)

	//cach 2
	st2 := Student{999, "Hieu"}
	fmt.Println(st2)

	//cach 3
	var st3 Student = struct {
		id   int
		name string
	}{id: 99, name: "Hieu"}
	fmt.Println(st3)

	//anonymus struct (vodanh)
	var anonymus = struct {
		email string
		age   int
	}{
		"hieunn@gmail",
		25,
	}
	fmt.Println(anonymus)

	//anonymus fields
	type noName struct {
		string
		int
	}
	n := noName{"Hieu", 99}
	fmt.Println(n)

	//con tro tro den struct
	pointer := &Student{
		999, "Hieu",
	}
	fmt.Println(pointer)
	fmt.Println((*pointer).id)
	fmt.Println(pointer.name)

	//struct long struct
	//cach 1
	student1 := school{
		class: 99,
		email: "hieunn.com",
		info: class{
			id:   10,
			name: "hieu",
		},
	}
	fmt.Println(student1)

	//cach2
	student2 := school{100, "hieunn.com", class{10, "hieu"}}
	fmt.Println(student2)

	//compare 2 struct

}
