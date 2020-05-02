package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	chgkGame      = 1 // Что? Где? Когда?
	brainRingGame = 2 // Брейн-ринг
	internetGame  = 3 // Интернет вопросы
	winglessGame  = 4 // Бескрылка
	ownGame       = 5 // Своя игра
	scholarGame   = 2 // Эрудитка
)

type Database interface {
	GetQuestion() (*Question, error)
	GetQuestionPacket(limit int) (*Packet, error)
}

type database struct {
	baseUrl  string
	gameType int
}

func NewDatabase(gameType int) Database {
	return &database{
		baseUrl:  "https://db.chgk.info",
		gameType: gameType,
	}
}

// GetQuestion returns one random question from storage
func (d *database) GetQuestion() (*Question, error) {
	packet, err := d.loadPacket(1)
	if err != nil {
		return nil, err
	}
	if len(packet.Questions) > 0 {
		return packet.Questions[0], nil
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
	return fmt.Sprintf(d.baseUrl+"/xml/random/types%d/limit%d", d.gameType, limit)
}

// request executes http request and read response body
func (d *database) request(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error on get response!\n")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error on read response body!\n")
	}
	return body, nil
}
