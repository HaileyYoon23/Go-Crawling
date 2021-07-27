package utils

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
)

const pageNumber = 1

var Result []string

func ScrapStart() {
	//dupCount := 0
	url := clienBaseUrl + clienResourcePath + strconv.Itoa(pageNumber)

	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, docErr := goquery.NewDocumentFromReader(res.Body)
	checkErr(docErr)

	//searchTR := doc.Find(clienTrSelector)
	searchTitle := doc.Find(clienTitleSelector)

	searchTitle.Each(func(i int, str *goquery.Selection) {
		Result = append(Result, str.Text())
	})

}