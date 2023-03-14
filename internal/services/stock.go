package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"shub.codes/ngrokgo-stock/internal/models"
)

const (
	alphaVantageURL = "https://www.alphavantage.co/query"
)

type AlphaVantageResponse struct {
	TimeSeriesDaily map[string]struct {
		Open  string `json:"1. open"`
		High  string `json:"2. high"`
		Low   string `json:"3. low"`
		Close string `json:"4. close"`
	} `json:"Time Series (Daily)"`
}

func GetStockData(ticker string) ([]*models.StockPriceData, error) {
	alphaVantageAPIKey := os.Getenv("avKey")
	fmt.Println("API Key:", alphaVantageAPIKey)
	url := fmt.Sprintf("%s?function=TIME_SERIES_DAILY_ADJUSTED&symbol=%s&apikey=%s", alphaVantageURL, ticker, alphaVantageAPIKey)
	log.Printf("Making request to %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(body)) // print response body

	var result AlphaVantageResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	var data []*models.StockPriceData
	for date, prices := range result.TimeSeriesDaily {
		dateLayout := "2006-01-02"
		d, err := time.Parse(dateLayout, date)
		if err != nil {
			return nil, err
		}

		open, _ := strconv.ParseFloat(prices.Open, 64)
		high, _ := strconv.ParseFloat(prices.High, 64)
		low, _ := strconv.ParseFloat(prices.Low, 64)
		close, _ := strconv.ParseFloat(prices.Close, 64)

		data = append(data, &models.StockPriceData{
			Date:  d,
			Open:  open,
			High:  high,
			Low:   low,
			Close: close,
		})
	}

	return data, nil
}
