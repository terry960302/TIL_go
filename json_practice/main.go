package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//영화 api를 가져와서 모델로 파싱하기
type Movie struct {
	Title      string   `json:"Title"`
	Year       string   `json:"Year"`
	Rated      string   `json:"Rated"`
	Released   string   `json:"Released"`
	Runtime    string   `json:"Runtime"`
	Genre      string   `json:"Genre"`
	Director   string   `json:"Director"`
	Writer     string   `json:"Writer"`
	Actors     []string `json:"Actors"`
	Plot       string   `json:"Plot"`
	Language   string   `json:"Language"`
	Country    string   `json:"country"`
	Awards     string   `json:"Awards"`
	Poster     string   `json:"Poster"`
	Ratings    []Rating `json:"Ratings"`
	MetaScore  string   `json:"Metascore"`
	imdbRating string   `json:"imdbRating"`
	imdbVotes  string   `json:"imdbVotes"`
	imdbID     string   `json:"imdbID"`
	Type       string   `json:"Type"`
	DVD        string   `json:"DVD"`
	BoxOffice  string   `json:"BoxOffice"`
	Production string   `json:"Production"`
	Website    string   `json:"Website"`
	Response   string   `json:"Response"`
}

type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

func main() {
	host := "http://www.omdbapi.com"

	apiKey := "cbccbe6a"

	client := &http.Client{}

	req, err := http.NewRequest("GET", host, nil)
	if err != nil {
		panic(err)
	}
	query := req.URL.Query()
	query.Add("apiKey", apiKey)
	query.Add("t", "avengers")

	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)

	result, _ := ioutil.ReadAll(res.Body)

	var movie Movie

	json.Unmarshal(result, &movie)

	fmt.Println(movie.Title)
	fmt.Println(movie.Ratings)
	fmt.Println(movie.Genre)
	fmt.Println(movie.Year)
}
