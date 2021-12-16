package routes

import (
	"AnalysisService/controllers"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("*.html")
	r.Use(CORSMiddleware())
	r.GET("/graph/region/:gameOne/:gameTwo", controllers.GetGameVsGameChart)
	r.GET("/graph/annual/:firstYear/:lastYear", controllers.GetAnnualSaleChart)
	r.GET("/graph/company/:companyOne/:companyTwo/:firstYear/:lastYear", controllers.GetCompanyVsCompanyChart)
	r.GET("/graph/genre/:firstYear/:lastYear", controllers.GetAllGenresChart)
	return r
}
