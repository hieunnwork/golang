package main

import "fmt"

func main() {
	/*
		//unbuffed channel ( khi gui gia tri vao chanel thi ctrinh se bi block tai cau lenh do luon)
		ch := make(chan int)
		go func() {
			ch <- 99
			fmt.Println("sent") //cau lenh nay se bi block
		}()
		//fmt.Println(<-ch)
		fmt.Println("done")

	*/

	//buffed channel
	/*
		ch1 := make(chan int, 2)
		ch1 <- 1
		ch1 <- 2
		fmt.Println(<-ch1)
		fmt.Println(<-ch1)

	*/
	//select
	/*
		queue := make(chan int)
		done := make(chan bool)

		go func() {
			for i := 0; i < 10; i++ {
				queue <- i
			}
			done <- true
		}()
		for {
			select {
			case v := <-queue:
				fmt.Println(v)
			case <-done:
				fmt.Println("done")
				return
			}

		}

	*/
	//close channel
	queue := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			if i > 5 {
				close(queue)
				break
			} else {
				queue <- i
			}
		}
	}()
	for v := range queue {
		fmt.Println(v)
	}

}
