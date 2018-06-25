package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
	"strings"
)

var BASE_URL = "https://www.fakepersongenerator.com"
var GENERATE_PATH = "/user-face-generator?new=refresh"

func main() {
	fmt.Println("=== Collect User Icons ======")
	doc := getHtmlDoc(BASE_URL + GENERATE_PATH)
	urls := parseDoc2ImageURLs(doc)

	fmt.Printf("===> Download Images(%d counts)\n", len(urls))
	for _, url := range urls {
		download(url)
	}
	fmt.Println("=== Exit Process ======")
}

func download(url string) {
	paths := strings.Split(url, "/")
	fname := paths[len(paths)-1]

	response, err := http.Get(url)
	checkErr(err)
	defer response.Body.Close()

	file, err := os.Create("./temp/" + fname)
	checkErr(err)
	defer file.Close()
	io.Copy(file, response.Body)
}

func parseDoc2ImageURLs(doc *goquery.Document) []string {
	var urls []string
	doc.Find(".face img").Each(func(_ int, s *goquery.Selection) {
		val, _ := s.Attr("src")
		urls = append(urls, BASE_URL+val)
	})

	return urls
}

func getHtmlDoc(url string) *goquery.Document {
	doc, err := goquery.NewDocument(url)
	checkErr(err)
	return doc
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
