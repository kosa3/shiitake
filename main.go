package main

import (
	"bufio"
	"fmt"
	"github.com/Songmu/flextime"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type ShiitakeSetting struct {
	Constellation string `yaml:"constellation"`
}

type Option struct {
	Ago int
}

const (
	AppVersion = "0.1.0"
	ConfigFile = "/tmp/.shiitake.yml"
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
func scanConstellation(ShiitakeSetting ShiitakeSetting) (string, error) {
	for alias, constellation := range constellations {
		fmt.Println(alias, "("+constellation+")")
	}

	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		ShiitakeSetting.Constellation = scanner.Text()
		// check validate
		if len(ShiitakeSetting.Constellation) == 0 {
			fmt.Println("星座を入力してください")
			continue
		}

		if _, ok := constellations[ShiitakeSetting.Constellation]; !ok {
			fmt.Println(ShiitakeSetting.Constellation + "という星座はありません。。")
			continue
		}
		break
	}

	return ShiitakeSetting.Constellation, nil
}

func getThisMonday(option Option) time.Time {
	date := flextime.Now().AddDate(0, 0, -7*option.Ago)
	weekday := date.Weekday()
	// 日曜日は0だから合わせるために7にする
	if weekday == 0 {
		weekday = 7
	}
	// 月曜にしたものを返却する
	return date.Add(time.Duration(-24*(weekday-1)) * time.Hour)
}

func formatPeriod(date time.Time) string {
	return "【" + date.Format("2006/1/2") + "~" + date.AddDate(0, 0, 6).Format("2006/1/2") + "】"
}

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name: "version", Aliases: []string{"v"},
		Usage: "print the version",
	}
	app := cli.App{
		Name:  "shiitake",
		Usage: "shiitake-fortune-telling",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "ago",
				Value: 0,
				Usage: "before fortune-telling",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("しいたけ占いへようこそ！")
			time.Sleep(time.Second * 1)
			fmt.Println("占いたい星座の英字を入力して")
			time.Sleep(time.Second * 1)

			var ShiitakeSetting ShiitakeSetting
			constellation, err := scanConstellation(ShiitakeSetting)
			if err != nil {
				log.Fatal("もう一度最初からやり直してください")
			}

			date := getThisMonday(Option{Ago: c.Int("ago")})
			fmt.Println(formatPeriod(date) + constellations[constellation] + "の運勢はこちら")
			ShiitakeResponse, err := fetchFortuneTelling(date)
			if err != nil {
				log.Fatal(err)
			}

			if err := ShiitakeResponse.showFortuneTellingByConstellation(constellation); err != nil {
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
					var ShiitakeSetting ShiitakeSetting
					constellation, err := scanConstellation(ShiitakeSetting)
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

					date := getThisMonday(Option{Ago: c.Int("ago")})
					fmt.Println(formatPeriod(date) + constellations[setting.Constellation] + "の運勢はこちら")

					ShiitakeResponse, err := fetchFortuneTelling(date)
					if err != nil {
						log.Fatal(err)
					}

					if err := ShiitakeResponse.showFortuneTellingByConstellation(setting.Constellation); err != nil {
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
