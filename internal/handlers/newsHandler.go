package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"shub.codes/ngrokgo-stock/internal/models"
	"shub.codes/ngrokgo-stock/internal/services"
)

func NewsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ticker := strings.ToUpper(r.FormValue("ticker"))
		articles, err := services.GetNewsArticles(ticker)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responseData := struct {
			Ticker   string                `json:"ticker"`
			Articles []*models.NewsArticle `json:"articles"`
		}{
			Ticker:   ticker,
			Articles: articles,
		}

		jsonData, err := json.Marshal(responseData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
