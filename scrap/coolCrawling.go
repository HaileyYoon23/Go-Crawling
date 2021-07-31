package scrap

import (
	"github.com/PuerkitoBio/goquery"
	util "github.com/jungAcat/Go-Crawling/utils"
	"strconv"
	"strings"
)

const coolPageNumber = 0

var CoolResultTitle 	[]string
var CoolResultCategory 	[]string
var CoolResultAnchor 	[]string
var	CoolResultDate 		[]string
var CoolResultContent 	[]string
var CoolResultLink		[]string

func CoolScrapStart() {
	//dupCount := 0
	url := util.CoolBaseUrl + util.CoolResourcePath + strconv.Itoa(coolPageNumber)

	doc := util.GetDocument(url)

	searchTR		:= doc.Find(util.CoolTrSelector)

	searchTR.Each(func(i int, sel *goquery.Selection) {
		anchor   	:= sel.Find(util.CoolSubSelector)

		// SubLink Create
		attr, _ := anchor.Attr("href")

		CoolResultAnchor 	= append(CoolResultAnchor, attr)


		subLink				:= attr
		CoolResultLink 	= append(CoolResultLink, subLink)

		docSubLink 	:= util.GetDocument(subLink)

		headerString	:= docSubLink.Find(util.CoolTitleSelector).Text()
		splitTitle 		:= strings.Split(headerString, "|")
		if len(splitTitle) >= 2 {
			CoolResultTitle 	= append(CoolResultTitle, splitTitle[1])
			CoolResultCategory	= append(CoolResultCategory, splitTitle[0])
		} else {
			CoolResultTitle 	= append(CoolResultTitle, " ")
			CoolResultCategory	= append(CoolResultCategory, " ")
		}

		CoolResultDate 		= append(CoolResultDate, docSubLink.Find(util.CoolDateSelector).Text())
		CoolResultContent  	= append(CoolResultContent, docSubLink.Find(util.CoolContentSelector).Text())

	})
}
