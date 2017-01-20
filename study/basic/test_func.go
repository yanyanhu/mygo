package main
import (
    "errors"
    "fmt"
    "test_func_pkg"
)


func Add(a int, b int) (ret int, err error) {
    if a < 0 || b < 0 {
        err = errors.New("Should be non-negative numbers.")
        return
    }

    return a + b, nil
}


func Del(a int, b int) (ret int) {
    return a - b
}


func myprint(args ...int) {
    for _, arg := range args {
	println(arg)
    }
}

func myprint2(args ...interface{}) {
    for _, arg := range args {
	fmt.Println(arg)
    }
}


func main() {
    x := true
    print(x)
    print("\n")
    y := !x
    print(y)
    print("\n")

    res, _ := Add(1, 2)
    print(res)
    print("\n")

    res = Del(3, 1)
    print(res)
    print("\n")

    func(m int) {
	print(m)
	print("\n")
    } (123)

    func (a int) {
        print(a)
        print("\n")
    } (1)

    zz := test_func_pkg.Multiply(11, 2)
    println(zz)

    myprint(1, 3, 5, 7, 9)
    //var a []int
    //a = []int{1, 2, 3}
    a := []int{1, 2, 3}
    myprint(a...)

    myprint2(1, "abc", 3.5)
    println("abc")
}
