package main

import "fmt"

func main(){
    test(1, 2, 3)
}

func test(a int, b ...int) {
    fmt.Println(b)
}
