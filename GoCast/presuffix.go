package main

import (
    "fmt"
    "strings"
)

func main() {
    var str string = "This is an example of a string"

    fmt.Printf("T/F? Does the string \"%s\" have prefix %s? \n", str, "Th")
    fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))

    fmt.Printf("\nT/F? Does the string \"%s\" have suffix %s? \n", str, "ing")
    fmt.Printf("%t\n", strings.HasSuffix(str, "ing"))

    fmt.Printf("\nT/F? Does the string \"%s\" have  %s? \n", str, "of")
    fmt.Printf("%t\n", strings.Contains(str, "of"))

    fmt.Printf("%d\n", strings.Index(str, "ofss"))
}
