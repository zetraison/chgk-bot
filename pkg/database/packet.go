package database

// Packet struct that contains questions slice
// Used for parse xml representation of packet of questions
type Packet struct {
	Questions []*Question `xml:"question"`
}
