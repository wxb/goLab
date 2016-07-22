// Show the specified directory structure
package main

import (
  "flag"
  "fmt"
  "os"
  "path"
  "strings"
)

const (
  INDENT = " "
)

var (
  rootPath string
)

func init() {
  flag.StringVar($rootPath, "p", "", "The path of target directory.")
}


func main() {
  flag.Parse()
  if len(rootPath) == 0 {

  }
}
