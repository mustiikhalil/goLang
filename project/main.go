package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
)

var newsMap = make(map[string]NewsMap)

type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

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

func NewsAgg() {
	var s Sitemapindex
	var n News

	// hash map to save all the data
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
	fmt.Println("done")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing news AGG", News: newsMap}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func main() {
	go NewsAgg() // concurrent go routine
	// http.HandleFunc("/", indexHandler)
	// http.ListenAndServe(":8000", nil)
}
