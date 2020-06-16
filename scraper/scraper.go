//Package scraper implements the extraction of trekking site URIs
package scraper

import (
	"io"
	"log"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

//Run runs the extraction of trekking site URIs and writes them into the given results slice
func Run(url string, results *[]string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	*results = extractURIs(&resp.Body)
}

func extractURIs(body *io.ReadCloser) (URIs []string) {
	z := html.NewTokenizer(*body)
	var wg sync.WaitGroup
	c := make(chan string)

	for {
		tt := z.Next()

		if tt == html.ErrorToken {
			break
		}

		if tt == html.StartTagToken {
			wg.Add(1)
			go processToken(&wg, c, z.Token())
		}
	}

	go func() {
		for URI := range c {
			URIs = append(URIs, URI)
		}
	}()

	wg.Wait()

	return URIs
}

func processToken(wg *sync.WaitGroup, c chan string, t html.Token) {
	defer wg.Done()
	if t.Data == "a" {
		for _, a := range t.Attr {
			if a.Key == "href" && strings.HasPrefix(a.Val, "uebernachtungsplaetze/") {
				c <- a.Val
			}
		}
	}
}
