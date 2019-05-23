package main

import (
	"fmt"
	"net/url"
)

func main() {
	uri := "app=http://baidu.com&id=1 2"
	enURI := url.QueryEscape(uri)
	deURI, _ := url.QueryUnescape(enURI)
	fmt.Println(enURI, deURI)

	parseURL, _ := url.Parse("http://116.62.132.145:9077/walle/deploy?taskId=513023")
	fmt.Println(parseURL, parseURL.Host, parseURL.IsAbs(), parseURL.Query(), parseURL.RequestURI())

	pparseURL, _ := parseURL.Parse("/walle/test")
	fmt.Println(pparseURL)

	u := url.UserPassword("wangxb", "wangxb5175")
	fmt.Println(u, u.Username())

	values, _ := url.ParseQuery("name=Ava&friend=Jess&friend=Sarah&friend=Zoe")
	values.Set("name", "wxb")
	fmt.Println(values, values.Get("friend"), values["friend"], values.Encode())
}
