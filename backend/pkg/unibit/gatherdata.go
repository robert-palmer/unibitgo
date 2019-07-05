package unibit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// APIKey ...
const (
	APIKey    = "UWOHOOPP9I7fP2fO3alOoDHPBmb8bPLg"
	baseURL   = "https://api.unibit.ai/api/"
	accessKey = "&AccessKey=" + APIKey
)

type (
	// CurrentPrice ...
	CurrentPrice struct {
		Price []RealTimeStockPrice `json:"Realtime Stock price"`
	}
	// RealTimeStockPrice ...
	RealTimeStockPrice struct {
		Ticker       string  `json:"ticker"`
		Date         string  `json:"date"`
		Minute       string  `json:"minute"`
		Volume       int     `json:"volume"`
		Price        float32 `json:"price"`
		Timezone     string  `json:"timezone"`
		TimezoneName string  `json:"timezone_name"`
	}
	// MetaData ...
	MetaData struct {
		Ticker        string `json:"ticker"`
		Datapoints    int    `json:"datapoints"`
		CreditCost    int    `json:"credit cost"`
		Timezone      string `json:"timezone"`
		LastRefreshed string `json:"last refreshed"`
	}
)

// CompanyProfiles ...
type CompanyProfiles struct {
	Profile CompanyProfile `json:"company profile"`
}

// CompanyProfile ...
type CompanyProfile struct {
	Ticker      string `json:"ticker"`
	Name        string `json:"company_name"`
	Website     string `json:"website"`
	Sector      string `json:"sector"`
	Industry    string `json:"industry"`
	Employees   string `json:"employee_number"`
	Description string `json:"company_description"`
}

// FinancialSummary - Company financial summary
type FinancialSummary struct {
	CompanyFinancialsSummary struct {
		Open                 string  `json:"open"`
		PreviousClose        string  `json:"previous_close"`
		MarketCap            string  `json:"market_cap"`
		ForwardDividendYield string  `json:"forward_dividend_yield"`
		AvgVolume            string  `json:"avg_volume"`
		ExDividendDate       string  `json:"ex_dividend_date"`
		Beta                 string  `json:"beta"`
		PE                   string  `json:"pe_ratio"`
		EPS                  string  `json:"eps"`
		Week52Range          string  `json:"week_52_range"`
		EarningsDate         string  `json:"earnings_date"`
		Sentiment            float32 `json:"sentiment"`
	} `json:"Company financials summary"`
}

type (
	// ExchangeCoverage - holds all the companies on a given exchange
	ExchangeCoverage []AssetCoverage
	// AssetCoverage - information for each company on a given exchange
	AssetCoverage struct {
		Ticker        string `json:"ticker,omitempty"`
		CompanyName   string `json:"companyName,omitempty"`
		Exchange      string `json:"exchange,omitempty"`
		ExchangeShort string `json:"exchangeShort,omitempty"`
		Currency      string `json:"currency,omitempty"`
		Timezone      string `json:"timezone,omitempty"`
	}
)

// assembles data into imputed struct (company interface{})
func getData(url *string, company interface{}) {
	res, err := http.Get(*url)
	if err != nil {
		fmt.Println(err)
		// log.Fatal(err)
		return
	}
	defer res.Body.Close()

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		// log.Fatal(err)
		return
	}
	err = json.Unmarshal(resData, &company)
	if err != nil {
		fmt.Println(err)
		// log.Fatal(err)
		return
	}
}

// GetCurrentPrice ...
func GetCurrentPrice(s string) CurrentPrice {
	url := baseURL + "realtimestock/" + s + "?datatype=json&size=1" + accessKey
	var company CurrentPrice

	getData(&url, &company)

	return company
}

// GetCompanyProfile ...
func GetCompanyProfile(s string) CompanyProfiles {
	url := baseURL + "companyprofile/" + s + "?datatype=json" + accessKey
	var company CompanyProfiles

	getData(&url, &company)

	return company
}

// GetFinancialSummary ...
func GetFinancialSummary(s string) FinancialSummary {
	url := baseURL + "financials/summary/" + s + "?datatype=json" + accessKey
	var company FinancialSummary

	getData(&url, &company)

	return company
}

// GetExchangeCoverage - show all the stocks available on a given exchange
func GetExchangeCoverage(s string) ExchangeCoverage {
	url := baseURL + "companylist" + "?exchange=" + s + accessKey
	var company ExchangeCoverage

	getData(&url, &company)

	return company
}
