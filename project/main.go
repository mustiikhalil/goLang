package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

var wg sync.WaitGroup

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
	fetchQueue := make(chan News, 100)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	resp.Body.Close()
	for _, Location := range s.Locations {
		wg.Add(1)
		go newsRoutine(fetchQueue, Location)
	}
	wg.Wait()
	close(fetchQueue)

	for n := range fetchQueue {
		for indx, _ := range n.Titles {
			newsMap[n.Titles[indx]] = NewsMap{n.Keywords[indx], n.Locations[indx]}
		}
	}

	fmt.Println("done")
}

func newsRoutine(c chan<- News, Location string) {
	var n News
	// Hack to solve a problem with https://
	s := strings.Split(Location, "//")

	resp, _ := http.Get("https://" + s[1])
	bytes, _ := ioutil.ReadAll(resp.Body)

	xml.Unmarshal(bytes, &n)
	resp.Body.Close()
	c <- n
	wg.Done()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing news AGG", News: newsMap}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func main() {
	runtime.GOMAXPROCS(4)
	go NewsAgg() // concurrent go routine
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
