package all

import "encoding/xml"

type Rule struct {
	XMLName         xml.Name  `xml:"NORMA" json:"NORMA"`
	ID              string    `xml:"idNorma,attr" json:"idNorma,attr"`
	PublishedDate   string    `xml:"FECHA_PUBLICACION" json:"FECHA_PUBLICACION"`
	PromulgatedDate string    `xml:"FECHA_PROMULGACION" json:"FECHA_PROMULGACION"`
	Title           string    `xml:"TITULO" json:"TITULO"`
	Text            string    `xml:"TEXTO" json:"TEXTO"`
	Organisms       Organisms `xml:"ORGANISMOS" json:"ORGANISMOS"`
}
