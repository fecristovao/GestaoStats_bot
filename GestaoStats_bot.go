package main

import (
	"log"
	"os"
	"time"

	"github.com/fecristovao/GestaoStats_bot/telegram"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	os.Setenv("TZ", "America/Sao_Paulo")
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	time.Local = loc
	b, err := tb.NewBot(tb.Settings{
		URL:    "http://powerful-ridge-88814.herokuapp.com:" + os.Getenv("PORT"),
		Token:  "5052397467:AAGOjmx5gP7z-iHCssTLD4aARlXPFt9F0_E",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	telegram.GblSts = make(map[string]telegram.Stats)

	b.Handle(tb.OnText, func(m *tb.Message) {
		telegram.CheckBetText(m, b)
		telegram.CheckGreen(m, b)
		telegram.CheckRed(m, b)
	})

	b.Handle(tb.OnPhoto, func(m *tb.Message) {
		telegram.CheckPhoto(m, b)
	})

	b.Handle("/stats", func(m *tb.Message) {
		//spew.Dump(m)
		var key string
		if m.Payload == "" {
			today := time.Now()
			key = today.Format("2006-02-01")
		} else {
			key = m.Payload
		}
		telegram.SendStats(m, b, key)
	})

	b.Start()

}
