package sites

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"strconv"
)

const ppominPageNumber = 0

var PpominResultTitle 		[]string
var PpominResultCategory 	[]string
var PpominResultAnchor 		[]string
var	PpominResultDate 		[]string
var PpominResultContent 	[]string
var PpominResultLink		[]string

func PpominScrapStart() {
	//dupCount := 0
	url := ppominBaseUrl + ppominResourcePath + strconv.Itoa(ppominPageNumber)

	doc := getDocument(url)

	searchTitle 	:= doc.Find(ppominTitleSelector)
	searchCategory 	:= doc.Find(ppominCategorySelector)
	searchAnchor	:= doc.Find(ppominSubSelector)

	searchTitle.Each(func(i int, str *goquery.Selection) {
		title, _:=  iconv.ConvertString(str.Text(), "euc-kr", "utf-8")
		PpominResultTitle = append(PpominResultTitle, title)
	})
	searchCategory.Each(func(i int, str *goquery.Selection) {
		category, _:=  iconv.ConvertString(str.Text(), "euc-kr", "utf-8")
		PpominResultCategory = append(PpominResultCategory, category)
	})
	searchAnchor.Each(func(i int, str *goquery.Selection) {
		attr, _ := str.Attr("href")

		PpominResultAnchor 	= append(PpominResultAnchor, attr)

		subLink		:= ppominBaseUrl + "/zboard/" +  attr

		PpominResultLink 	= append(PpominResultLink, subLink)

		docSubLink 		:= getDocument(subLink)

		date, _:=  iconv.ConvertString(docSubLink.Find(ppominDateSelector).Text(), "euc-kr", "utf-8")
		content, _:=  iconv.ConvertString(docSubLink.Find(ppominContentSelector).Text(), "euc-kr", "utf-8")

		PpominResultDate 		= append(PpominResultDate, date)
		PpominResultContent  	= append(PpominResultContent, content)
	})
}
