package sites

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"../utils"
)

const quasPageNumber = 1

var QuasResultTitle 	[]string
var QuasResultCategory 	[]string
var QuasResultAnchor 	[]string
var	QuasResultDate 		[]string
var QuasResultContent 	[]string
var QuasResultLink		[]string

func QuasScrapStart() {
	//dupCount := 0
	url := quasBaseUrl + quasResourcePath + strconv.Itoa(quasPageNumber)

	doc := getDocument(url)

	searchTitle 	:= doc.Find(quasTitleSelector)
	searchCategory 	:= doc.Find(quasCategorySelector)
	searchAnchor	:= doc.Find(quasSubSelector)

	searchTitle.Each(func(i int, str *goquery.Selection) {
		QuasResultTitle = append(QuasResultTitle, str.Text())
	})
	searchCategory.Each(func(i int, str *goquery.Selection) {
		QuasResultCategory = append(QuasResultCategory, str.Text())
	})
	searchAnchor.Each(func(i int, str *goquery.Selection) {
		attr, _ := str.Attr("href")

		QuasResultAnchor 	= append(QuasResultAnchor, attr)

		subLink		:= quasBaseUrl + attr

		QuasResultLink 		= append(QuasResultLink, subLink)

		docSubLink 		:= getDocument(subLink)

		QuasResultDate 		= append(QuasResultDate, docSubLink.Find(quasDateSelector).Text())
		QuasResultContent  	= append(QuasResultContent, docSubLink.Find(quasContentSelector).Text())
	})
}
