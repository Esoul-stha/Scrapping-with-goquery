package main

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var (
	targetUrl = "#url"
)

func main() {
	doc, err := goquery.NewDocument(targetUrl)
	if err != nil {
		fmt.Println(err)
	}
	u := url.URL{}
	u.Scheme = doc.Url.Scheme
	u.Host = doc.Url.Host

	doc.Find("#PATH > a").Each(func(_ int, s *goquery.Selection) {

		url, _ := s.Attr("href")
		re := regexp.MustCompile("[0-9]+")
		storeNum := re.FindAllString(url, -1)

		numbers := make([]int, len(storeNum))

		for i, numberString := range storeNum {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				panic(err)
			}
			numbers[i] = number
		}

		// fmt.Printf("%T\n", numbers)
		name, _ := s.Html()
		fmt.Println(name, numbers)
	})

}
