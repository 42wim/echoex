package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type Collect struct {
	Uid         string    `json:"uid"`
	Timestamp   time.Time `json:"timestamp"`
	Ip          string    `json:"ip"`
	App         string    `json:"app"`
	Fingerprint string    `json:"fingerprint"`
	Auth        string    `json:"auth"`
	Country     string    `json:"country"`
	Id          string    `json:"id"`
}

func collectId(c echo.Context) error {
	id := c.Param("id")
	collect := new(Collect)
	if err := c.Bind(collect); err != nil {
		return err
	}
	if collect.Timestamp.IsZero() {
		collect.Timestamp = time.Now()
	}
	collect.Id = id
	return c.JSON(http.StatusCreated, collect)
}

func main() {
	e := echo.New()
	e.POST("/collect/:id", collectId)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":7755"))
}
