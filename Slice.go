package main

import "fmt"

func main() {
	//khai bao slice
	var silce []int
	fmt.Println(silce)

	//khai bao va khoi tao slice
	silce1 := []int{1, 2, 3, 4}
	fmt.Println(silce1)

	// tao 1 slice tu 1 array
	array := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice2 := array[1:8] // array[1]-> array[7]
	slice3 := array[:]
	slice4 := array[2:]
	slice5 := array[:8]
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)

	//slice -> reference type
	var array1 = [5]int{1, 2, 3, 4, 5}

	slice6 := array1[:]
	slice6[0] = 999
	fmt.Println(slice6)
	fmt.Println(array1)

	//length va capacity cua slice
	Name := [...]string{"Hieu", "Thao", "Thuy", "Huong"}
	slice7 := Name[1:2]
	fmt.Println(slice7)
	fmt.Println(len(slice7))
	fmt.Println(cap(slice7))
	//len: so luong phan tu cua slice
	//cap: so luong phan tu cua underlying array bat dau tu vi tri strat khi ma slice duoc tao ( từ vị trí "thao" rồi dếm đén cuối dãy nên =3)

	fmt.Println("==============")

	//make ( cung cap gia tri cu the cho len va cap cua 1 slice)
	slice8 := make([]int, 2, 5)
	fmt.Println(slice8)
	fmt.Println(len(slice8))
	fmt.Println(cap(slice8))

	fmt.Println("==============")
	//append ( them gia tri vao slice)
	var slice9 = []int{1, 2, 3}
	slice9 = append(slice9, 89)
	fmt.Println(slice9)

	fmt.Println("==============")

	//copy
	src := []int{1, 2, 3, 4, 5}
	dest := make([]int, 4)
	copy(dest, src)
	fmt.Println(dest)

	fmt.Println("==============")

	//delete item with index =1
	src = append(src[:1], src[2:]...) // slice-slice = append(slice1,slice2...)
	fmt.Println(src)

}
