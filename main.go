package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

// yaml用
type constellationSetting struct {
	constellation string
}

// TODO: 占い結果
type shiitakeResponse struct {
}

const appVersion = "0.1.0"

const configFile = ".shiitake.yml"

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
			// TODO: 毎回スクレイピングはあれなのでサーバーでjsonをホスティング配信する
			// TODO: jsonを取得しその内容を表示する
			return nil
		},
		Version: appVersion,
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

					fp, err := os.Create(configFile)
					if err != nil {
						log.Fatal(err)
					}
					defer fp.Close()

					fp.WriteString("constellation: " + constellation)

					fmt.Println("あなたの星座は" + constellation + "ですね")
					fmt.Println("次からshiitake meで自分の占い結果が見れるよ")
					return nil
				},
			},
			{
				Name:  "me",
				Usage: "my shiitake result",
				Action: func(c *cli.Context) error {
					// TODO: 設定ファイル読み込み
					fmt.Println("今週のあなたの星座の運勢は")

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
