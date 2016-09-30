package main

import (
    "fmt"
)

func main() {
    f()
    fmt.Println("The function is main")
}

func f() {
    for i := 0; i < 5; i++ {
        defer fmt.Printf("%d", i)
    }
    defer fmt.Printf("The New Line \n")
    defer fmt.Println("Last output")
}
