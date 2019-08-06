package memo1_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/wxb/goLab/gopl/ch9/memo1"
)

var httpGetBody = func(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

var incomingURLs = []string{
	"http://baidu.com",
	"http://2345.com",
	"https://github.com",
	"https://coding.net",

	"http://baidu.com",
	"http://2345.com",
	"https://github.com",
	"https://coding.net",
}

func TestMemo1(t *testing.T) {
	m := memo1.New(httpGetBody)
	for _, url := range incomingURLs {
		start := time.Now()

		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}

		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

func TestMemo1A(t *testing.T) {
	m := memo1.New(httpGetBody)
	wg := sync.WaitGroup{}

	for _, url := range incomingURLs {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			defer func(s time.Time) {
				fmt.Printf("%s \n", time.Since(s))
				wg.Done()
			}(start)

			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}

			fmt.Printf("%s %d ", url, len(value.([]byte)))
		}(url)
	}

	wg.Wait()
}

func BenchmarkMemo01(b *testing.B) {

}
