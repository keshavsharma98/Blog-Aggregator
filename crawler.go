package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RSSFeed struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Atom    string   `xml:"atom,attr"`
	Channel struct {
		Text  string `xml:",chardata"`
		Title string `xml:"title"`
		Link  struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Description   string `xml:"description"`
		Generator     string `xml:"generator"`
		Language      string `xml:"language"`
		LastBuildDate string `xml:"lastBuildDate"`
		Item          []struct {
			Text        string `xml:",chardata"`
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			PubDate     string `xml:"pubDate"`
			Guid        string `xml:"guid"`
			Description string `xml:"description"`
		} `xml:"item"`
	} `xml:"channel"`
}

func Crawler(feedURL string, n, j int) int {
	resp, err := http.Get(feedURL)
	if err != nil {
		log.Println(err)
	}
	log.Println("Hit the URL : ", feedURL)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var rssFeed RSSFeed
	err = xml.Unmarshal(body, &rssFeed)
	if err != nil {
		log.Println(err)
	}

	log.Println("Printing next ", n, " item's titles:")
	for i := j; i < j+n && i < len(rssFeed.Channel.Item); i++ {
		fmt.Println(i+1, ". ", rssFeed.Channel.Item[i].Title)
	}
	return j
}
