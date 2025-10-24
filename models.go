package main

import ()

type Inner struct {
	Raw			string		`json:"raw"`
	Fmt			string		`json:"fmt"`
	LongFmt			string		`json:"longFmt"`
}

// ---- Calendar Data Models ---- //

type Dividends struct{}

type Earnings struct{}

type EconomicEvents struct{}

type IPO struct{}

type PublicOfferings struct{}

type StockSplits struct{}



// ---- Indicator Data Models ---- //

type ADX struct{}

type MACD struct{}

type RSI struct{}

type SMA struct{}



// ---- Market Data Models ---- //

type News struct{}

type Quotes struct{}

type Screener struct{}

type Tickers struct{}



// ---- Options Data Models ---- //

type Options struct{}

type MostActive struct{}

type UnusualActivity struct{}



// ---- Stock Data Models ---- //

type StockSearch struct{}

type StockProfile struct {
	Name			string		`json:"name"`
	Address			string		`json:"address"`
	City			string		`json:"city"`
	State			string		`json:"state"`
	Zip			string		`json:"zip"`
	Country			string		`json:"country"`
	Phone			string		`json:"phone"`
	Website			string		`json:"website"`
	Sector			string		`json:"sector"`
	Industry		string		`json:"industry"`
	FulltimeEmployees	string		`json:"fulltime_employees"`
	Description		string		`json:"description"`
	CompanyOfficers		map[int]struct {
		Name			string		`json:"name"`
		Title			string		`json:"title"`
		Pay			Inner		`json:"pay"`
		Exercised		Inner		`json:"exercised"`
		YearBorn		string		`json:"year_born"`	
	} `json:"officers"`
}


type StockBalanceSheet struct{
	TotalAssets struct {
		CurrentAssets struct {
			Receivables struct {
				AccountsReceivable 			string
			}
			Inventory struct {
				RawMaterials 				string
				WorkInProcess 				string
				FinishedGoods 				string
			}
			CashAndCashEquivalent				string
			OtherShortTermInvestments			string
			PrepaidAssets					string
			OtherCurrentAssets				string
		}
		TotalNonCurrentAssets struct {
			NetPPE struct {
				GrossPPE struct {
					Properties			string
					LandAndImprovements		string
					BuildingsAndImprovements	string
					MachineryFurnitureEquipment	string
					OtherProperties			string
					ConstructionInProgress		string
				}
				AccumulatedDepreciation 		string
			}
			
		}
	} `json:"totalAssets"`
	TotalLiabilitiesNetMinorityInterest struct {
		CurrentLiabilities struct {
		}
		TotalNonCurrentLiabilitiesNetMinorityInterest struct {
		}
	} `json:"TotalLiabilitiesNetMinorityInterest"`
	TotalEquityGrossMinorityInterest struct {
		StockholdersEquity struct {
		} `json:"StockholdersEquity"`
	} `json:"TotalEquityGrossMinorityInterest"`
	TotalCapitalization     string `json:"TotalCapitalization"`
	CommonStockEquity       string `json:"CommonStockEquity"`
	CapitalLeaseObligations string `json:"CapitalLeaseObligations"`
	NetTangibleAssets       string `json:"NetTangibleAssets"` 
	WorkingCapital          string `json:"WorkingCapital"`
	InvestedCapital         string `json:"InvestedCapital"`
	TangibleBookValue       string `json:"TangibleBookValue"`
	TotalDebt               string `json:"TotalDebt"`
	NetDebt                 string `json:"NetDebt"`
	ShareIssued             string `json:"ShareIssued"`
	OrdinarySharesNumber    string `json:"OrdinarySharesNumber"`
}

type StockCalendarEvents struct{}

type StockCashflowStatement struct{}

//type Earnings struct{}

//type FinancialData struct{}

type StockHistory struct{}

type StockIncomeStatement struct{}

type StockIndexTrend struct{}

type StockInsideHolders struct{}

type StockInsiderTransactions struct{}

type StockInstitutionOwner struct{}

type StockModules struct{}

type StockNetSharePurchaseActivity struct{}

type StockRecommendationTrend struct{}

type StockSECFilings struct{}

type StockStatistics struct{}

type StockUpgradeDowngradeHistory struct{}
