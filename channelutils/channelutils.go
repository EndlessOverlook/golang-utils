package channelutils

import "fmt"

func TestChannel1() {
	c := make(chan int)
	go func() {
		for i := 1; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main执行结束")
}
