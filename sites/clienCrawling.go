package sites

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

const clienPageNumber = 0

var ClienResultTitle 		[]string
var ClienResultCategory 	[]string
var ClienResultAnchor 		[]string
var	ClienResultDate 		[]string
var ClienResultContent 		[]string
var ClienResultLink			[]string

func ClienScrapStart() {
	//dupCount := 0
	url := clienBaseUrl + clienResourcePath + strconv.Itoa(clienPageNumber)

	doc := getDocument(url)

	searchTitle 	:= doc.Find(clienTitleSelector)
	searchCategory 	:= doc.Find(clienCategorySelector)
	searchAnchor	:= doc.Find(clienSubSelector)

	searchTitle.Each(func(i int, str *goquery.Selection) {
		ClienResultTitle = append(ClienResultTitle, str.Text())
	})
	searchCategory.Each(func(i int, str *goquery.Selection) {
		ClienResultCategory = append(ClienResultCategory, str.Text())
	})
	searchAnchor.Each(func(i int, str *goquery.Selection) {
		attr, _ := str.Attr("href")

		ClienResultAnchor 	= append(ClienResultAnchor, attr)
		subLink				:= clienBaseUrl + attr
		ClienResultLink 	= append(ClienResultLink, subLink)

		docSubLink 	:= getDocument(subLink)

		ClienResultDate 		= append(ClienResultDate, docSubLink.Find(clienDateSelector).Text())
		ClienResultContent  	= append(ClienResultContent, docSubLink.Find(clienContentSelector).Text())
	})
}
