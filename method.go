package main

import "fmt"

type student struct {
	name string
	id   int
}

//define method
//func (t Type) methodName (params) return {//body code}
// (t Type)=> Receiver
//1. Value receiver
//2. Pointer Receiver
func (s1 student) getName() string {
	return s1.name
}

// 1. Value Receiver
func (s1 student) changeName() {
	s1.name = "Tuan"

}

//2. Pointer Receiver
func (s1 *student) changeName2() string {
	s1.name = "Tuan"
	return s1.name
}

//non struct (la type ma define lai tu type cua golang )
type String string

func (s String) append(str string) string {
	return str
}
func main() {
	student := student{"Hieu", 10}
	name := student.getName()
	fmt.Println(name)

	name2 := student.changeName2()
	fmt.Println(name2)

	s1 := String("A")
	newStr := s1.append("A")

	fmt.Println(newStr)
}
