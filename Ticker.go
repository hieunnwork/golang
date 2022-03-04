package main

import (
	"fmt"
	"time"
)

// dung de lap di lap lai dieu do 1 cach deu dan (vi du kiem tra du lieu or he thong con dang hoat dong hay k)
func main() {
	duration := time.Duration(1) * time.Second
	tk := time.NewTicker(duration)
	i := 0
	for range tk.C {
		i++
		CheckSomething()
		if i > 9 {
			tk.Stop()
			break
		}
	}
}

func CheckSomething() {
	fmt.Println("Checking...")

}
