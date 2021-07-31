package scrap

import (
	"github.com/PuerkitoBio/goquery"
	util "github.com/jungAcat/Go-Crawling/utils"
	"strconv"
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
	url := util.QuasBaseUrl + util.QuasResourcePath + strconv.Itoa(quasPageNumber)

	doc := util.GetDocument(url)

	searchTR		:= doc.Find(util.QuasTrSelector)

	searchTR.Each(func(i int, sel *goquery.Selection) {
		title 		:= sel.Find(util.QuasTitleSelector)
		category 	:= sel.Find(util.QuasCategorySelector)
		anchor   	:= sel.Find(util.QuasSubSelector)

		QuasResultTitle 	= append(QuasResultTitle, title.Text())
		QuasResultCategory 	= append(QuasResultCategory, category.Text())

		// SubLink Create
		attr, _ := anchor.Attr("href")

		QuasResultAnchor 	= append(QuasResultAnchor, attr)

		subLink				:= util.QuasBaseUrl + attr
		QuasResultLink 		= append(QuasResultLink, subLink)

		docSubLink 		:= util.GetDocument(subLink)

		QuasResultDate 		= append(QuasResultDate, docSubLink.Find(util.QuasDateSelector).Text())
		QuasResultContent  	= append(QuasResultContent, docSubLink.Find(util.QuasContentSelector).Text())

	})
}
