package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"shub.codes/ngrokgo-stock/internal/services"
)

func StockHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ticker := strings.ToUpper(r.FormValue("ticker"))
		data, err := services.GetStockData(ticker)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonData, err := json.Marshal(data)
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
