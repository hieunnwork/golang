package main

import "fmt"

type Animal interface {
	speak()
}

type Dog struct{}

func (d Dog) speak() {
	fmt.Println("gow gow")
}

//multiple interfaces
type Movement interface {
	move()
}

func (d Dog) move() {
	fmt.Println("meow meow")
}

//Embed interface
type NextAnimal interface {
	Movement
	Animal
}

//empty interface (khi khai bao 1 param là interface thì nó co thể sử dụng với bất kì với kiểu type nào )
func goout(i interface{}) {
	fmt.Println(i)
}
func main() {
	animal := Dog{}

	/*
		var m Movement = animal
		m.move()
		var a Animal = animal
		a.speak()
	*/

	var m NextAnimal = animal
	m.move()
	m.speak()

	goout(10)
	goout(10.02)

}
