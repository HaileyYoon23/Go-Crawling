package utils

import (
	"encoding/csv"
	"fmt"
	"hash/fnv"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)


//Scrapper function
func Scrapper(query string) {
	var jobs []Indeed
	var baseURL string = "https://kr.indeed.com/jobs?q=" + query
	c1 := make(chan []Indeed)
	TotalPage := getPages(baseURL)
	fmt.Println("TotalPage...", TotalPage)

	for i := 0; i < TotalPage; i++ {
		go getCard(i, baseURL, c1)
	}

	for i := 0; i < TotalPage; i++ {
		extractJobs := <-c1
		//merge slices or arrays
		jobs = append(jobs, extractJobs...)
	}
	writeJobs(jobs)
	fmt.Println("Done")
}

func writeJobs(jobs []Indeed) {
	file, err := os.Create("jobs.csv")
	CheckErr(err)

	w := csv.NewWriter(file)
	//Write data to the file
	defer w.Flush()

	header := []string{"ID", "TITLE", "LOCATION", "SUMMARY"}

	wErr := w.Write(header)
	CheckErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.summary}
		jobErr := w.Write(jobSlice)
		CheckErr(jobErr)
	}
}

func getCard(page int, baseURL string, c1 chan []Indeed) {
	var jobs []Indeed
	c := make(chan Indeed)
	URL := baseURL + "&start=" + strconv.Itoa(page*10)
	fmt.Println(URL)

	res, err := http.Get(URL)
	CheckErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	CheckErr(err)
	searchCards := doc.Find(".jobsearch-SerpJobCard")

	searchCards.Each(func(i int, s *goquery.Selection) {
		go extractJob(s, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	c1 <- jobs
}

func extractJob(s *goquery.Selection, c chan<- Indeed) {
	id, _ := s.Attr("data-jk")
	title := CleanString(s.Find(".title>a").Text())
	location := CleanString(s.Find(".accessible-contrast-color-location").Text())
	summary := CleanString(s.Find(".summary").Text())

	c <- Indeed{
		id:       id,
		title:    title,
		location: location,
		summary:  summary}
}

func getPages(baseURL string) int {
	pages := 0
	res, err := http.Get(baseURL)
	CheckErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	CheckErr(err)

	doc.Find(".pagination-list").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("li").Length()
	})

	return pages
}

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalf("Status code err: %d %s", res.StatusCode, res.Status)
	}
}

//CleanString function
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func GetDocument(url string) (*goquery.Document) {
	response, error := http.Get(url)
	CheckErr(error)
	checkCode(response)

	defer response.Body.Close()

	doc, docErr := goquery.NewDocumentFromReader(response.Body)
	CheckErr(docErr)

	return doc
}

func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func DownloadFile(URL, fileName string) {
	response, error := http.Get(URL)
	CheckErr(error)
	checkCode(response)

	defer response.Body.Close()

	//Create a empty file
	file, error := os.Create(fileName)
	CheckErr(error)
	defer file.Close()

	//Write the contents to the file
	_, error = io.Copy(file, response.Body)
	CheckErr(error)
}