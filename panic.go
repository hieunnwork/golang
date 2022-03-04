package main

import "fmt"

//panic dc su dung de khai bao 1 loi khong mong muon khi thuc thi ctrinh ( same try-catch)
//defer : tat ca cac lenh o phia sau tu khoa defer dc thuc hien trc khi ham do tra ve ket qua
//recover: kiem tra xem co panic nao khong , neu co tra ve thong tin cua panic
func main() {
	fmt.Println("Start")
	panicker()
	fmt.Println("End")
}
func panicker() {
	fmt.Println("create panic")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()
	panic("chia mot so la 0")
}
