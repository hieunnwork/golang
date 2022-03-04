package main

import (
	"fmt"
	"sync"
)

func g1() {
	fmt.Println("g1")
	wg.Done()
}
func g2() {
	fmt.Println("g2")
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	// go func()
	/*
		go g1()
		ng := runtime.NumGoroutine()
		fmt.Println(ng)
		time.Sleep(time.Second)
	*/
	//synchronized goroutines
	wg.Add(2)
	go g1()
	go g2()
	wg.Wait()
}
