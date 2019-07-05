package main

import (
	uni "github.com/robert-palmer/unibitgo/pkg/unibit"
)

func main() {
	// fmt.Println(uni.GetCurrentPrice("AMD"))
	// fmt.Println(uni.GetCompanyProfile("CSX"))
	// fmt.Println(uni.GetFinancialSummary("AMD"))
	// exchanges := []string{"NASDAQ", "NYSE", "LSE", "HKEX"}
	// for _, ex := range exchanges {
	// 	exCov := uni.GetExchangeCoverage(ex)
	// 	fmt.Println(len(exCov))

	// }
	// exgCoverage := uni.GetExchangeCoverage("NASDAQ")
	// for i, ticker := range exgCoverage {
	// 	fmt.Println(i, ticker.Ticker)
	// 	fmt.Println(uni.GetCompanyProfile(ticker.Ticker))
	// }
	uni.StartDB()
}
