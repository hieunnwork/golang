package main

import "fmt"

func main() {

	var myMap = make(map[string]int)
	fmt.Println(myMap)
	// neu su dung "make" thi gia tri dau tien se khac nil
	if myMap == nil {
		fmt.Println("= nil")
	} else {
		fmt.Println("!= nil")
	}
	fmt.Println("==========")
	//khoi tao map
	var myMap1 map[string]int
	fmt.Println(myMap1)
	// gia tri dau tien se la nil
	if myMap1 == nil {
		fmt.Println("= nil")
	}
	fmt.Println("==========")
	//khai bao voi gia tri khoi tao
	myMap2 := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}
	fmt.Println(myMap2)

	fmt.Println("==========")
	//them 1 phan tu vao map
	myMap2["key4"] = 4
	myMap2["key5"] = 5
	fmt.Println(myMap2)

	fmt.Println("==========")
	//xoa 1 phan tu trong map= delete(map,key)
	delete(myMap2, "key1")

	fmt.Println("==========")
	//map la 1 reference type
	myMap3 := myMap2
	fmt.Println(myMap3)
	myMap3["key5"] = 999
	fmt.Println(myMap3)

	fmt.Println("==========")
	// truy cap 1 phan tu trong map

	value := myMap2["key5"]
	fmt.Println(value)

	fmt.Println("==========")
	//trong map khogn co toan tu

}
