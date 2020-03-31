package main

//간단한 golang dockerize하기
//참고 : https://blog.puppyloper.com/menus/Golang/articles/Golang%EA%B3%BC%20docker%EB%A5%BC%20%EC%9D%B4%EC%9A%A9%ED%95%9C%20%EA%B0%9C%EB%B0%9C%ED%99%98%EA%B2%BD%20%EB%A7%8C%EB%93%A4%EA%B8%B0%20(feat.%20hot%20reload)
import "net/http"

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World everyone~"))
	})

	http.ListenAndServe(":9999", nil)
}
