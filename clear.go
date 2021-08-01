package main

import (
	"fmt"
	util "github.com/jungAcat/Go-Crawling/utils"
)

func main() {
	util.RemoveFile("./ImageNameList.txt")
	fmt.Println("End!")
}