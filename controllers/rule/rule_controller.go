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

	c.HTML(http.StatusOK,
		"index.html",
		gin.H{
			"title": "Normas publicadas",
			"data":  data.Rule,
		})

}

func GetRuleDataById(c *gin.Context) {
	id := c.Param("id")
	resp, _ := http.Get("https://www.leychile.cl/Consulta/obtxml?opt=7&idNorma=" + id)
	defer resp.Body.Close()

	data := utils.ProcessRuleBody(resp.Body)

	c.HTML(http.StatusOK,
		"rule.html",
		gin.H{
			"title":    data.Metadata.Title,
			"subtitle": data.Identifier.Organisms.Organism,
		})
}
