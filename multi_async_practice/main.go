package main

import (
	"fmt"
	"net/http"
	"sync"
)

// 고루틴과 채널을 합친 간단한 공부하기 좋은 예시
// [참고] https://www.eventslooped.com/posts/golang-make-multiple-api-calls/

var links = []string{
	"https://github.com/fabpot",
	"https://github.com/andrew",
	"https://github.com/taylorotwell",
	"https://github.com/egoist",
	"https://github.com/HugoGiraudel",
}

func main() {
	checkUrls(links)
}

func checkUrls(urls []string) {
	c := make(chan string)

	var wg sync.WaitGroup

	for _, link := range urls {
		wg.Add(1) // wait group에다가 한 가지일 하고 있는 중이라고 기다리라고 알림.(이거 안하면 걍 함수 끝나면 close됨.)
		go checkUrl(link, c, &wg)
	}

	go func() {
		wg.Wait() // waitgroup이 0이 될때가지 고루틴을 막음.
		close(c)  // 채널은 무조건 닫아줘야함.
	}()

	for msg := range c {
		fmt.Println(msg)
	}
}

func checkUrl(url string, c chan string, wg *sync.WaitGroup) {
	defer (*wg).Done() //  waitgroup이 끝날 때까지 이 함수를 Close하지 않음

	_, err := http.Get(url)

	if err != nil {
		c <- "[ERROR] Could not reach url...."
	} else {
		c <- "[SUCCESS] reached URL!!"
	}
}
