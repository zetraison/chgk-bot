package database

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURL = "https://db.chgk.info"

// Database describes available database functions
type Database interface {
	GetQuestion() (*Question, error)
	GetQuestionPacket(limit int) (*Packet, error)
}

type database struct {
	mode int
}

// NewDatabase returns new database api instance
func NewDatabase(mode int) Database {
	return &database{
		mode: mode,
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

	return nil, errors.New("packet is empty")
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
	data, err := d.request(limit)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(data, &packet)
	if err != nil {
		return nil, err
	}
	return
}

// request executes http request and read response body
func (d *database) request(limit int) ([]byte, error) {
	databaseURL := fmt.Sprintf(baseURL+"/xml/random/types%d/limit%d", d.mode, limit)
	log.Printf("request: %s\n", databaseURL)

	resp, err := http.Get(databaseURL) // #nosec G107
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
