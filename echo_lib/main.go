// rest_api_test2 project main.go
package main

import (
	_ "fmt"
	"io"
	"os"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"

	"net/http"
)

func main() {
	e := echo.New()

	//미들웨어
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "김태완" && password == "1234" {
			return true, nil
		}
		return false, nil
	}))

	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /users")
			return next(c)
		}
	}

	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users")
	}, track)

	e.GET("/", func(c echo.Context) error {
		//화면에 띄우는 거 (Node에서 send같은 느낌)
		return c.String(http.StatusOK, "Hello World!")
	})

	e.GET("/users/:id", getUser)

	e.GET("/show", show)

	e.POST("/save", save)

	e.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, u)
	})

	e.Logger.Fatal(e.Start(":8000"))
}

//함수들
func getUser(c echo.Context) error {
	id := c.Param("id") //주소에 있는 id를 그대로 받아옴
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team : "+team+", member : "+member)
}

//이미지 업로드
func save(c echo.Context) error {
	name := c.FormValue("name")

	avatar, err := c.FormFile("avatar")

	if err != nil {
		return err
	}

	src, err := avatar.Open()

	if err != nil {
		return err
	}

	defer src.Close()

	dst, err := os.Create(avatar.Filename)

	if err != nil {
		return err
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b> 고마워!! "+name+"</b>")
}

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}
