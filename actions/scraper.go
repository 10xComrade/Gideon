package actions

import (
	"Gideon/config"
	"log"
	"strings"

	colly "github.com/gocolly/colly/v2"
)

func Scrape(char rune, url string, charlimit int) string {
	var texts []string

	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	// c.SetRequestTimeout(60)

	if config.GlobalConfig.Proxy.Enabled {
		c.SetProxy(config.GlobalConfig.Proxy.URL)
	}

	// c.WithTransport(&http.Transport{
	// 	DialContext:           (&net.Dialer{Timeout: 60 * time.Second, KeepAlive: 60 * time.Second}).DialContext,
	// 	MaxIdleConns:          60,
	// 	IdleConnTimeout:       60 * time.Second,
	// 	ExpectContinueTimeout: 60 * time.Second,
	// })

	c.OnHTML("p", func(e *colly.HTMLElement) {
		text := strings.TrimSpace(e.Text)
		if text != "" && len(strings.Join(texts, " ")+text) <= charlimit {
			texts = append(texts, text)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit(url)

	joinedText := strings.Join(texts, " ")

	return joinedText
}
