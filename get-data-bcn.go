package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

type Organisms struct {
	XMLName  xml.Name `xml:"ORGANISMOS"`
	Organism string   `xml:"ORGANISMO"`
}

type Rule struct {
	XMLName         xml.Name  `xml:"NORMA"`
	PublishedDate   string    `xml:"FECHA_PUBLICACION"`
	PromulgatedDate string    `xml:"FECHA_PROMULGACION"`
	Title           string    `xml:"TITULO"`
	Text            string    `xml:"TEXTO"`
	Organisms       Organisms `xml:"ORGANISMOS"`
}

type RulesResult struct {
	XMLName xml.Name `xml:"NORMAS"`
	Rule    []Rule   `xml:"NORMA"`
}

type ChileanRulesResponse struct {
	RulesResult string `xml:"RulesResult"`
}

var (
	r ChileanRulesResponse
)

func getDataBCN(c *gin.Context) {
	resp, _ := http.Get(os.Getenv("BCN_URL"))

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	result := RulesResult{}

	_ = xml.Unmarshal([]byte(data), &result)

	for i := 0; i < len(result.Rule); i++ {
		fmt.Println("Título: " + result.Rule[i].Title)
		fmt.Println("Fecha de promulgación: " + result.Rule[i].PromulgatedDate)
		fmt.Println("Organismo: " + result.Rule[i].Organisms.Organism)
		fmt.Println(" ")
	}

}

func main() {
	router := gin.Default()
	router.GET("/data-bcn", getDataBCN)

	router.Run("localhost:8080")
}
