package main

import (
    "fmt"
    "errors"
)


func Foo(a int) (n int, err error) {
    if a != 0 {
	n = a*a
    } else {
	err = errors.New("Input parameter can not be 0!")
    }

    return n, err
}

func main() {
    a := 0
    res, err := Foo(a)
    if err != nil {
	fmt.Println(err)
    } else {
        fmt.Println(res)
    }
}
