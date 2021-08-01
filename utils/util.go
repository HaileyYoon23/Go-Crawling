package utils

import (
	"bufio"
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

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

func Hash(s string) string {
	hash := sha1.New()
	hash.Write([]byte(s))
	temp := hash.Sum(nil)
	return hex.EncodeToString(temp)
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

func InsertPost(db *sql.DB, p *Post) {
	result, err := db.Exec("INSERT INTO posts(title, date, content, is_image_exist, hash, is_notified, site_id, link) VALUES (?,?,?,?,?,?,?,?)",p.Title, p.Date, p.Content, p.Is_image_exist, p.Hash, p.Is_notified, p.Site_id, p.Link)

	if err != nil {
		log.Fatal(err)
	}

	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}
}

func IsDuplicate(db *sql.DB, hash string) bool {
	var count int
	var result = false

	row := db.QueryRow("SELECT COUNT(*) FROM posts WHERE hash= ?",hash)
	err := row.Scan(&count)
	CheckErr(err)

	if count > 0 {
		result = true
	}

	return result
}

func UploadToTenth() {
	_, err := exec.Command("/bin/sh", "test2.sh").Output()
	if err != nil {
		log.Fatal(err)
	}
}

func RemoveFile(fileName string) {
	file, err := os.Open(fileName)
	CheckErr(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		err = os.Remove(scanner.Text())
		CheckErr(err)
	}
	file.Close()

	err = os.Remove(fileName)
	CheckErr(err)
}