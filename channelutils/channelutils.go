package channelutils

import (
    "fmt"
    "time"
)

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

func TestChannel2() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    // 开启goroutine将0-100的数发送到ch1中
    go func() {
        for i := 0; i < 2; i++ {
            ch1 <- i
        }
        close(ch1)
    }()
    // 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
    go func() {
        for {
            i, ok := <-ch1 // 通道关闭后再取值ok=false
            if !ok {
                break
            }
            ch2 <- i * i
        }
        close(ch2)
    }()
    // 在主goroutine中从ch2中接收值打印
    for i := range ch2 { // 通道关闭后会退出for range循环
        fmt.Println(i)
    }
}

func TestSelect1(ch chan string) {
    time.Sleep(time.Second * 5)
    ch <- "test1"
}

func TestSelect2(ch chan string) {
    time.Sleep(time.Second * 2)
    ch <- "test2"
}
