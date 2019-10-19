package memo5_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/wxb/goLab/gopl/ch9/memo5"
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

func TestMemo5(t *testing.T) {
	m := memo5.NewMemo(httpGetBody)
	wg := sync.WaitGroup{}

	for k, url := range incomingURLs {
		wg.Add(1)
		go func(k int, url string) {
			start := time.Now()
			defer wg.Done()

			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}

			t.Logf("%d, %s, %s, %d bytes\n", k, url, time.Since(start), len(value.([]byte)))
		}(k, url)
	}

	wg.Wait()
}
