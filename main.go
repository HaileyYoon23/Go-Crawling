package main

import (
	scrapper "./utils"
	"fmt"
	"github.com/labstack/echo"
	"strings"
)
var fileName = "jobs.csv"

func main() {
	/*e := echo.New()
	e.GET("/", handler)
	e.POST("/scrape", handleFunc)
	e.Logger.Fatal(e.Start(":1323"))*/
	scrapper.ScrapStart()
	for _, v := range scrapper.Result {
		fmt.Println(scrapper.CleanString(v))
	}
}

func handler(c echo.Context) error {
	return c.File("./resources/home.html")
}

func handleFunc(c echo.Context) error {
	query := strings.ToLower(scrapper.CleanString(c.FormValue("query")))
	fmt.Println(query)
	scrapper.Scrapper(query)

	return c.Attachment(fileName, query + ".csv")
}