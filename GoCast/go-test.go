package main

import "fmt"
import _ "strings"


func main() {
    //var name string = "王晓勃";
    //length := len(name)
    //fmt.Printf("%s 's length is %d \n", name, length)
    //fmt.Printf("%t\n", strings.HasPrefix(name, "王"))
    i := 0
    for { //since there are no checks, this is an infinite loop
        if i >= 3 { break }
        //break out of this for loop when this condition is met
        fmt.Println("Value of i is:", i)
        i++;
    }
    fmt.Println("A statement just after for loop.")
}
