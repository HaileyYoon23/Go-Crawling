package scrap

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	util "github.com/jungAcat/Go-Crawling/utils"
	"io"
	"os"
	"strconv"
	"strings"
)

const clienPageNumber = 0

func ClienScrapStart(db *sql.DB) {
	//dupCount := 0
	url := util.ClienBaseUrl + util.ClienResourcePath + strconv.Itoa(clienPageNumber)

	doc := util.GetDocument(url)

	var ImageNameBuffer bytes.Buffer

	ImageNameFile, err := os.Create("./ImageNameList.txt")
	util.CheckErr(err)

	searchTR		:= doc.Find(util.ClienTrSelector)

	searchTR.Each(func(i int, sel *goquery.Selection) {
		title 		:= util.CleanString(sel.Find(util.ClienTitleSelector).Text())
		anchor   	:= sel.Find(util.ClienSubSelector)
		img 		:= sel.Find("div.list_img > div > a > img")

		// scrap img url
		imgStr, _ 	:= img.Attr("src")

		// SubLink Create
		attr, _ := anchor.Attr("href")

		subLink				:= util.ClienBaseUrl + attr
		linkHash			:= util.Hash(subLink)

		if util.IsDuplicate(db, linkHash) {
			// Duplicate : Skip
		} else {
			docSubLink 	:= util.GetDocument(subLink)

			dates	:= strings.Split(docSubLink.Find(util.ClienDateSelector).Text(),"수정일 : ")
			var isImageExist = 0
			date	:= dates[0]
			//}

			content := util.CleanString(docSubLink.Find(util.ClienContentSelector).Text())

			// Image Download
			if len(imgStr) > 1 {
				isImageExist = 1
					// Image Download
					fileFullName 	:= linkHash + ".png"

					ImageNameBuffer.WriteString(fileFullName + "\n")

					imgURL	:= imgStr
					util.DownloadFile(imgURL, fileFullName)

				} else {
					fmt.Println("Image Null")
			}

			newPost	:= util.Post{title, date, content, isImageExist, linkHash, 0,1, subLink}

			util.InsertPost(db, &newPost)
		}
	})

	_, err = io.WriteString(ImageNameFile, ImageNameBuffer.String())
	util.CheckErr(err)

}
