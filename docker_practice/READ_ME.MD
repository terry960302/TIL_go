# 간단한 Go프로젝트 Dockerize하기

## main.go 파일

```
// main.go
package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
        w.Write([]byte("Hello world!")
    })
    http.ListenAndServe(":9999", nil)
}
```

</br>

## go 모듈(go.mod) 초기화

```
go mod init
```

</br>

## Docker 파일 만들기

```
# 리눅스 환경의 go언어
FROM golang:1.12.5-alpine3.9

# container 내부에 위치할 app디렉토리에 위치
WORKDIR /app
COPY . .

# go모듈은 기본적으로 깃헙을 사용하기에 git설치하고 flex라이브러리도 설치
RUN apk update && \
    apk add git && \
    go get github.com/cespare/reflex

# container가 open할 포트 설정
EXPOSE 9999

# container 시작을 위한 명령어들
CMD ["reflex", "-c", "reflex.conf"]
```

- 참고로 flex라이브러리는 node.js의 nodemon같은 역할을 담당하고 있음.

</br>

## reflex.conf 파일 만들기

```
-r '(\.go$|go\.mod)' -s go run main.go
```

- -r : watch recursively
- -s : run as service

</br>

## Docker파일 빌드하고 실행하기

```
// 빌드
docker build --tag golang-docker-tutorial:test .

//실행
docker run -p 9999:9999 -v ${pwd}:/app golang-docker-tutorial:test
```

- 만약 여기서 에러가 난다면 docker 프로그램 setting에 들어가서 file sharing에 c드라이브에 체크하고 프로그램을 재시작해주면 됩니다.
