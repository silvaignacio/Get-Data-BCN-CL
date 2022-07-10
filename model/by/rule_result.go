package by

import (
	"encoding/xml"
)

type Organisms struct {
	XMLName  xml.Name `xml:"Organismos"`
	Organism string   `xml:"Organismo"`
}

type Identifier struct {
	XMLName         xml.Name  `xml:"Identificador"`
	PromulgatedDate string    `xml:"fechaPromulgacion,attr"`
	PublicatedDate  string    `xml:"fechaPublicacion"`
	Organisms       Organisms `xml:"Organismos"`
}

type Subjects struct {
	XMLName xml.Name `xml:"Materias"`
	Subject string   `xml:"Materia"`
}

type Metadata struct {
	XMLName          xml.Name `xml:"Metadatos"`
	Title            string   `xml:"TituloNorma"`
	Subjects         string   `xml:"Materias"`
	SourceIdentifier string   `xml:"IdentificacionFuente"`
	SourceNumber     int      `xml:"NumeroFuente"`
}

type Head struct {
	XMLName     xml.Name `xml:"Encabezado"`
	VersionDate string   `xml:"fechaVersion,attr"`
	Derogated   string   `xml:"derogado,attr"`
	Text        string   `xml:"Texto"`
}

type RuleResult struct {
	XMLName     xml.Name   `xml:"Norma"`
	Derogate    string     `xml:"derogado,attr"`
	VersionDate string     `xml:"fechaVersion,attr"`
	Identifier  Identifier `xml:"Identificador"`
	Metadata    Metadata   `xml:"Metadatos"`
	Head        Head       `xml:"Encabezado"`
}
