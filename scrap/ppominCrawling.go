package scrap

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	util "github.com/jungAcat/Go-Crawling/utils"
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
	url := util.PpominBaseUrl + util.PpominResourcePath + strconv.Itoa(ppominPageNumber)

	doc := util.GetDocument(url)

	searchTR		:= doc.Find(util.PpominTrSelector)

	searchTR.Each(func(i int, sel *goquery.Selection) {
		tempTitle 		:= sel.Find(util.PpominTitleSelector)
		tempCategory 	:= sel.Find(util.PpominCategorySelector)
		anchor   	:= sel.Find(util.PpominSubSelector)

		titleText, _	:=  iconv.ConvertString(tempTitle.Text(), "euc-kr", "utf-8")
		categoryText, _	:=  iconv.ConvertString(tempCategory.Text(), "euc-kr", "utf-8")

		PpominResultTitle = append(PpominResultTitle, titleText)
		PpominResultCategory = append(PpominResultCategory, categoryText)

		// SubLink Create
		attr, _ := anchor.Attr("href")

		PpominResultAnchor 	= append(PpominResultAnchor, attr)

		subLink				:= util.PpominBaseUrl + "/zboard/" +  attr
		PpominResultLink 	= append(PpominResultLink, subLink)

		docSubLink 	:= util.GetDocument(subLink)

		dateText, _		:=  iconv.ConvertString(docSubLink.Find(util.PpominDateSelector).Text(), "euc-kr", "utf-8")
		contentText, _	:=  iconv.ConvertString(docSubLink.Find(util.PpominContentSelector).Text(), "euc-kr", "utf-8")

		PpominResultDate 		= append(PpominResultDate, dateText)
		PpominResultContent  	= append(PpominResultContent, contentText)
	})
}
