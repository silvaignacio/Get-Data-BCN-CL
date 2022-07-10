package all

import "encoding/xml"

type Rule struct {
	XMLName         xml.Name  `xml:"NORMA"`
	ID              string    `xml:"idNorma,attr"`
	PublishedDate   string    `xml:"FECHA_PUBLICACION"`
	PromulgatedDate string    `xml:"FECHA_PROMULGACION"`
	Title           string    `xml:"TITULO"`
	Text            string    `xml:"TEXTO"`
	Organisms       Organisms `xml:"ORGANISMOS"`
}
