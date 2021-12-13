package test

import (
	"AnalysisService/utils/visualizer"
)

func GameVSGameChart_test() {
	visualizer.GameVsGameChart("game1", "game2", []float32{10.1, 23.4, 45, 6, 13}, []float32{20.17, 29.9, 35, 16.7, 34.6})
}

func AnnualSaleChart_test() {
	visualizer.AnnualSaleChart(2010, 2016, []float32{100, 200, 234.56, 56.5, 650.1, 460, 45})
}
