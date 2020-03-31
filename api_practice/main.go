package main

import (
	"encoding/json"
	"first_golang/api_practice/model"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//https://app.swaggerhub.com/apis-docs/Promptech/public-mask-info/20200307-oas3#/
const (
	port    = ":3000"
	baseUrl = "https://8oi9s0nnth.apigw.ntruss.com/corona19-masks/v1"

	storeInfoUrl    = "/stores/json"
	storesByGeoUrl  = "/storesByGeo/json"
	storesByAddrUrl = "/storesByAddr/json"

	myLat = 37.5768224
	myLng = 127.0507571
)

type User struct {
	Name string `json:"name"`
}

func main() {
	e := echo.New()

	e.POST("/", func(c echo.Context) error {

		m := echo.Map{}
		if err := c.Bind(&m); err != nil {
			return err
		}
		return c.JSON(200, m)
	})

	go e.Logger.Fatal(e.Start(port))

}

//내 위치를 기준으로 특정 거리 범위의 약국 마스크 재고확인
func getStoresByGeo(_lat float64, _lng float64, radius int) []model.SortedStore {
	client := &http.Client{}

	req, err := http.NewRequest("GET", baseUrl+storesByGeoUrl, nil)
	manageErr(err)

	query := req.URL.Query()
	query.Add("lat", strconv.FormatFloat(_lat, 'f', 6, 64))
	query.Add("lng", strconv.FormatFloat(_lng, 'f', 6, 64))
	query.Add("m", strconv.Itoa(radius))

	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)
	manageErr(err)

	jsonData, err := ioutil.ReadAll(res.Body)
	manageErr(err)
	defer res.Body.Close()

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonData), &jsonMap)
	manageErr(err)

	// fmt.Println("근처 약국 개수 : ", jsonMap["count"])

	bodyBytes, err := json.Marshal(jsonMap["stores"])
	manageErr(err)

	var storeList []model.SortedStore

	err = json.Unmarshal([]byte(bodyBytes), &storeList)
	manageErr(err)

	// fmt.Println("총 약국 리스트 : ", storeList)
	return storeList
}

func getStoreInfo() []model.StoreInfo {
	res, err := http.Get(baseUrl + storeInfoUrl)
	manageErr(err)
	defer res.Body.Close()

	jsonData, err := ioutil.ReadAll(res.Body)
	manageErr(err)

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonData), &jsonMap)
	manageErr(err)

	bodyBytes, err := json.Marshal(jsonMap["storeInfos"])
	manageErr(err)

	var storeList []model.StoreInfo

	json.Unmarshal([]byte(bodyBytes), &storeList)

	return storeList
}

//거리계산(위도, 경도)
func distance(myLocation model.Location, targetLocation model.Location, unit ...string) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * myLocation.Lat / 180)
	radlat2 := float64(PI * targetLocation.Lat / 180)

	theta := float64(myLocation.Lng - targetLocation.Lng)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}

//에러 핸들링
func manageErr(err error) {
	if err != nil {
		fmt.Println("에러 발생 : ", err)
		panic(err)
	}
}
