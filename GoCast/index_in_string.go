package main

import (
    "fmt"
    "strings"
)

func main() {
    var str string = "Hi, I'm Wangxb, Hi."
    fmt.Printf("The position of 'Wangxb' is: \n")
    fmt.Printf("%d\n", strings.Index(str, "Wangxb"))

    fmt.Printf("The position of the first instance of 'Hi' is: \n")
    fmt.Printf("%d\n", strings.Index(str, "Hi"))

    fmt.Printf("The position of the last instance of 'Hi' is: \n")
    fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

    fmt.Printf("The position of the last instance of 'Hi' is: \n")
    fmt.Printf("%d\n", strings.LastIndex(str, "Test"))
}
