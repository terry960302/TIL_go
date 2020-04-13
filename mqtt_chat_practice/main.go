package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type TestModel struct {
	Lat   float64 `json:"lat"`
	Lng   float64 `json:"lng"`
	Range int64   `json:"range"`
}

func main() {
	e := echo.New()
	e.POST("/", getModel)

	go e.Logger.Fatal(e.Start(":3000"))
}

func getModel(c echo.Context) error {
	req := c.Request()

	decoder := json.NewDecoder(req.Body)

	var testModel TestModel = TestModel{}
	err := decoder.Decode(&testModel)
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, testModel)

}
