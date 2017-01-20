package main

import (
    "runtime"
)

const CPU = runtime.NumCPU()

func (v Vector) DoSome(i, n int, u Vector, c chan int) {
    for ; i < n; i++ {
	v[i] += u.Op(v[i])
    }
    c <- 1
}

func (v Vector) DoAll(u Vector) {
    c := make(chan int, NCPU)

    for i := 0; i < CPU; i++ {
	go v.DoSome(i*len(v)/CPU, (i+1)*len(v)/CPU, u, c)
    }

    for i := 0; i < CPU; i++ {
	<-c
    }
}

func main() {
}
