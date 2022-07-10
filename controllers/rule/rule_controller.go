package rule

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/utils"
	"net/http"
	"os"
)

func GetDataBCN(c *gin.Context) {

	resp, _ := http.Get(os.Getenv("BCN_URL"))

	defer resp.Body.Close()

	data := utils.ProcessRulesBody(resp.Body)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": data, // cast it to string before showing
	})

}

func GetRuleDataById(c *gin.Context) {
	id := c.Param("id")
	resp, _ := http.Get("https://www.leychile.cl/Consulta/obtxml?opt=7&idNorma=" + id)
	defer resp.Body.Close()

	data := utils.ProcessRuleBody(resp.Body)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": data, // cast it to string before showing
	})
}
