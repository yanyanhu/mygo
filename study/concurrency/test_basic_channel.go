package main

import (
    "fmt"
    "time"
)

func Count(ch chan int) {
    ch <- 1
    fmt.Println("counting...")
}

//func receive(ch1 chan int, ch2 chan string) {
//    //for i := range(ch1) {
//    //    fmt.Println(i)
//    //}
//    for {
//	    i := <-ch1
//            fmt.Println(i)
//	    if i == 9 {
//                break
//	}
//    }
//    fmt.Println("ok")
//    ch2 <- "ok"
//}

func main() {
//    chs := make([]chan int, 10)
//    for i := 0; i < 10; i++ {
//	chs[i] = make(chan int)
//	go Count(chs[i])
//    }
//
//    for _, ch := range(chs) {
//	<-ch
//    }

//    ch_timeout := make(chan bool)
//    go func() {
//	time.Sleep(5e9)
//	ch_timeout <- true
//    }()
//
//    for j := 0; j < 10; j++ {
//        select {
//	    case <-ch_timeout:
//		fmt.Println("Counting is timeout.")
//		return
//	    default:
//	        fmt.Println(j)
//	        time.Sleep(1e9)
//        }
//    }

    //The example above will abort the send operation if the send blocks for more than 5 seconds.
    strChan := make(chan string)
    select {
        case strChan <- "value":
	    fmt.Println("got value")
        case <-time.After(5 * time.Second):
	    fmt.Println("timeout")
    }

}
