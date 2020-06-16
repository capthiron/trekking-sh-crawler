package main

import (
	"log"

	"github.com/capthiron/trekking-sh-crawler/scraper"
)

const url string = "https://www.wildes-sh.de/"

func main() {
	var trekkingSiteURIs []string
	scraper.Run(url, &trekkingSiteLinks)
	log.Println(trekkingSiteURIs)
}
