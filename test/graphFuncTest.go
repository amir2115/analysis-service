package test

import (
	"AnalysisService/utils"
)

func GameVSGameChart_test() {
	utils.GameVsGameChart("game1", "game2", []float32{10.1, 23.4, 45, 6, 13}, []float32{20.17, 29.9, 35, 16.7, 34.6})
}

func AnnualSaleChart_test() {
	utils.AnnualSaleChart(2010, 2016, []float32{100, 200, 234.56, 56.5, 650.1, 460, 45})
}

func CompanyVsCompanyChart_test() {
	utils.CompanyVsCompanyChart("company1", "company2", 2000, 2006, []float32{100, 200, 234.56, 56.5, 650.1, 460, 45}, []float32{45, 460, 650.1, 56.5, 234.56, 200, 100})
}
func AllGenresChart_test() {
	utils.AllGenresChart(2000, 2006, []float32{100, 200, 234.56, 56.5, 650.1, 460, 45, 154, 854, 510, 321, 457})
}
