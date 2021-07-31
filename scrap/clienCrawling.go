package scrap

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	util "github.com/jungAcat/Go-Crawling/utils"
	"io"
	"os"
	"strconv"
)

const clienPageNumber = 0

var ClienResultTitle 		[]string
var ClienResultCategory 	[]string
var ClienResultAnchor 		[]string
var	ClienResultDate 		[]string
var ClienResultContent 		[]string
var ClienResultLink			[]string
var ClienResultImg			[]string
var ClienResultHash			[]uint32

func ClienScrapStart() {
	//dupCount := 0
	url := util.ClienBaseUrl + util.ClienResourcePath + strconv.Itoa(clienPageNumber)

	doc := util.GetDocument(url)

	var imageNameBuffer 	bytes.Buffer

	// 1.
	f1, err := os.Create("./ImageNameList.txt")
	util.CheckErr(err)

	searchTR		:= doc.Find(util.ClienTrSelector)

	searchTR.Each(func(i int, sel *goquery.Selection) {
		title 		:= sel.Find(util.ClienTitleSelector)
		category 	:= sel.Find(util.ClienCategorySelector)
		anchor   	:= sel.Find(util.ClienSubSelector)
		img 		:= sel.Find("div.list_img > div > a > img")
		ClienResultTitle 	= append(ClienResultTitle, title.Text())
		ClienResultCategory = append(ClienResultCategory, category.Text())

		// scrap img url
		imgStr, _ 	:= img.Attr("src")
		ClienResultImg	= append(ClienResultImg, imgStr)

		// SubLink Create
		attr, _ := anchor.Attr("href")

		ClienResultAnchor 	= append(ClienResultAnchor, attr)

		subLink				:= util.ClienBaseUrl + attr
		linkHash			:= util.Hash(subLink)
		ClienResultLink 	= append(ClienResultLink, subLink)
		ClienResultHash		= append(ClienResultHash, linkHash)

		docSubLink 	:= util.GetDocument(subLink)

		ClienResultDate 		= append(ClienResultDate, docSubLink.Find(util.ClienDateSelector).Text())
		ClienResultContent  	= append(ClienResultContent, docSubLink.Find(util.ClienContentSelector).Text())

		// Image Download
		if len(imgStr) > 1 {
			// Image Download
			fileFullName 	:= strconv.Itoa(int(linkHash)) + ".png"

			imageNameBuffer.WriteString(fileFullName + "\n")

			//imgURL	:= imgStr
			//util.DownloadFile(imgURL, fileFullName)

		} else {
			fmt.Println("Image Null")
		}

	})

	_, err = io.WriteString(f1, imageNameBuffer.String())
	util.CheckErr(err)

}
