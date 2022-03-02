package main

import "fmt"

func main() {
	a := 100
	var pointer *int // hoac pointer :=new(int)
	pointer = &a
	fmt.Println(pointer)
	fmt.Printf("%T", pointer)

	//zero value
	var pointer1 *int
	fmt.Println(pointer1)
	// neu su dung toan tu new thi chuong trinh se tu cap phat 1 dia chi nen k con gia tri nil nua
	pointer2 := new(int)
	fmt.Println(pointer2)

	//con tro tro toi 1 bien
	pointer1 = &a
	*pointer1 = 999
	fmt.Println(*pointer1)
	fmt.Println(a)

	//con tro tro toi mang
	array := [4]int{1, 2, 3, 4}
	var pointer3 *[4]int
	pointer3 = &array
	fmt.Println(pointer3)

	//truyen con tro vao ham
	c := 89
	var pointer4 *int
	pointer4 = &c
	applyPointer(pointer4)
	fmt.Println(c)
}

func applyPointer(p1 *int) {
	*p1 = 777
}
