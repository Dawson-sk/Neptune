package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
//	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/cdproto/cdp"
	"golang.org/x/net/html"
)

func (cfg *AppConfig) tmp(w http.ResponseWriter, req *http.Request) {
	log.Println("Passed into next handler")
	return
}

func getNewsFromHTML(rawHTML string) ([]string, error) {
	_, err := html.Parse(strings.NewReader(rawHTML))
	if err != nil {
		return nil, fmt.Errorf("Error: %v occurred while getting the news.")
	}
	return []string{}, nil
	/*
	for n := range body.Descendants() {
		if n.Type == html.ElementNode && n.Data == {
			
		}
	} */
}

// Get stocks that match specified search params.
func getStocksFromHTML(rawHTML string, param string) ([][]string, error) {
	// String: index 1 = tickers, index 2 = names
	_, err := html.Parse(strings.NewReader(rawHTML))
	if err != nil {
		return nil, fmt.Errorf("Error: %v occurred while retrieving stocks", err)
	}
	return [][]string{}, nil
	/*
	// Needs to be able to flip over pages
	for n := body.Descendants() {
	
	} */
}

// ----- Handlers

func checkSelector(ctx context.Context, selector string) (bool, error) {
	var res string
	js := `
		function checkSelector() {
			let res
			try {
				const element = document.evaluate('`+selector+`', document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue;
				res = element ? "exists" : "not found";
			} catch (e) {
				res = "invalid";
			}
			return res;
		}
		checkSelector()
	`
	if err := chromedp.Run(ctx,
		chromedp.Evaluate(js, &res),
	); err != nil {
		log.Println(err)
		return false, err
	}
	
//	log.Println(res)
	return res == "exists", nil
}

// Returning hundreds of null maps
func formatRow(dates []string, textElements []string) (map[string]any, error) {
	var cache string
	var count int
	var results map[string]any
	for _, n := range textElements {
		if matched, _ := regexp.Match(`[A-Za-z]+ [A-Za-z]+|[A-Za-z]+`, []byte(n)); matched {
//			log.Println(n)
			cache = n
			count = 0
		}
	
		// Not matching currency
		if matched, _ := regexp.Match(``, []byte(n)); matched {
			log.Println(n)
			inner, _ := formatCurrency(n)
			innerMap := map[string]Inner{cache:inner}
			results = map[string]any{dates[0]: innerMap}
			count++
		}
	}
	
	return results, nil
}

func formatDate(s string) (string) {
	return s
}

func formatCurrency(s string) (Inner, error) {
	return Inner{}, nil
}

func (cfg *AppConfig) handleSearch(w http.ResponseWriter, req *http.Request) {
	search := req.PathValue("search")
	var results [][]string
	
	for i := range 99 {
		body, err := getHTML(fmt.Sprintf("https://finance.yahoo.com/research-hub/screener/equity/?start=%v?count=%v", i*100, (i*100)+100))
		if err != nil {
			log.Printf("Error '%v' occured while retrieving HTML.", err)
			return
		}
		
		matches, err := getStocksFromHTML(body, search)
		if err != nil {
			log.Printf("Error '%v' occured while searching for matching stocks.", err)
			return
		}
		time.Sleep(time.Second * 1)
		fmt.Println(results, matches, i)
	}
	
	// Stops on 24 requests for some reason
	
	// Format matches
	// Respond with matches
}

// Calendar
func (cfg *AppConfig) handleCalendarDividends(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleCalendarEarnings(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleCalendarEconomicsEvents(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleCalendarIPO(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleCalendarPublicOfferings(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleCalendarStockSplits(w http.ResponseWriter, req *http.Request) {
	return
}

// Indicators
func (cfg *AppConfig) handleIndicatorsADX(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleIndicatorsMACD(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleIndicatorRSI(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleIndicatorSMA(w http.ResponseWriter, req *http.Request) {
	return
}

// Market
func (cfg *AppConfig) handleMarketTickers(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleMarketNews(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleMarketScreener(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleMarketQuotes(w http.ResponseWriter, req *http.Request) {
	return
}

// Options
func (cfg *AppConfig) handleOptions(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleOptionsMostActive(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleOptionsUnusualActivity(w http.ResponseWriter, req *http.Request) {
	return
}

// Stocks
func (cfg *AppConfig) handleStockProfile(w http.ResponseWriter, req *http.Request) {
	type inner struct {
		Raw	string
		Fmt	string
		LongFmt	string
	}

	type officer struct {
		Name		string
		Title		string
		Pay		inner
		Exercised	inner
		YearBorn	string
	}

	type results struct {
		Name			string	`json:"name"`
		Address			string	`json:"address"`
		City			string
		State			string
		Zip			string
		Country			string
		Phone			string
		Website			string
		Sector			string
		Industry		string
		FulltimeEmployees	string
		Description		string
		CompanyOfficers		map[int]officer
	}

	ticker := req.PathValue("ticker")
	
	rawHTML, err := getHTML(fmt.Sprintf("https://finance.yahoo.com/quote/%s/profile/", ticker))
	if err != nil {
		log.Printf("Error occured while getting HTML: %v\n", err)
		return
	}

	body, err := html.Parse(strings.NewReader(rawHTML))
	if err != nil {
		log.Printf("Error occurred while parsing HTML: %v\n", err)
		return
	}
	
	state := regexp.MustCompile(`[A-Z][A-Z] `)
	zip := regexp.MustCompile(`[0-9]+`)
	city := regexp.MustCompile(`[A-Za-z]+ [A-Za-z]+|[A-Za-z]+`)
	var assetProfile []string
	var officers []string
	for n := range body.Descendants() {
		if n.Type == html.ElementNode && (n.Data == "section" || n.Data == "table") {
			for _, a := range n.Attr {
				if (a.Key == "class") && (a.Val == "yf-mj92za") {
					for r := range n.Descendants() {
						if r.Type == html.TextNode {
							officers = append(officers, r.Data)
							// Got proper data coming just needs parsed into map
						}
					}
				}
				
				if (a.Key == "data-testid") && (a.Val == "asset-profile" || a.Val == "description") {
					for r := range n.Descendants() {
						if r.Type == html.TextNode {
							assetProfile = append(assetProfile, r.Data)
						}
					}
				}
			}
			//break
		}
	}

	result := results{
		Name: assetProfile[0],
		Address: assetProfile[3],
		City: city.FindString(assetProfile[5]),
		State: state.FindString(assetProfile[5]),
		Zip: zip.FindString(assetProfile[5]),
		Country: assetProfile[7],
		Phone: assetProfile[9],
		Website: assetProfile[11],
		Sector: assetProfile[15],
		Industry: assetProfile[18],
		FulltimeEmployees: assetProfile[21],
		Description: assetProfile[26],
		CompanyOfficers: make(map[int]officer),
	}
	
	for i,n,t,p,e,y := 0,6,7,8,9,10; i < len(officers); i,n,t,p,e,y = i+1,n+6,t+6,p+6,e+6,y+6 {
		if y > len(officers) {
			break
		}
		result.CompanyOfficers[i] = officer{
			Name: officers[n],
			Title: officers[t],
			Pay: inner{Fmt: officers[p]},
			Exercised: inner{Fmt: officers[e]},
			YearBorn: officers[y],
		}
	}
	
	respondWithJSON(w, http.StatusOK, result)
	return
	// Refactor officer loop?
}

func (cfg *AppConfig) handleStockBalanceSheet(w http.ResponseWriter, req *http.Request) {
// Finished until a solid format of returning data is found.
	cfg.wg.Add(1)
	cfg.pool <- func(){
		defer cfg.wg.Done()
		
		ticker := req.PathValue("ticker")

		ctx, _ := context.WithTimeout(context.Background(), time.Second * 10)
		ctx, _ = chromedp.NewContext(ctx)
		
		if err := chromedp.Run(ctx,
			chromedp.Navigate(fmt.Sprintf("https://finance.yahoo.com/quote/%v/balance-sheet/", ticker)),
			chromedp.Sleep(time.Millisecond * 350),
		); err != nil {
			log.Printf("Error '%v' occurred while navigating to url.", err)
			respondWithError(w, http.StatusInternalServerError, "Error occurred while navigating to url.", err)
			return
		}	


		var i int
		for {
			selector := fmt.Sprintf(`//div[@class="row lv-%v yf-t22klz"][descendant::button]`, i)

			var nodes []*cdp.Node
			if exists, _ := checkSelector(ctx, selector); exists == true {
				if err := chromedp.Run(ctx,
					chromedp.WaitReady(selector, chromedp.BySearch),
					chromedp.Nodes(selector, &nodes, chromedp.BySearch, chromedp.AtLeast(0)),
				); err != nil {
					log.Printf("Error '%v' occurred while getting NodeIDs.", err)
					respondWithError(w, http.StatusInternalServerError, "Error occurred while getting NodeIDs.", err)
					return
				}
			}
			
			log.Println(nodes, len(nodes))
			
			if len(nodes) == 0 {
				break
			}
			
			for _, node := range nodes {
				if err := chromedp.Run(ctx,
					chromedp.WaitVisible(`button`, chromedp.ByQuery, chromedp.FromNode(node)),
					chromedp.Click(`button`, chromedp.ByQuery, chromedp.FromNode(node)),
				); err != nil {
					log.Printf("Error '%v' occurred while clicking buttons.", err)
					respondWithError(w, http.StatusInternalServerError, "Error occurred while clicking buttons.", err)
					return
				}
			}
			i++
		}
			
		var rawHTML string
		if err := chromedp.Run(ctx,
			chromedp.OuterHTML(`//div[@class="table yf-yuwun0"]//div[@class="tableBody yf-yuwun0"]`, &rawHTML, chromedp.BySearch),
		); err != nil {
			log.Printf("Error '%v' occurred while getting html.", err)
			respondWithError(w, http.StatusInternalServerError, "Error occurred while getting html.", err)
		}
		
		body, err := html.Parse(strings.NewReader(rawHTML))
		if err != nil {
			log.Printf("Error occurred while parsing HTML: %v\n", err)
			return
		}
		
//		map[annual][date][header][value]
		
//		var cache string
		results := make(map[string][]any)
		var dates []string
		for n := range body.Descendants() {
			if n.Type == html.ElementNode && n.Data == "div" {
				for _, a := range n.Attr {
					if a.Key == "class" && (strings.Contains(a.Val, "row") && strings.Contains(a.Val, "yf-t22klz")) {						
						var textElements []string
						for r := range n.Descendants() { 
							if r.Type == html.TextNode {
								if matched, _ := regexp.Match(`[0-9]+/[0-9]+/[0-9]+`, []byte(r.Data)); matched {
									dates = append(dates, formatDate(r.Data))
								} else {				
									textElements = append(textElements, r.Data)
								}
							}
						}
						
						innerMap, err := formatRow(dates, textElements)
						if err != nil {
							respondWithError(w, http.StatusInternalServerError, "Error occurred while parsing HTML.", err)
							return
						}
						// Cause of null maps ???
						results["annual"] = append(results["annual"], innerMap)
					}
				}
			}
		}
		respondWithJSON(w, http.StatusOK, results)
		return
	}
	cfg.wg.Wait()
}

func (cfg *AppConfig) handleStockCalendarEvents(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleStockCashFlowStatements(w http.ResponseWriter, req *http.Request) {
// Finished until a solid format of returning data is found.
	cfg.wg.Add(1)
	cfg.pool <- func(){
		defer cfg.wg.Done()
		ticker := req.PathValue("ticker")

		ctx, _ := context.WithTimeout(context.Background(), time.Second * 10)
		ctx, _ = chromedp.NewContext(ctx)
		
		if err := chromedp.Run(ctx,
			chromedp.Navigate(fmt.Sprintf("https://finance.yahoo.com/quote/%v/cash-flow/", ticker)),
			chromedp.Sleep(time.Millisecond * 350),
		); err != nil {
			log.Printf("Error '%v' occurred while navigating to url.", err)
			respondWithError(w, http.StatusInternalServerError, "Error occurred while navigating to url.", err)
			return
		}	


		var i int
		for {
			selector := fmt.Sprintf(`//div[@class="row lv-%v yf-t22klz"][descendant::button]`, i)

			var nodes []*cdp.Node
			if exists, _ := checkSelector(ctx, selector); exists == true {
				if err := chromedp.Run(ctx,
					chromedp.WaitReady(selector, chromedp.BySearch),
					chromedp.Nodes(selector, &nodes, chromedp.BySearch, chromedp.AtLeast(0)),
				); err != nil {
					log.Printf("Error '%v' occurred while getting NodeIDs.", err)
					respondWithError(w, http.StatusInternalServerError, "Error occurred while getting NodeIDs.", err)
					return
				}
			}
			
			log.Println(nodes, len(nodes))
			
			if len(nodes) == 0 {
				break
			}
			
			for _, node := range nodes {
				if err := chromedp.Run(ctx,
					chromedp.WaitVisible(`button`, chromedp.ByQuery, chromedp.FromNode(node)),
					chromedp.Click(`button`, chromedp.ByQuery, chromedp.FromNode(node)),
				); err != nil {
					log.Printf("Error '%v' occurred while clicking buttons.", err)
					respondWithError(w, http.StatusInternalServerError, "Error occurred while clicking buttons.", err)
					return
				}
			}
			i++
		}
			
		var rawHTML string
		if err := chromedp.Run(ctx,
			chromedp.OuterHTML(`//div[@class="table yf-yuwun0"]//div[@class="tableBody yf-yuwun0"]`, &rawHTML, chromedp.BySearch),
		); err != nil {
			log.Printf("Error '%v' occurred while getting html.", err)
			respondWithError(w, http.StatusInternalServerError, "Error occurred while getting html.", err)
		}
		
		body, err := html.Parse(strings.NewReader(rawHTML))
		if err != nil {
			log.Printf("Error occurred while parsing HTML: %v\n", err)
			return
		}
		
		var textElements []string
		for n := range body.Descendants() {
			if n.Type == html.ElementNode && n.Data == "div" {
				for _, a := range n.Attr {
					if a.Key == "class" && (strings.Contains(a.Val, "row") && strings.Contains(a.Val, "yf-t22klz")) {
						for r := range n.Descendants() { 
							if r.Type == html.TextNode {
								textElements = append(textElements, r.Data)
							}
						}
					}
				}
			}
		}

		for i, n := range textElements {
			log.Println(i, n)
		}

		respondWithJSON(w, http.StatusOK, map[string]string{"test": "success"})
		return
	}
	cfg.wg.Wait()
}

func (cfg *AppConfig) handleStockEarnings(w http.ResponseWriter, req *http.Request) {
	// Net Profits
	ticker := req.PathValue("ticker")
	
	rawHTML, err := getHTML(fmt.Sprintf("https://finance.yahoo.com/quote/%s/key-statistics", ticker))
	if err != nil {
		log.Printf("Error occured while getting HTML: %v\n", err)
		return
	}
	
	body, err := html.Parse(strings.NewReader(rawHTML))
	if err != nil {
		log.Printf("Error occurred while parsing HTML: %v\n", err)
		return
	}
	
	var isEarnings bool
	var textElements []string
	for n := range body.Descendants() {
		if n.Type == html.ElementNode && n.Data == "tr" {
			for r := range n.Descendants() {
				if r.Type == html.TextNode && r.Data == "Gross Profit  (ttm)" {
					textElements = append(textElements, fmt.Sprintf(r.Data))
					isEarnings = true
					continue
				}
				
				if match, _ := regexp.MatchString(`[0-9]+\.[0-9][A-Za-z]`, r.Data); r.Type == html.TextNode && match && isEarnings {
					textElements = append(textElements, fmt.Sprintf(r.Data))					
					isEarnings = false
					
					results := map[string]string{
						textElements[0]: textElements[1],
					}
					
					respondWithJSON(w, http.StatusOK, results)
				}
			}
		}
	}
}

func (cfg *AppConfig) handleStockFinancialData(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleStockHistory(w http.ResponseWriter, req *http.Request) {
	//https://finance.yahoo.com/quote/NVDA/history/
	// ????
	return
}

func (cfg *AppConfig) handleStockIncomeStatement(w http.ResponseWriter, req *http.Request) {
// Finished until a solid format of returning data is found.
	cfg.wg.Add(1)
	cfg.pool <- func(){
		defer cfg.wg.Done()
		ticker := req.PathValue("ticker")

		ctx, _ := context.WithTimeout(context.Background(), time.Second * 10)
		ctx, _ = chromedp.NewContext(ctx)
		
		if err := chromedp.Run(ctx,
			chromedp.Navigate(fmt.Sprintf("https://finance.yahoo.com/quote/%v/financials/", ticker)),
			chromedp.Sleep(time.Millisecond * 350),
		); err != nil {
			log.Printf("Error '%v' occurred while navigating to url.", err)
			respondWithError(w, http.StatusInternalServerError, "Error occurred while navigating to url.", err)
			return
		}	


		var i int
		for {
			selector := fmt.Sprintf(`//div[@class="row lv-%v yf-t22klz"][descendant::button]`, i)

			var nodes []*cdp.Node
			if exists, _ := checkSelector(ctx, selector); exists == true {
				if err := chromedp.Run(ctx,
					chromedp.WaitReady(selector, chromedp.BySearch),
					chromedp.Nodes(selector, &nodes, chromedp.BySearch, chromedp.AtLeast(0)),
				); err != nil {
					log.Printf("Error '%v' occurred while getting NodeIDs.", err)
					respondWithError(w, http.StatusInternalServerError, "Error occurred while getting NodeIDs.", err)
					return
				}
			}
			
			log.Println(nodes, len(nodes))
			
			if len(nodes) == 0 {
				break
			}
			
			for _, node := range nodes {
				if err := chromedp.Run(ctx,
					chromedp.WaitVisible(`button`, chromedp.ByQuery, chromedp.FromNode(node)),
					chromedp.Click(`button`, chromedp.ByQuery, chromedp.FromNode(node)),
				); err != nil {
					log.Printf("Error '%v' occurred while clicking buttons.", err)
					respondWithError(w, http.StatusInternalServerError, "Error occurred while clicking buttons.", err)
					return
				}
			}
			i++
		}
			
		var rawHTML string
		if err := chromedp.Run(ctx,
			chromedp.OuterHTML(`//div[@class="table yf-yuwun0"]//div[@class="tableBody yf-yuwun0"]`, &rawHTML, chromedp.BySearch),
		); err != nil {
			log.Printf("Error '%v' occurred while getting html.", err)
			respondWithError(w, http.StatusInternalServerError, "Error occurred while getting html.", err)
		}
		
		body, err := html.Parse(strings.NewReader(rawHTML))
		if err != nil {
			log.Printf("Error occurred while parsing HTML: %v\n", err)
			return
		}
		
		var results map[string][]any
		var dates []string
		for n := range body.Descendants() {
			if n.Type == html.ElementNode && n.Data == "div" {
				for _, a := range n.Attr {
					if a.Key == "class" && (strings.Contains(a.Val, "row") && strings.Contains(a.Val, "yf-t22klz")) {
						var textElements []string
						for r := range n.Descendants() { 
							if r.Type == html.TextNode {
								textElements = append(textElements, r.Data)
							}
						}
						innerMap, err := formatRow(dates, textElements)
						if err != nil {
							respondWithError(w, http.StatusInternalServerError, "Error occurred while parsing HTML", err)
							return
						}
						results["annual"] = append(results["annual"], innerMap)
					}
				}
			}
		}

		respondWithJSON(w, http.StatusOK, map[string]string{"test": "success"})
		return
	}
	cfg.wg.Wait()
}

func (cfg *AppConfig) handleStockIndexTrend(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleStockInsideHolders(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleStockInsiderTransactions(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleStockInstitutionOwner(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleStockModules(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleStockNetSharePurchaseActivity(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleStockRecommendationTrend(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleStockSECFilings(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleStockStatistics(w http.ResponseWriter, req *http.Request) {
	return
}

func (cfg *AppConfig) handleStockUpgradeDowngradeHistory(w http.ResponseWriter, req *http.Request) {
	return
}





