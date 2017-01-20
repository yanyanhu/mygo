package main

import (
   "fmt"
)


func modify(array [2][10]int) {
    array[0][0] = 10
    fmt.Println("In modify(), array values:", array)
}


func main() {
    array := [2][10]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}

    modify(array)

    fmt.Println(array)

    array1 := [10]int{1, 2, 3, 4, 5, 6, 7}
    fmt.Println(array1)

    myslice := array1[:5]
    fmt.Println(myslice)
    fmt.Println(len(myslice))

    myslice2 := make([]int, 5, 10)
    fmt.Println(myslice2)
    fmt.Println(len(myslice2))

    myslice2 = append(myslice2, 1)
    fmt.Println(myslice2)
    fmt.Println(len(myslice2))
    fmt.Println(cap(myslice2))

    myslice2 = append(myslice2, 1, 2, 3, 4, 5)
    fmt.Println(myslice2)
    fmt.Println(len(myslice2))
    fmt.Println(cap(myslice2))

    fmt.Println(myslice)
    fmt.Println(len(myslice))
    fmt.Println(cap(myslice))
    copy(myslice, myslice2)
    fmt.Println(myslice)
    fmt.Println(len(myslice))
    fmt.Println(cap(myslice))

}
