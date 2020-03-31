package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

func main() {
	fileName := flag.String("f", "uos_study_video.mp4", "filename to stream")
	flag.Parse()

	if fileName == nil || *fileName == "" {
		log.Println("파일을 찾을 수 없습니다.")
		return
	}

	baseName := path.Base(*fileName)

	file, err := os.Open(*fileName)
	manageErr(err, "파일을 열 수 없습니다.")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeContent(w, r, baseName, time.Unix(0, 0), file)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))

}

func manageErr(err error, desc string) {
	if err != nil {
		fmt.Println(desc, " : ", err)
		panic(err)
	}
}
