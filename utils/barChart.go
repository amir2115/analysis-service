package utils

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"os"
)

func GameVsGameChart(gameOne string, gameTwo string, gameOneSales []float32, gameTwoSales []float32) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    fmt.Sprintf("%s vs %s", gameOne, gameTwo),
		Subtitle: "sales",
	}))
	bar.SetXAxis(RegionNames).
		AddSeries("Category A", generateGameVsGameBarItems(gameOneSales)).
		AddSeries("Category B", generateGameVsGameBarItems(gameTwoSales))
	f, _ := os.Create(fmt.Sprintf("%s%s.html", gameOne, gameTwo))
	_ = bar.Render(f)
}
func AnnualSaleChart(firstYear int, LastYear int, totalGameSales []float32) {
	bar := charts.NewBar()
	bar.XYReversal()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    fmt.Sprintf("Annual Sales %d - %d", firstYear, LastYear),
		Subtitle: "sales",
	}))
	bar.SetXAxis(annualSalesLabelMaker(firstYear, LastYear)).
		AddSeries("Category A", generateAnnualSaleBarItems(totalGameSales, LastYear-firstYear+1))
	f, _ := os.Create(fmt.Sprintf("annualSales%d-%d.html", firstYear, LastYear))
	_ = bar.Render(f)
}

func CompanyVsCompanyChart(companyOne string, companyTwo string, firstYear int, LastYear int, companyOneSales []float32, companyTwoSales []float32) {
	bar := charts.NewBar()
	bar.XYReversal()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    fmt.Sprintf("%s vs %s", companyOne, companyTwo),
		Subtitle: "sales",
	}))
	bar.SetXAxis(annualSalesLabelMaker(firstYear, LastYear)).
		AddSeries("Category A", generateCompanyVsCompanyBarItems(companyOneSales, LastYear-firstYear+1)).
		AddSeries("Category B", generateCompanyVsCompanyBarItems(companyTwoSales, LastYear-firstYear+1))
	f, _ := os.Create(fmt.Sprintf("%s%s-%d-%d.html", companyOne, companyTwo, firstYear, LastYear))
	_ = bar.Render(f)
}
func AllGenresChart(firstYear int, LastYear int, categoryGameSales []float32) {
	bar := charts.NewBar()
	bar.XYReversal()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    fmt.Sprintf("Category Sales %d - %d", firstYear, LastYear),
		Subtitle: "sales",
	}))

	bar.SetXAxis(Genres).
		AddSeries("Category A", generateAllGenresBarItems(categoryGameSales))
	f, _ := os.Create(fmt.Sprintf("genres-%d-%d.html", firstYear, LastYear))
	_ = bar.Render(f)
}
func generateAllGenresBarItems(gameSales []float32) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < len(Genres); i++ {
		items = append(items, opts.BarData{Value: gameSales[i], Label: &opts.Label{Show: true}})
	}
	return items
}
func generateCompanyVsCompanyBarItems(gameSales []float32, yearCount int) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < yearCount; i++ {
		items = append(items, opts.BarData{Value: gameSales[i], Label: &opts.Label{Show: true}})
	}
	return items
}
func generateAnnualSaleBarItems(gameSales []float32, yearCount int) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < yearCount; i++ {
		items = append(items, opts.BarData{Value: gameSales[i], Label: &opts.Label{Show: true}})
	}
	return items
}

func generateGameVsGameBarItems(gameSales []float32) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < len(RegionNames); i++ {
		items = append(items, opts.BarData{Value: gameSales[i], Label: &opts.Label{Show: true}})
	}
	return items
}

func annualSalesLabelMaker(firstYear int, LastYear int) []string {
	labels := make([]string, 0)

	for year := firstYear; year <= LastYear; year++ {
		labels = append(labels, fmt.Sprintf("%d", year))
	}
	return labels
}
