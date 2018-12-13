package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// for parsing the base index.xml
type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

// Parsing the news
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

// site map
type NewsMap struct {
	Keywords string
	Location string
}

func main() {
	var s Sitemapindex
	var n News

	// hash map to save all the data
	newsMap := make(map[string]NewsMap)
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		// Hack to solve a problem with https://
		s := strings.Split(Location, "//")

		resp, _ := http.Get("https://" + s[1])
		bytes, _ := ioutil.ReadAll(resp.Body)

		xml.Unmarshal(bytes, &n)

		for idx, _ := range n.Titles {
			newsMap[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}
	fmt.Println(newsMap)
}
