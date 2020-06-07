package HttpUtil

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	urls      []string
	totalSize int
}

func (p *Page) SetUrls(urls []string) {

	p.urls = urls

}

func (p *Page) getTotalSize() int {

	sendRequest(p)
	return p.totalSize
}

func sendRequest(p *Page) {

	totalSize := request(p.urls)
	p.totalSize = totalSize
}

func request(urls []string) int {

	size := 0
	sizes := make(chan int)

	for _, url := range urls {
		go response(url, sizes)
	}

	for i := 0; i < len(urls); i++ {
		sizeFrom := <-sizes
		size = size + sizeFrom
	}

	return size
}

func response(url string, channel chan int) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- len(body)
}
