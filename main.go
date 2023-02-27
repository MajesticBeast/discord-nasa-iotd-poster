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
	apodCI string
	npodCI string
}

func main() {

	cfg := currentItems{
		apodCI: "",
		npodCI: "",
	}

	for {
		time.Sleep(4 * time.Hour)

		getNPOD(&cfg)
	}
}
