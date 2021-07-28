package main

import (
	scrapper "./utils"
	"fmt"
)
var fileName = "jobs.csv"

func main() {
	/*e := echo.New()
	e.GET("/", handler)
	e.POST("/scrape", handleFunc)
	e.Logger.Fatal(e.Start(":1323"))*/
	scrapper.PpomexScrapStart()
	fmt.Println("!Title!")
	for _, v := range scrapper.PpomexResultTitle {
		fmt.Println(scrapper.CleanString(v))
	}
	fmt.Println("!Category!")
	for _, v := range scrapper.PpomexResultCategory {
		fmt.Println(scrapper.CleanString(v))
	}
	fmt.Println("!Anchor!")
	for _, v := range scrapper.PpomexResultAnchor {
		fmt.Println(scrapper.CleanString(v))
	}
	fmt.Println("!Date!")
	for _, v := range scrapper.PpomexResultDate {
		fmt.Println(scrapper.CleanString(v))
	}
	fmt.Println("!Content!")
	for _, v := range scrapper.PpomexResultContent {
		fmt.Println(scrapper.CleanString(v))
	}
}


/*
func handler(c echo.Context) error {
	return c.File("./resources/home.html")
}

func handleFunc(c echo.Context) error {
	query := strings.ToLower(scrapper.CleanString(c.FormValue("query")))
	fmt.Println(query)
	scrapper.Scrapper(query)
	return c.Attachment(fileName, query + ".csv")
}*/