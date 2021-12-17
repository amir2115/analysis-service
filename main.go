package main

import (
	"AnalysisService/config"
	"AnalysisService/routes"
)

func main() {
	config.ConnectDatabase()
	r := routes.SetupRouter()

	r.Run("0.0.0.0:8090")
}
