package main

import (
	"encoding/base64"
	"fmt"
	"net/url"
)

func main() {
	params := "app=http://baidu.com&id=1 2"
	fmt.Println(url.QueryEscape(params))
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(params)))
}
