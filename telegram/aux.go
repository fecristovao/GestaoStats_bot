package telegram

import (
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

const format = "2006-02-01"

func checkTime(m *tb.Message) time.Time {
	var timeResult int64

	if m.OriginalUnixtime != 0 {
		timeResult = int64(m.OriginalUnixtime)
	} else {
		timeResult = m.Unixtime
	}

	return time.Unix(timeResult, 0)
}

func registerBet(m *tb.Message, b *tb.Bot) {
	time := checkTime(m)
	day := time.Format(format)
	stat := GblSts[day]
	stat.Bets++
	GblSts[day] = stat
	b.Send(m.Chat, "Bet registrada")
}

func registerGreen(m *tb.Message, b *tb.Bot) {
	time := checkTime(m)
	day := time.Format(format)
	stat := GblSts[day]
	stat.Green++
	GblSts[day] = stat
	b.Send(m.Chat, "Green Registrado")
}

func registerRed(m *tb.Message, b *tb.Bot) {
	time := checkTime(m)
	day := time.Format(format)
	stat := GblSts[day]
	stat.Red++
	GblSts[day] = stat
	b.Send(m.Chat, "Red Registrado")
}
