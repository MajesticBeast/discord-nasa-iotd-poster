package main

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

type msgItem struct {
	title string
	desc  string
	url   string
}

func main() {

	feedURL := "https://www.nasa.gov/rss/dyn/lg_image_of_the_day.rss"
	fp := gofeed.NewParser()
	var currentItem string

	for {
		time.Sleep(5 * time.Second)

		// Load the feed
		nasaRSS, err := fp.ParseURL(feedURL)
		if err != nil {
			fmt.Printf("Error opening RSS feed: %s\n", feedURL)
		}

		// Check if there is an update to the feed,
		if currentItem == nasaRSS.Items[0].Title {
			fmt.Println("No updated image.")
			continue
		}

		// Assign newest item to current item (to check against next time for update)
		currentItem = nasaRSS.Items[0].Title

		// Print the new item
		item := msgItem{
			title: nasaRSS.Items[0].Title,
			desc:  nasaRSS.Items[0].Description,
			url:   nasaRSS.Items[0].Enclosures[0].URL,
		}
		fmt.Println(nasaRSS.Items[0].Title)
		fmt.Println(nasaRSS.Items[0].Description)
		fmt.Println(nasaRSS.Items[0].Enclosures[0].URL)

		publishToHook(item.title, item.desc, item.url)
	}
}
