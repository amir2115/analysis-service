package visualizer

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"os"
)

var regionSales = [...]string{"NA_Sales", "EU_Sales", "JP_Sales", "Other_Sales", "Global_Sales"}

func GameVsGameChart(gameOne string, gameTwo string, gameOneSales []float32, gameTwoSales []float32) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    fmt.Sprintf("%s vs %s", gameOne, gameTwo),
		Subtitle: "sales",
	}))
	bar.SetXAxis(regionSales).
		AddSeries("Category A", generateGameVsGameBarItems(gameOneSales)).
		AddSeries("Category B", generateGameVsGameBarItems(gameTwoSales))
	f, _ := os.Create(fmt.Sprintf("%s%s.html", gameOne, gameTwo))
	_ = bar.Render(f)
}
func AnnualSaleChart(firstYear int, LastYear int, totalGameSales []float32) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    fmt.Sprintf("Annual Sales %d - %d", firstYear, LastYear),
		Subtitle: "sales",
	}))
	bar.SetXAxis(annualSalesLabelMaker(firstYear, LastYear)).
		AddSeries("Category A", generateAnnualSaleBarItems(totalGameSales, LastYear-firstYear+1))
	f, _ := os.Create(fmt.Sprintf("annualSales%d-%d.html", firstYear, LastYear))
	_ = bar.Render(f)
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
	for i := 0; i < len(regionSales); i++ {
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