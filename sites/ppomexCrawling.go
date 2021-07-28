package sites

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"strconv"
)

const ppomexPageNumber = 0

var PpomexResultTitle 		[]string
var PpomexResultCategory 	[]string
var PpomexResultAnchor 		[]string
var	PpomexResultDate 		[]string
var PpomexResultContent	 	[]string
var PpomexResultLink		[]string

func PpomexScrapStart() {
	//dupCount := 0
	url := ppomexBaseUrl + ppomexResourcePath + strconv.Itoa(ppomexPageNumber)

	doc := getDocument(url)

	searchTitle 	:= doc.Find(ppomexTitleSelector)
	searchCategory 	:= doc.Find(ppomexCategorySelector)
	searchAnchor	:= doc.Find(ppomexSubSelector)

	searchTitle.Each(func(i int, str *goquery.Selection) {
		title, _:=  iconv.ConvertString(str.Text(), "euc-kr", "utf-8")
		PpomexResultTitle = append(PpomexResultTitle, title)
	})
	searchCategory.Each(func(i int, str *goquery.Selection) {
		category, _:=  iconv.ConvertString(str.Text(), "euc-kr", "utf-8")
		PpomexResultCategory = append(PpomexResultCategory, category)
	})
	searchAnchor.Each(func(i int, str *goquery.Selection) {
		attr, _ := str.Attr("href")

		PpomexResultAnchor 	= append(PpomexResultAnchor, attr)

		subLink		:= ppomexBaseUrl + "/zboard/" +  attr

		PpomexResultLink 	= append(PpomexResultLink, subLink)

		docSubLink 	:= getDocument(subLink)

		date, _		:=  iconv.ConvertString(docSubLink.Find(ppomexDateSelector).Text(), "euc-kr", "utf-8")
		content, _	:=  iconv.ConvertString(docSubLink.Find(ppomexContentSelector).Text(), "euc-kr", "utf-8")

		PpomexResultDate 		= append(PpomexResultDate, date)
		PpomexResultContent  	= append(PpomexResultContent, content)
	})
}

