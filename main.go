package main

import (
	"github.com/PuerkitoBio/goquery"
	scrap "github.com/jungAcat/Go-Crawling/scrap"
	util "github.com/jungAcat/Go-Crawling/utils"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"fmt"
)
// var fileName = "jobs.csv"

const testNumber = 0

var resultTitle 	*[]string
var resultCategory 	*[]string
var resultAnchor	*[]string
var resultLink		*[]string
var resultDate		*[]string
var resultContent	*[]string
var resultImg		*[]string
var resultHash		*[]uint32

func main() {
	/*e := echo.New()
	e.GET("/", handler)
	e.POST("/scrape", handleFunc)
	e.Logger.Fatal(e.Start(":1323"))*/

	// imgTest()

	decideTestNum(testNumber)
	fmt.Println("!Title!")
	for _, v := range *resultTitle {
		fmt.Println(util.CleanString(v))
	}

	if testNumber != 2 {
		fmt.Println("!Category!")
		for _, v := range *resultCategory {
			fmt.Println(util.CleanString(v))
		}
	}
	fmt.Println("!Anchor!")
	for _, v := range *resultAnchor {
		fmt.Println(util.CleanString(v))
	}
	fmt.Println("!Link!")
	for _, v := range *resultLink {
		fmt.Println(util.CleanString(v))
	}
	fmt.Println("!Date!")
	for _, v := range *resultDate {
		fmt.Println(util.CleanString(v))
	}
	fmt.Println("!Content!")
	for _, v := range *resultContent {
		fmt.Println(util.CleanString(v) + "\n")
	}
	fmt.Println("!Image!")
	for _, v := range *resultImg {
		fmt.Println(v)//util.CleanString(v))
	}
	fmt.Println("!Hash!")
	for _, v := range *resultHash {
		fmt.Println(v)
	}
}


/*
func handler(c echo.Context) error {
	return c.File("./resources/home.html")
}

func handleFunc(c echo.Context) error {
	query := strings.ToLower(scrap.CleanString(c.FormValue("query")))
	fmt.Println(query)
	scrap.Scrapper(query)
	return c.Attachment(fileName, query + ".csv")
}*/

func decideTestNum(i int) {
	switch i {
	case 0:				// Clien
		scrap.ClienScrapStart()
		resultTitle 	= &(scrap.ClienResultTitle)
		resultCategory 	= &(scrap.ClienResultCategory)
		resultAnchor	= &(scrap.ClienResultAnchor)
		resultLink		= &(scrap.ClienResultLink)
		resultDate		= &(scrap.ClienResultDate)
		resultContent	= &(scrap.ClienResultContent)
		resultImg		= &(scrap.ClienResultImg)
		resultHash		= &(scrap.ClienResultHash)

	case 1:				// Cool
		scrap.CoolScrapStart()
		resultTitle 	= &(scrap.CoolResultTitle)
		resultCategory 	= &(scrap.CoolResultCategory)
		resultAnchor	= &(scrap.CoolResultAnchor)
		resultLink		= &(scrap.CoolResultLink)
		resultDate		= &(scrap.CoolResultDate)
		resultContent	= &(scrap.CoolResultContent)

	case 2:				// Eomi
		scrap.EomiScrapStart()
		resultTitle 	= &(scrap.EomiResultTitle)
		//resultCategory 	= &(scrap.EomiResultCategory)
		resultAnchor	= &(scrap.EomiResultAnchor)
		resultLink		= &(scrap.EomiResultLink)
		resultDate		= &(scrap.EomiResultDate)
		resultContent	= &(scrap.EomiResultContent)

	case 3:				// Ppomex
		scrap.PpomexScrapStart()
		resultTitle 	= &(scrap.PpomexResultTitle)
		resultCategory 	= &(scrap.PpomexResultCategory)
		resultAnchor	= &(scrap.PpomexResultAnchor)
		resultLink		= &(scrap.PpomexResultLink)
		resultDate		= &(scrap.PpomexResultDate)
		resultContent	= &(scrap.PpomexResultContent)

	case 4:				// Ppomin
		scrap.PpominScrapStart()
		resultTitle 	= &(scrap.PpominResultTitle)
		resultCategory 	= &(scrap.PpominResultCategory)
		resultAnchor	= &(scrap.PpominResultAnchor)
		resultLink		= &(scrap.PpominResultLink)
		resultDate		= &(scrap.PpominResultDate)
		resultContent	= &(scrap.PpominResultContent)

	case 5:				// Quas
		scrap.QuasScrapStart()
		resultTitle 	= &(scrap.QuasResultTitle)
		resultCategory 	= &(scrap.QuasResultCategory)
		resultAnchor	= &(scrap.QuasResultAnchor)
		resultLink		= &(scrap.QuasResultLink)
		resultDate		= &(scrap.QuasResultDate)
		resultContent	= &(scrap.QuasResultContent)
	}
}


type Sites struct {
	url    string
	images []string
	folder string
}

var crawlers sync.WaitGroup
var downloaders sync.WaitGroup
var verbose bool = false

func (Site *Sites) Crawl() {
	defer crawlers.Done()

	resp, err := goquery.NewDocument(Site.url)
	if err != nil {
		fmt.Printf("ERROR: Failed to crawl \"" + Site.url + "\"\n\n")
		os.Exit(3)
	}
	// use CSS selector found with the browser inspector
	// for each, use index and item
	resp.Find("*").Each(func(index int, item *goquery.Selection) {
		linkTag := item.Find("img")
		link, _ := linkTag.Attr("src")

		if link != "" {
			Site.images = append(Site.images, link)
		}
	})

	fmt.Printf("%s found %d unique images\n", Site.url, len(Site.images))

	pool := len(Site.images) / 3
	if pool > 10 {
		pool = 10
	}

	l := 0
	counter := len(Site.images) / pool

	for i := counter; i < len(Site.images); i += counter {
		downloaders.Add(1)
		go Site.DownloadImg(Site.images[l:i])
		l = i
	}

	downloaders.Wait()
}

func (Site *Sites) DownloadImg(images []string) {

	defer downloaders.Done()

	os.Mkdir(Site.folder, os.FileMode(0777))

	Site.images = SliceUniq(images)

	for _, url := range Site.images {
		if url[:4] != "http" {
			url = "http:" + url
		}
		parts := strings.Split(url, "/")
		name := parts[len(parts)-1]
		file, _ := os.Create(string(Site.folder + "/" + name))
		resp, _ := http.Get(url)
		io.Copy(file, resp.Body)
		file.Close()
		resp.Body.Close()
		if verbose == true {
			fmt.Printf("Saving %s \n", Site.folder+"/"+name)
		}
	}
}

func SliceUniq(s []string) []string {
	for i := 0; i < len(s); i++ {
		for i2 := i + 1; i2 < len(s); i2++ {
			if s[i] == s[i2] {
				// delete
				s = append(s[:i2], s[i2+1:]...)
				i2--
			}
		}
	}
	return s
}

func imgTest() {

	var seedUrls []string

	seedUrls = append(seedUrls, "https://www.clien.net/service/board/jirum?&od=T31&po=0")
	/*
	if len(os.Args) < 2 {
		fmt.Println("ERROR : Less Args\nCommand should be of type : imagescraper [websites]\n\n")
		os.Exit(3)
	}
	if os.Args[1] == "-v" || os.Args[1] == "--verbose" {
		verbose = true
		seedUrls = os.Args[2:]
	}else {
		seedUrls = os.Args[1:]
	}*/
	Site := make([]Sites, len(seedUrls))

	// Crawl process (concurrently)
	for i, name := range seedUrls {
		if name[:4] != "http" {
			name = "http://" + name
		}
		u, err := url.Parse(name)
		if err != nil {
			fmt.Printf("could not fetch page - %s %v", name, err)
		}
		Site[i].folder = u.Host
		Site[i].url = name
		crawlers.Add(1)
		go Site[i].Crawl()
	}

	crawlers.Wait()

	fmt.Printf("\n\nScraped succesfully\n\n")

}