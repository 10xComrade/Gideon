package actions

import (
	"Gideon/config"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

var GlobalNewsResponse *NewsResponse
var CurrentPage *int

type NewsReader struct {
	apiKey string
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
}

type NewsResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

func NewNewsReader(apikey string) *NewsReader {
	return &NewsReader{
		apiKey: apikey,
	}
}

func (N *NewsReader) FetchNews(subject string, sortBy string) (*NewsResponse, error) {
	client := resty.New()

	if config.GlobalConfig.Proxy.Enabled {
		client.SetProxy(config.GlobalConfig.Proxy.URL)
	}

	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&sortBy=%s&apiKey=%s",
		subject,
		sortBy,
		N.apiKey,
	)

	response, err := client.R().Get(url)

	if err != nil {
		log.Fatal("Could not send GET request !")
		return nil, err
	}

	var newsResponse NewsResponse
	err = json.Unmarshal(response.Body(), &newsResponse)
	if err != nil {
		log.Fatal("Could not parse results !")
		return nil, err
	}

	GlobalNewsResponse = &newsResponse
	return GlobalNewsResponse, nil
}

func LimitCurrentPage(currentPage *int) {
	if currentPage == nil {
		currentPage = new(int)
		*currentPage = 0

	} else if *currentPage < 0 {
		*currentPage = 0

	} else if *currentPage >= len(GlobalNewsResponse.Articles) {
		*currentPage = len(GlobalNewsResponse.Articles) - 1
	}

	CurrentPage = currentPage
}
