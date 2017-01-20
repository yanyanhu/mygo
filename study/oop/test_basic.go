package main

import (
    "fmt"
)


type Integer int

func (a Integer) Less (b Integer) (res bool) {
    res = a < b
    return res
}

func (a *Integer) Add (b Integer) {
    *a += b
}

type Rect struct {
    x, y float64
    width, length float64
}

func (r *Rect) Area() float64 {
    return r.width * r.length
}

func main() {
    var i Integer = 3
    var j Integer = 2
    if i.Less(j) {
        fmt.Println("i is less than j.")
    } else {
        fmt.Println("i is not less than j.")
    }

    fmt.Println(i)
    i.Add(j)
    fmt.Println(i)

    r := new(Rect)
    r.width = 10
    r.length = 10
    a := r.Area()
    fmt.Println(a)

    r1 := &Rect{x:1, y:2, width:10, length:20}
    fmt.Println(r1.Area())

}
