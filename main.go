package main

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

func main() {

	feedURL := "https://www.nasa.gov/rss/dyn/lg_image_of_the_day.rss"
	fp := gofeed.NewParser()
	var currentItem string

	for {
		time.Sleep(5 * time.Minute)

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
		fmt.Println(nasaRSS.Items[0].Title)
		fmt.Println(nasaRSS.Items[0].Enclosures[0].URL)

		publishToHook(nasaRSS.Items[0].Title, nasaRSS.Items[0].Enclosures[0].URL)
	}

}
