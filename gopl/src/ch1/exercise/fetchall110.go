package exercise

import (
	"time"
	"net/http"
	"io/ioutil"
	"fmt"
	"io"
)

/**
* Created by wanjx on 2018/8/25.
*/

func main() {
	var urls = []string{
		"http://godoc.org",
		"http://gopl.io",
	}

	start := time.Now()
	ch := make(chan string)

	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan <- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}

