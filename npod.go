package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

var item msgItem

func getNPOD(cfg *currentItems) msgItem {

	feedURL := "https://www.nasa.gov/rss/dyn/lg_image_of_the_day.rss"
	fp := gofeed.NewParser()

	nasaRSS, err := fp.ParseURL(feedURL)
	if err != nil {
		fmt.Printf("Error opening RSS feed: %s\n", feedURL)
	}

	// Check if there is an update to the feed,
	if cfg.npodCI == nasaRSS.Items[0].Title {
		fmt.Println("No updated image.")
		return item
	}

	cfg.npodCI = nasaRSS.Items[0].Title

	item := msgItem{
		title: nasaRSS.Items[0].Title,
		desc:  nasaRSS.Items[0].Description,
		url:   nasaRSS.Items[0].Enclosures[0].URL,
	}

	return item
}
