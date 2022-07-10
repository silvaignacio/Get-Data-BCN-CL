package all

import "encoding/xml"

type RulesResult struct {
	XMLName xml.Name `xml:"NORMAS"`
	Rule    []Rule   `xml:"NORMA"`
}
