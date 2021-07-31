package scrap

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	util "github.com/jungAcat/Go-Crawling/utils"
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
	url := util.PpomexBaseUrl + util.PpomexResourcePath + strconv.Itoa(ppomexPageNumber)

	doc := util.GetDocument(url)

	searchTR		:= doc.Find(util.PpomexTrSelector)

	searchTR.Each(func(i int, sel *goquery.Selection) {
		tempTitle 		:= sel.Find(util.PpomexTitleSelector)
		tempCategory 	:= sel.Find(util.PpomexCategorySelector)
		anchor   		:= sel.Find(util.PpomexSubSelector)

		titleText, _ 	:= iconv.ConvertString(tempTitle.Text(), "euc-kr", "utf-8")
		categoryText, _ := iconv.ConvertString(tempCategory.Text(), "euc-kr", "utf-8")

		PpomexResultTitle = append(PpomexResultTitle, titleText)
		PpomexResultCategory = append(PpomexResultCategory, categoryText)

		// SubLink Create
		attr, _ := anchor.Attr("href")

		PpomexResultAnchor 	= append(PpomexResultAnchor, attr)

		subLink				:= util.PpomexBaseUrl + "/zboard/" + attr
		PpomexResultLink 	= append(PpomexResultLink, subLink)

		docSubLink 	:= util.GetDocument(subLink)

		dateText, _ 	:= iconv.ConvertString(docSubLink.Find(util.PpomexDateSelector).Text(), "euc-kr", "utf-8")
		contentText, _ := iconv.ConvertString(docSubLink.Find(util.PpomexContentSelector).Text(), "euc-kr", "utf-8")

		PpomexResultDate 		= append(PpomexResultDate, dateText)
		PpomexResultContent  	= append(PpomexResultContent, contentText)

	})
}

