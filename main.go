package main

import (
	"time"
)

type msgItem struct {
	title string
	desc  string
	url   string
}

type currentItems struct {
	npodCI string
}

func main() {

	cfg := currentItems{
		npodCI: "",
	}

	item := getNPOD(&cfg)
	publishToHook(item.title, item.desc, item.url)

	for {
		time.Sleep(4 * time.Hour)
		item = getNPOD(&cfg)

		if len(item.title) < 1 {
			continue
		}

		publishToHook(item.title, item.desc, item.url)
	}
}
