package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Payload struct {
	Analysis string `json:"analysis"`
	Advice   string `json:"advice"`
	PowerUp  string `json:"power_up"`
	CoolDown string `json:"cool_down"`
}

type Aries struct {
	Payload
}

type Taurus struct {
	Payload
}

type Gemini struct {
	Payload
}

type Cancer struct {
	Payload
}

type Leo struct {
	Payload
}

type Virgo struct {
	Payload
}

type Libra struct {
	Payload
}

type Scorpio struct {
	Payload
}

type Sagittarius struct {
	Payload
}

type Capricorn struct {
	Payload
}

type Aquarius struct {
	Payload
}

type Pisces struct {
	Payload
}

type ShiitakeResponse struct {
	Aries       `json:"aries"`
	Taurus      `json:"taurus"`
	Gemini      `json:"gemini"`
	Cancer      `json:"cancer"`
	Leo         `json:"leo"`
	Virgo       `json:"virgo"`
	Libra       `json:"libra"`
	Scorpio     `json:"scorpio"`
	Sagittarius `json:"sagittarius"`
	Capricorn   `json:"capricorn"`
	Aquarius    `json:"aquarius"`
	Pisces      `json:"pisces"`
}

func fetchFortuneTelling(date time.Time) (ShiitakeResponse, error) {
	ShiitakeResponse := ShiitakeResponse{}
	url := BaseUrl + date.Format("20060102") + ".json"
	response, err := http.Get(url)
	if err != nil {
		return ShiitakeResponse, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ShiitakeResponse, err
	}

	if err := json.Unmarshal(body, &ShiitakeResponse); err != nil {
		return ShiitakeResponse, err
	}

	return ShiitakeResponse, nil
}

func (ShiitakeResponse *ShiitakeResponse) showFortuneTellingByConstellation(constellation string) error {
	switch constellation {
	case "aries":
		ShiitakeResponse.Aries.Payload.showMessage()
	case "taurus":
		ShiitakeResponse.Taurus.Payload.showMessage()
	case "gemini":
		ShiitakeResponse.Gemini.Payload.showMessage()
	case "cancer":
		ShiitakeResponse.Cancer.Payload.showMessage()
	case "leo":
		ShiitakeResponse.Leo.Payload.showMessage()
	case "virgo":
		ShiitakeResponse.Virgo.Payload.showMessage()
	case "libra":
		ShiitakeResponse.Libra.Payload.showMessage()
	case "scorpio":
		ShiitakeResponse.Scorpio.Payload.showMessage()
	case "sagittarius":
		ShiitakeResponse.Sagittarius.Payload.showMessage()
	case "capricorn":
		ShiitakeResponse.Capricorn.Payload.showMessage()
	case "aquarius":
		ShiitakeResponse.Aquarius.Payload.showMessage()
	case "pisces":
		ShiitakeResponse.Pisces.Payload.showMessage()
	default:
		fmt.Println("何かおかしい")
	}
	return nil
}

func (payload Payload) showMessage() {
	fmt.Println("-----------分析結果-----------")
	fmt.Println(payload.Analysis)
	fmt.Println("----------アドバイス----------")
	fmt.Println(payload.Advice)
	fmt.Println("-------パワーアップカラー-------")
	fmt.Println(payload.PowerUp)
	fmt.Println("-------クールダウンカラー-------")
	fmt.Println(payload.CoolDown)
}
