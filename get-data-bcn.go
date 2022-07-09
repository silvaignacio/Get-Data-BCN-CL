package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Organismos struct {
	XMLName   xml.Name `xml:"ORGANISMOS"`
	Organismo string   `xml:"ORGANISMO"`
}

type Norma struct {
	XMLName           xml.Name   `xml:"NORMA"`
	FechaPublicacion  string     `xml:"FECHA_PUBLICACION"`
	FechaPromulgacion string     `xml:"FECHA_PROMULGACION"`
	Titulo            string     `xml:"TITULO"`
	Texto             string     `xml:"TEXTO"`
	Organismos        Organismos `xml:"ORGANISMOS"`
}

type NormasResult struct {
	XMLName xml.Name `xml:"NORMAS"`
	Norma   []Norma  `xml:"NORMA"`
}

type NormaChilenaResponse struct {
	NormasResult string `xml:"NormasResult"`
}

var (
	r NormaChilenaResponse
)

func getDataBCN(c *gin.Context) {
	resp, _ := http.Get("https://www.leychile.cl/Consulta/obtxml?opt=3&cantidad=5")

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	result := NormasResult{}

	_ = xml.Unmarshal([]byte(data), &result)

	for i := 0; i < len(result.Norma); i++ {
		fmt.Println("Título: " + result.Norma[i].Titulo)
		fmt.Println("Fecha de promulgación: " + result.Norma[i].FechaPromulgacion)
		fmt.Println("Organismo: " + result.Norma[i].Organismos.Organismo)
		fmt.Println(" ")
	}

}

func main() {
	router := gin.Default()
	router.GET("/data-bcn", getDataBCN)

	router.Run("localhost:8080")
}
