package chgkapi

type Packet struct {
	Questions []*Question `xml:"question"`
}
