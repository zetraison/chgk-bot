package main

type Question struct {
	Question     string `xml:"Question"`
	Answer       string `xml:"Answer"`
	PassCriteria string `xml:"PassCriteria"`
	Authors      string `xml:"Authors"`
	Sources      string `xml:"Sources"`
	Comments     string `xml:"Comments"`
	Tournament   string `xml:"tournamentTitle"`
}

type Packet struct {
	Questions []*Question `xml:"question"`
}
