package main

import (
	"MrMohebi/tracking-post-telegram-bot/requests"
	"github.com/anaskhan96/soup"
	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

func main() {
	println("aaaaa")

	err := godotenv.Load()

	pref := tele.Settings{
		Token:  os.Getenv("TEL_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(tele.OnText, func(c tele.Context) error {
		var (
		//user = c.Sender()
		//text = c.Text()
		)
		err := c.Send("wait a second...")
		if err != nil {
			return err
		}
		bodyString, err := requests.PostSendRequest("104530134300045270076178")
		if err != nil {

		}
		doc := soup.HTMLParse(string(bodyString))
		links := doc.Find("div", "id", "txtVoteR").FindAll("div")
		return c.Send(links)
	})

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Start()
}
