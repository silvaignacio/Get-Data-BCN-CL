package all

import "encoding/xml"

type RulesResult struct {
	XMLName xml.Name `xml:"NORMAS" json:"NORMAS"`
	Rule    []Rule   `xml:"NORMA" json:"NORMA"`
}
