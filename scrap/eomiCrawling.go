package scrap

import (
	"github.com/PuerkitoBio/goquery"
	util "github.com/jungAcat/Go-Crawling/utils"
	"strconv"
)

const eomiPageNumber = 1

var EomiResultTitle 	[]string
var EomiResultCategory 	string
var EomiResultAnchor 	[]string
var	EomiResultDate 		[]string
var EomiResultContent 	[]string
var EomiResultLink		[]string

func EomiScrapStart() {
	//dupCount := 0
	url := util.EomiBaseUrl + util.EomiResourcePath + strconv.Itoa(eomiPageNumber)

	doc := util.GetDocument(url)

	EomiResultCategory = "의류/잡화"

	searchTR		:= doc.Find(util.EomiTrSelector)

	searchTR.Each(func(i int, sel *goquery.Selection) {
		// SubLink Create
		anchor 		:= sel.Find(util.EomiSubSelector)
		attr, _		:= anchor.Attr("href")

		EomiResultAnchor 	= append(EomiResultAnchor, attr)

		subLink				:= attr
		EomiResultLink 	= append(EomiResultLink, subLink)

		docSubLink 	:= util.GetDocument(subLink)

		EomiResultTitle = append(EomiResultTitle, docSubLink.Find(util.EomiTitleSelector).Text())
		EomiResultDate 		= append(EomiResultDate, docSubLink.Find(util.EomiDateSelector).Text())
		EomiResultContent  	= append(EomiResultContent, docSubLink.Find(util.EomiContentSelector).Text())

	})
}
