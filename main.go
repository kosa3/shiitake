package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

// TODO: 占い結果
type shiitakeResponse struct {
}

const appVersion = "0.1.0"

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

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name: "version", Aliases: []string{"v"},
		Usage: "print only the version",
	}
	app := cli.App{
		Name:  "shiitake",
		Usage: "shiitake-fortune-telling",
		Action: func(c *cli.Context) error {
			fmt.Println("しいたけ占いへようこそ！")
			time.Sleep(time.Second * 1)
			fmt.Println("占いたい星座の英字を入力して")
			time.Sleep(time.Second * 1)
			//TODO 関数として切り出して失敗した時に呼び出す形でも良さそう
			for alias, constellation := range constellations {
				fmt.Println(alias, "(" + constellation + ")")
			}
			fmt.Print("> ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			constellation := scanner.Text()

			// check
			if _, ok := constellations[constellation]; !ok {
				log.Fatal(constellation + "という星座はありません。。")
			}

			fmt.Println(constellations[constellation], "の運勢はこちら")
			// 毎回スクレイピングはあれなのでサーバーでjsonをホスティング配信する
			// jsonを取得しその内容を表示する
			return nil
		},
		Version: appVersion,
		Commands: []*cli.Command{
			{
				Name:    "configure",
				Aliases: []string{"c"},
				Usage:   "setting your profile",
				Action: func(c *cli.Context) error {
					scanner := bufio.NewScanner(os.Stdin)
					fmt.Println("あなたの星座はどれ？")

					scanner.Scan()
					fmt.Println(scanner.Text())
					// 終わったらホーム配下に設定ファイルを作り書き込む
					fmt.Println("あなたの星座は~~ですね")
					fmt.Println("次からshiitake meで自分の占い結果が見れるよ")
					return nil
				},
			},
			{
				Name:  "me",
				Usage: "my shiitake result",
				Action: func(c *cli.Context) error {
					// 設定ファイル
					fmt.Println("今週のあなたの星座の運勢は")

					return nil
				},
			},
		},
		Before: func(c *cli.Context) error {
			fmt.Println("before")
			// ここで設定ファイルをチェックしてなければconfigureして
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
