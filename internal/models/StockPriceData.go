package models

import "time"

type StockPriceData struct {
	Date  time.Time `json:"date"`
	Open  float64   `json:"open"`
	High  float64   `json:"high"`
	Low   float64   `json:"low"`
	Close float64   `json:"close"`
}
