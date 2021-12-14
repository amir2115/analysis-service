package main

import "AnalysisService/test"

func main() {
	//config.ConnectDatabase()
	//r := routes.SetupRouter()
	//
	//r.Run("0.0.0.0:8000")
	test.GameVSGameChart_test()
	test.AnnualSaleChart_test()
	test.CompanyVsCompanyChart_test()
	test.AllCategoriesChart_test()
}
