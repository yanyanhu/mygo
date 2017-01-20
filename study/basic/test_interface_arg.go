package main

import "fmt"

func MyPrintf(args ...interface{}) {
    for _, arg := range args {
        switch arg.(type) {
        case int:
            fmt.Println(arg, " is an int value.")
        case string:
            fmt.Println(arg, " is a string value.")
        case int64:
            fmt.Println(arg, " is an int64 value.")
        default:
            fmt.Println(arg, " is an unknown type value.")
        }
    }
}

func main() {
    MyPrintf(1, "111", int64(1), 1.23)
}
