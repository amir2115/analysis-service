package main

import (
	"AnalysisService/config"
	"AnalysisService/routes"
)

func main() {
	config.ConnectDatabase()
	r := routes.SetupRouter()

	r.Run("0.0.0.0:8090")
	//test.GameVSGameChart_test()
	//test.AnnualSaleChart_test()
	//test.CompanyVsCompanyChart_test()
	//test.AllGenresChart_test()
}
