package ch8_4_test

import (
	"io"
	"net/http"
	"os"
	"testing"
)

func TestMyCurl(t *testing.T) {

	tests := map[string]string{
		"http://baidu.com": "/tmp/baidu.com.txt",
	}

	for url, file := range tests {
		r, err := http.Get(url)
		if err != nil {
			t.Errorf("Request:%s; Result:%s; Want nil", url, err.Error())
		}

		f, err := os.Create(file)
		if err != nil {
			t.Errorf("Request:%s; Result:%s; Want nil", url, err.Error())
		}
		defer f.Close()

		dest := io.MultiWriter(os.Stdout, f)
		_, err = io.Copy(dest, r.Body)
		if err != nil {
			t.Errorf("Request:%s; Result:%s; Want nil", url, err.Error())
		}

		err = r.Body.Close()
		if err != nil {
			t.Errorf("Request:%s; Result:%s; Want nil", url, err.Error())
		}
	}
}
