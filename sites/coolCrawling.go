package sites

import (
	"github.com/PuerkitoBio/goquery"
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
	url := coolBaseUrl + coolResourcePath + strconv.Itoa(coolPageNumber)

	doc := getDocument(url)

	searchTitle 	:= doc.Find(coolTitleSelector)
	searchAnchor	:= doc.Find(coolSubSelector)

	searchTitle.Each(func(i int, str *goquery.Selection) {
		splitTitle := strings.Split(str.Text(), "|")
		CoolResultTitle 	= append(CoolResultTitle, splitTitle[1])
		CoolResultCategory	= append(CoolResultCategory, splitTitle[0])
	})
	searchAnchor.Each(func(i int, str *goquery.Selection) {
		attr, _ := str.Attr("href")

		CoolResultAnchor = append(CoolResultAnchor, attr)
		subLink			:= attr

		CoolResultLink 	 = append(CoolResultLink, subLink)

		docSubLink 		:= getDocument(subLink)

		headerString	:= docSubLink.Find(coolTitleSelector).Text()

		splitTitle 		:= strings.Split(headerString, "|")
		if len(splitTitle) >= 2 {
			CoolResultTitle 	= append(CoolResultTitle, splitTitle[1])
			CoolResultCategory	= append(CoolResultCategory, splitTitle[0])
		} else {
			CoolResultTitle 	= append(CoolResultTitle, " ")
			CoolResultCategory	= append(CoolResultCategory, " ")
		}


		CoolResultDate 		= append(CoolResultDate, docSubLink.Find(coolDateSelector).Text())
		CoolResultContent  	= append(CoolResultContent, docSubLink.Find(coolContentSelector).Text())
	})
}
