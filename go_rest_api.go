package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rest-api/controllers"
	"go-rest-api/controllers/rule"
	"log"
	"net/http"
)

var (
	NotFoundError = fmt.Errorf("resource could not be found")
)

func main() {
	router := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	router.Use(
		controllers.ErrorHandler(
			controllers.Map(NotFoundError).ToResponse(func(c *gin.Context, err error) {
				c.Status(http.StatusNotFound)
				c.Writer.Write([]byte(err.Error()))
			}),
		))
	router.Use(gin.Logger())

	router.GET("/data-bcn", rule.GetDataBCN)
	router.GET("/data-bcn/:id", rule.GetRuleDataById)

	router.Run()
}
