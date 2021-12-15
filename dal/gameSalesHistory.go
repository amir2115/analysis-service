package dal

import (
	"AnalysisService/config"
	"AnalysisService/utils"
)

func GetGameRegionSales(gameName string) []float32 {
	sales := make([]float32, 0)
	for _, region := range utils.RegionNames {
		var temp float32
		config.DB.Table("GameSalesHistory").Select("SUM(?)", region).Where("name = ?", gameName).Scan(&temp)
		sales = append(sales, temp)
	}
	return sales
}

func GetAnnualSales(firstYear int, lastYear int) []float32 {
	sales := make([]float32, 0)
	for year := firstYear; year <= lastYear; year++ {
		var temp float32
		config.DB.Table("GameSalesHistory").Select("SUM(global_sales)").Where("year = ?", year).Scan(&temp)
		sales = append(sales, temp)
	}
	return sales
}

func GetCompanySales(companyName string, firstYear int, lastYear int) []float32 {
	sales := make([]float32, 0)
	for year := firstYear; year <= lastYear; year++ {
		var temp float32
		config.DB.Table("GameSalesHistory").Select("SUM(global_sales)").Where("year = ? AND publisher = ?", year, companyName).Scan(&temp)
		sales = append(sales, temp)
	}
	return sales
}

func GetGenreSales(firstYear int, lastYear int) []float32 {
	sales := make([]float32, 0)
	for _, genre := range utils.Genres {
		var sum float32
		sum = 0
		for year := firstYear; year <= lastYear; year++ {
			var temp float32
			config.DB.Table("GameSalesHistory").Select("SUM(global_sales)").Where("year = ? AND genre = ?", year, genre).Scan(&temp)
			sum += temp
		}
		sales = append(sales, sum)
	}
	return sales
}
