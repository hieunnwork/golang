package main

import "fmt"

func main() {
	//khai bao array
	var array1 [9]int
	fmt.Println(array1)
	//gan gia tri cho array
	array1[0] = 123
	fmt.Println(array1)
	//khai bao 1 array co khoi tao gia tri
	array2 := [9]int{1, 2, 3, 4}
	fmt.Println(array2)
	//size mang
	fmt.Println(len(array2))
	// khai bao mang k can set size
	array3 := [...]string{"H", "i", "e", "u"}
	fmt.Println(array3)
	fmt.Println(len(array3))
	//array la value type khong phai la reference type

	fmt.Println("============")
	array4 := [...]string{"H", "i", "e", "u"}
	copyArray4 := array4 // tao ra 1 vung nho copy array4 gan cho copyArray4 -> viec thay doi gia tri cua copyArray4 k lam anh huong den array4
	//a->(A) b=a <-> b->(A) b=x-> a=x ( reference type)

	copyArray4[0] = "Hieu"
	fmt.Println(array4)
	fmt.Println(copyArray4)

	fmt.Println("============")

	//loop array
	// cach 1
	for i := 0; i < len(array4); i++ {
		fmt.Println(array4[i])
	}

	//cach 2
	for index, value := range array4 {
		fmt.Printf("i= %d value= %s", index, value)
	}
	//blank indentifier
	for index, _ := range array4 {
		fmt.Printf("index=%d", index)
	}

	// mang  hai chieu [hang][cot]
	matrix := [4][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(matrix)

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(matrix[i][j])
		}
	}

}
