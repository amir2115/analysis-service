package controllers

import (
	"AnalysisService/dal"
	"AnalysisService/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetGameVsGameChart(c *gin.Context) {
	gameOne := c.Param("gameOne")
	gameTwo := c.Param("gameTwo")
	gameOneSales := dal.GetGameRegionSales(gameOne)
	gameTwoSales := dal.GetGameRegionSales(gameTwo)
	chartFileName := utils.GameVsGameChart(gameOne, gameTwo, gameOneSales, gameTwoSales)
	c.HTML(http.StatusOK, chartFileName, nil)
}

func GetAnnualSaleChart(c *gin.Context) {
	_, firstYear := utils.GetParamUInt(c, "firstYear")
	_, lastYear := utils.GetParamUInt(c, "lastYear")
	totalSales := dal.GetAnnualSales(firstYear, lastYear)
	chartFileName := utils.AnnualSaleChart(firstYear, lastYear, totalSales)
	c.HTML(http.StatusOK, chartFileName, nil)
}

func GetCompanyVsCompanyChart(c *gin.Context) {
	companyOne := c.Param("companyOne")
	companyTwo := c.Param("companyTwo")
	_, firstYear := utils.GetParamUInt(c, "firstYear")
	_, lastYear := utils.GetParamUInt(c, "lastYear")
	companyOneSales := dal.GetCompanySales(companyOne, firstYear, lastYear)
	companyTwoSales := dal.GetCompanySales(companyTwo, firstYear, lastYear)
	chartFileName := utils.CompanyVsCompanyChart(companyOne, companyTwo, firstYear, lastYear, companyOneSales, companyTwoSales)
	c.HTML(http.StatusOK, chartFileName, nil)
}

func GetAllGenresChart(c *gin.Context) {
	_, firstYear := utils.GetParamUInt(c, "firstYear")
	_, lastYear := utils.GetParamUInt(c, "lastYear")
	genreSales := dal.GetGenreSales(firstYear, lastYear)
	chartFileName := utils.AllGenresChart(firstYear, lastYear, genreSales)
	c.HTML(http.StatusOK, chartFileName, nil)
}
