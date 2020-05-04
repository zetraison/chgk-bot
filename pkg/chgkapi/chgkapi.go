package chgkapi

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Database interface {
	GetQuestion() (*Question, error)
	GetQuestionPacket(limit int) (*Packet, error)
}

type database struct {
	baseUrl string
	mode    int
}

func NewDatabase(mode int) Database {
	return &database{
		baseUrl: "https://db.chgk.info",
		mode:    mode,
	}
}

// GetQuestion returns one random question from storage
func (d *database) GetQuestion() (*Question, error) {
	packet, err := d.loadPacket(1)
	if err != nil {
		return nil, err
	}
	if len(packet.Questions) > 0 {
		question := packet.Questions[0]

		log.Print(question.String())

		return question, nil
	}
	return nil, fmt.Errorf("Packet is empty\n")
}

// GetQuestion returns random questions packet from storage with limit
func (d *database) GetQuestionPacket(limit int) (*Packet, error) {
	packet, err := d.loadPacket(limit)
	if err != nil {
		return nil, err
	}
	return packet, nil
}

// load gets question list and wrap to packet
func (d *database) loadPacket(limit int) (packet *Packet, err error) {
	url := d.buildUrl(limit)

	data, err := d.request(url)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(data, &packet)
	if err != nil {
		return nil, err
	}
	return
}

// buildUrl returns http url
func (d *database) buildUrl(limit int) string {
	return fmt.Sprintf(d.baseUrl+"/xml/random/types%d/limit%d", d.mode, limit)
}

// request executes http request and read response body
func (d *database) request(url string) ([]byte, error) {
	resp, err := http.Get(url)

	log.Printf("request: %s\n", url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
