package memo4_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/wxb/goLab/gopl/ch9/memo4"
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

func TestMemo4(t *testing.T) {
	m := memo4.New(httpGetBody)
	for _, url := range incomingURLs {
		start := time.Now()

		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}

		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

func TestMemo4A(t *testing.T) {
	m := memo4.New(httpGetBody)
	wg := sync.WaitGroup{}

	for _, url := range incomingURLs {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			defer wg.Done()

			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}

			fmt.Printf("%s, %s, %d \n", url, time.Since(start), len(value.([]byte)))
		}(url)
	}

	wg.Wait()

	for _, url := range incomingURLs {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			defer wg.Done()

			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}

			fmt.Printf("%s, %s, %d \n", url, time.Since(start), len(value.([]byte)))
		}(url)
	}

	wg.Wait()
}

func TestChannelClose(t *testing.T) {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go func(i int) {
			t.Logf("goroutine[%d] start!", i)
			<-ch
			t.Logf("goroutine[%d] finish!", i)
			wg.Done()
		}(i)
	}

	t.Logf("main goroutine sleep 3 second!")
	time.Sleep(time.Second * 3)

	close(ch)
	wg.Wait()
	t.Logf("main goroutine finish!")
}
