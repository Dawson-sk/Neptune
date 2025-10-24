package main

import (
//	"github.com/chromedp/chromedp"
//	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
//	"sync/atomic"
)

func mid(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		t := next
		for _, m := range middleware {
			t = m(t)
		}
		t.ServeHTTP(w, req)
	})
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		next.ServeHTTP(w, req)
	})
}

func recoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Log panic and stack trace
				log.Println("Error: ", err, string(debug.Stack()))
				
				// Send error message to client
				
			}
		}()
		next.ServeHTTP(w, req)
	})
}

var middleware = []func(next http.HandlerFunc)http.HandlerFunc{
	recoveryMiddleware,
	loggingMiddleware,
}

func main() {
	cfg, err := configure()
	if err != nil {
		log.Printf("Error '%v' occurred while loading configurations.", err)
		return
	}	
	cfg.Start()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /tmp", mid(cfg.tmp))

	mux.HandleFunc("GET /search/{search}", mid(cfg.handleSearch))
	
	mux.HandleFunc("GET /calendar/dividends", cfg.tmp)
	mux.HandleFunc("GET /calendar/earnings", cfg.tmp)
	mux.HandleFunc("GET /calendar/economic-events", cfg.tmp)
	mux.HandleFunc("GET /calendar/ipo", cfg.tmp)
	mux.HandleFunc("GET /calendar/public-offerings", cfg.tmp)
	mux.HandleFunc("GET /calendar/stock-splits", cfg.tmp)

	
	mux.HandleFunc("GET /indicators/adx", cfg.handleIndicatorsADX)
	mux.HandleFunc("GET /indicators/macd", cfg.handleIndicatorsMACD)
	mux.HandleFunc("GET /indicators/rsi", cfg.handleIndicatorRSI) 
	mux.HandleFunc("GET /indicators/sma", cfg.handleIndicatorSMA)

	
	mux.HandleFunc("GET /market/news", cfg.handleMarketNews)
	mux.HandleFunc("GET /market/quotes", cfg.handleMarketQuotes)
	mux.HandleFunc("GET /market/screener", cfg.handleMarketScreener)
	mux.HandleFunc("GET /market/tickers", cfg.handleMarketTickers)

	
	mux.HandleFunc("GET /options", cfg.tmp)
	mux.HandleFunc("GET /options/most-active", cfg.tmp)
	mux.HandleFunc("GET /options/unusual-activity", cfg.tmp)	

	
	mux.HandleFunc("GET /stock/profile/{ticker}", mid(cfg.handleStockProfile))
	mux.HandleFunc("GET /stock/balance-sheet/{ticker}", mid(cfg.handleStockBalanceSheet))
	
//	mux.HandleFunc("GET /stock/calendar-events", cfg.tmp)
	mux.HandleFunc("GET /stock/cashflow-statement/{ticker}", mid(cfg.handleStockCashFlowStatements))

//	mux.HandleFunc("GET /stock/earnings/{ticker}", mid(cfg.handleStockEarnings))
//	mux.HandleFunc("GET /stock/financial-data", cfg.tmp)
//	mux.HandleFunc("GET /stock/history", cfg.tmp)
	mux.HandleFunc("GET /stock/income-statement/{ticker}", mid(cfg.handleStockIncomeStatement))

	mux.HandleFunc("GET /stock/index-trend", cfg.tmp)
	mux.HandleFunc("GET /stock/inside-holders", cfg.tmp)
	mux.HandleFunc("GET /stock/insider-transactions", cfg.tmp)
	mux.HandleFunc("GET /stock/institution-owner", cfg.tmp)
	mux.HandleFunc("GET /stock/modules", cfg.tmp)
	mux.HandleFunc("GET /stock/net-share-purchase-activity", cfg.tmp)
	mux.HandleFunc("GET /stock/recommendation-trend", cfg.tmp)
	mux.HandleFunc("GET /stock/sec-filings", cfg.tmp)
	mux.HandleFunc("GET /stock/statistics", cfg.tmp)
	mux.HandleFunc("GET /stock/upgrade-downgrade-history", cfg.tmp)
	
	srv := &http.Server{
		Addr: ":" + "8080",
		Handler: mux,
	}
	
	log.Printf("Serving on port: %s\n", "8080")
	log.Fatal(srv.ListenAndServe())
}



