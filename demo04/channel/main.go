package main

import "fmt"

func fibonaci(c, quit chan int) {
	x := 0
	for {
		fmt.Println("执行一次select")
		select {
		case c <- x:
			x++
		case <-quit:
			fmt.Printf("fibonaci 数列已经终止···")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	fmt.Println("测试select监控多路channel")
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonaci(c, quit)

}
