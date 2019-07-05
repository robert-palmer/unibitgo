package unibit

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	// ...
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// CompanyProfiles ...
// type CompanyProfiles struct {
// 	ID          uint
// 	Ticker      string
// 	Name        string
// 	Website     string
// 	Sector      string
// 	Industry    string
// 	Employees   string
// 	Description string
// }

func dbExgCoverage(db *gorm.DB) {
	exgCoverage := GetExchangeCoverage("NASDAQ")
	for _, ticker := range exgCoverage {
		fmt.Println(ticker.Ticker)
		profile := GetCompanyProfile(ticker.Ticker).Profile
		db.Create(&profile)
	}
}

// StartDB ...
func StartDB() {
	fmt.Println("Connecting to SQL Db ...")
	connStr := "user=postgres password=rober dbname=stocks host=localhost sslmode=disable"
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SingularTable(true)
	dbExgCoverage(db)

	db.AutoMigrate(&CompanyProfile{}) // create table

}
