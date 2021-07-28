package sites

import (
	"github.com/PuerkitoBio/goquery"
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
	url := eomiBaseUrl + eomiResourcePath + strconv.Itoa(eomiPageNumber)

	doc := getDocument(url)

	searchAnchor	:= doc.Find(eomiSubSelector)

	EomiResultCategory = "의류/잡화"

	searchAnchor.Each(func(i int, str *goquery.Selection) {
		attr, _ := str.Attr("href")

		EomiResultAnchor 	= append(EomiResultAnchor, attr)
		subLink			:= attr
		EomiResultLink 		= append(EomiResultLink, subLink)

		docSubLink 		:= getDocument(subLink)

		EomiResultTitle 	= append(EomiResultTitle, docSubLink.Find(eomiDateSelector).Text())
		EomiResultDate 		= append(EomiResultDate, docSubLink.Find(eomiDateSelector).Text())
		EomiResultContent  	= append(EomiResultContent, docSubLink.Find(eomiContentSelector).Text())
	})
}
