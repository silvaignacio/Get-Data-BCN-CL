package all

import "encoding/xml"

type Organisms struct {
	XMLName  xml.Name `xml:"ORGANISMOS" xml:"ORGANISMOS"`
	Organism string   `xml:"ORGANISMO" xml:"ORGANISMO"`
}
