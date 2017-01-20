package main

import (
    "fmt"
)

func main() {
    str1 := "Hello, world"
    print(str1)
    print("\n")

    str2 := fmt.Sprintf("String \"str1\" is \"%s\"\n", str1)
    print(str2)

    strlen := len(str1)
    println(strlen)

    for i := 0; i < strlen; i++ {
	println(i)
	println(str1[i])
    }
}
