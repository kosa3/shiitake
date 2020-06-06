package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type ShiitakeSetting struct {
	Constellation string `yaml:"constellation"`
}

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

const (
	AppVersion = "0.1.0"
	ConfigFile = ".shiitake.yml"
	BaseUrl    = "https://shiitake-fortune-telling.s3-ap-northeast-1.amazonaws.com/"
)

var constellations = map[string]string{
	"aries":       "おひつじ座",
	"taurus":      "おうし座",
	"gemini":      "ふたご座",
	"cancer":      "かに座",
	"leo":         "しし座",
	"virgo":       "おとめ座",
	"libra":       "てんびん座",
	"scorpio":     "さそり座",
	"sagittarius": "いて座",
	"capricorn":   "やぎ座",
	"aquarius":    "みずがめ座",
	"pisces":      "うお座",
}

// 星座を選択する
func scanConstellation() (string, error) {
	for alias, constellation := range constellations {
		fmt.Println(alias, "("+constellation+")")
	}

	var constellation string
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		constellation = scanner.Text()
		// check validate
		if len(constellation) == 0 {
			fmt.Println("星座を入力してください")
			continue
		}

		if _, ok := constellations[constellation]; !ok {
			fmt.Println(constellation + "という星座はありません。。")
			continue
		}
		break
	}

	return constellation, nil
}

func fetchFortuneTelling() (ShiitakeResponse, error) {
	// TODO: 月曜日を取得
	ShiitakeResponse := ShiitakeResponse{}
	url := BaseUrl + "20200601.json"
	response, _ := http.Get(url)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(body, &ShiitakeResponse); err != nil {
		log.Fatal(err)
	}

	return ShiitakeResponse, nil
}

func showFortuneTellingByConstellation(constellation string, ShiitakeResponse ShiitakeResponse) error {
	switch constellation {
	case "aries":
		showMessage(ShiitakeResponse.Aries.Payload)
	case "taurus":
		showMessage(ShiitakeResponse.Taurus.Payload)
	case "gemini":
		showMessage(ShiitakeResponse.Gemini.Payload)
	case "cancer":
		showMessage(ShiitakeResponse.Cancer.Payload)
	case "leo":
		showMessage(ShiitakeResponse.Leo.Payload)
	case "virgo":
		showMessage(ShiitakeResponse.Virgo.Payload)
	case "libra":
		showMessage(ShiitakeResponse.Libra.Payload)
	case "scorpio":
		showMessage(ShiitakeResponse.Scorpio.Payload)
	case "sagittarius":
		showMessage(ShiitakeResponse.Sagittarius.Payload)
	case "capricorn":
		showMessage(ShiitakeResponse.Capricorn.Payload)
	case "aquarius":
		showMessage(ShiitakeResponse.Aquarius.Payload)
	case "pisces":
		showMessage(ShiitakeResponse.Pisces.Payload)
	default:
		fmt.Println("何かおかしい")
	}
	return nil
}

func showMessage(payload Payload) {
	fmt.Println("-------分析結果-------")
	fmt.Println(payload.Analysis)
	fmt.Println("-------アドバイス-------")
	fmt.Println(payload.Advice)
	fmt.Println("-------パワーアップカラー-------")
	fmt.Println(payload.PowerUp)
	fmt.Println("-------クールダウンカラー-------")
	fmt.Println(payload.CoolDown)
}

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name: "version", Aliases: []string{"v"},
		Usage: "print the version",
	}
	app := cli.App{
		Name:  "shiitake",
		Usage: "shiitake-fortune-telling",
		Action: func(c *cli.Context) error {
			fmt.Println("しいたけ占いへようこそ！")
			time.Sleep(time.Second * 1)
			fmt.Println("占いたい星座の英字を入力して")
			time.Sleep(time.Second * 1)

			constellation, err := scanConstellation()
			if err != nil {
				log.Fatal("もう一度最初からやり直してください")
			}

			fmt.Println(constellations[constellation], "の運勢はこちら")

			ShiitakeResponse, err := fetchFortuneTelling()
			if err != nil {
				log.Fatal(err)
			}

			if err := showFortuneTellingByConstellation(constellation, ShiitakeResponse); err != nil {
				log.Fatal(err)
			}

			return nil
		},
		Version: AppVersion,
		Commands: []*cli.Command{
			{
				Name:    "configure",
				Aliases: []string{"c"},
				Usage:   "setting your profile",
				Action: func(c *cli.Context) error {
					constellation, err := scanConstellation()
					if err != nil {
						log.Fatal("もう一度最初からやり直してください")
					}

					fp, err := os.Create(ConfigFile)
					if err != nil {
						log.Fatal(err)
					}
					defer fp.Close()

					_, err = fp.WriteString("constellation: " + constellation)
					if err != nil {
						log.Fatal(err)
					}

					fmt.Println("あなたの星座は" + constellations[constellation] + "ですね")
					fmt.Println("次からshiitake meで自分の占い結果が見れるよ")
					return nil
				},
			},
			{
				Name:  "me",
				Usage: "my shiitake result",
				Action: func(c *cli.Context) error {
					file, err := ioutil.ReadFile(ConfigFile)
					if err != nil {
						log.Fatal("shiitake configure を実行してください。")
					}

					setting := ShiitakeSetting{}
					err = yaml.Unmarshal([]byte(file), &setting)
					if err != nil {
						log.Fatal(err)
					}

					fmt.Println("今週の" + constellations[setting.Constellation] + "の運勢は")

					ShiitakeResponse, err := fetchFortuneTelling()
					if err != nil {
						log.Fatal(err)
					}

					if err := showFortuneTellingByConstellation(setting.Constellation, ShiitakeResponse); err != nil {
						log.Fatal(err)
					}
					return nil
				},
			},
		},
		Before: func(c *cli.Context) error {
			fmt.Println("--------------------------------")
			return nil
		},
		After: func(c *cli.Context) error {
			fmt.Println("--------------------------------")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
