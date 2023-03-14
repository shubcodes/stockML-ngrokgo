package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"shub.codes/ngrokgo-stock/internal/models"
)

const (
	newsAPIURL = "https://api.nytimes.com/svc/search/v2/articlesearch.json"
)

func GetNewsArticles(ticker string) ([]*models.NewsArticle, error) {
	newsAPIKey := os.Getenv("newsKey")
	query := fmt.Sprintf("%s stock", strings.ToUpper(ticker))
	url := fmt.Sprintf("%s?q=%s&api-key=%s", newsAPIURL, query, newsAPIKey)
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//fmt.Println("INITIAL RECEIVE:  ", string(body)) // print response body

	var result struct {
		Response struct {
			Docs []struct {
				Headline struct {
					Main string `json:"main"`
				} `json:"headline"`
				Abstract      string `json:"abstract"`
				Lede          string `json:"lede"`
				WebURL        string `json:"web_url"`
				PublishedDate string `json:"pub_date"`
				Multimedia    []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"multimedia"`
			} `json:"docs"`
		} `json:"response"`
	}

	//fmt.Println("STRUCT:  ", result)

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	fmt.Println("UNMARSHALL:  ", err)

	var articles []*models.NewsArticle
	for _, doc := range result.Response.Docs {
		article := &models.NewsArticle{
			Headline:       doc.Headline.Main,
			Abstract:       doc.Abstract,
			LeadParagraph:  doc.Lede,
			WebURL:         doc.WebURL,
			PublishedDate:  doc.PublishedDate,
			ThumbnailWidth: 0,
			//ThumbnailHeight: 0,
		}
		if len(doc.Multimedia) > 0 {
			article.ThumbnailURL = doc.Multimedia[0].URL
			article.ThumbnailWidth = doc.Multimedia[0].Width
			//article.ThumbnailHeight = doc.Multimedia[0].Height
		}
		articles = append(articles, article)
	}

	fmt.Println("SENT TO HANDLER:  ", articles)

	return articles, nil
}
