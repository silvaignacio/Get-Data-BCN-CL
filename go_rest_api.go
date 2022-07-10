package go_rest_api

import (
	"example.com/go-rest-api/controllers"
	"example.com/go-rest-api/controllers/rule"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	NotFoundError = fmt.Errorf("resource could not be found")
)

func Main() {
	router := gin.Default()
	router.Use(
		controllers.ErrorHandler(
			controllers.Map(NotFoundError).ToResponse(func(c *gin.Context, err error) {
				c.Status(http.StatusNotFound)
				c.Writer.Write([]byte(err.Error()))
			}),
		))
	router.LoadHTMLGlob("templates/*")
	router.GET("/data-bcn", rule.GetDataBCN)
	router.GET("/data-bcn/:id", rule.GetRuleDataById)

	router.Run("localhost:8080")
}
