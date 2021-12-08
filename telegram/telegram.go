package telegram

import (
	"fmt"
	"log"
	"strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

var GblSts map[string]Stats

// CheckPhoto checks if a caption of photo belongs to an bet365 bet
func CheckPhoto(m *tb.Message, b *tb.Bot) bool {
	time := checkTime(m)
	if strings.Contains(m.Caption, BetURL) {
		log.Printf("[%v] - Photo Bet Received: %s\n", time, m.Caption)
		registerBet(m, b)
		return true
	}
	return false
}

// CheckBetText check if a message is a bet
func CheckBetText(m *tb.Message, b *tb.Bot) bool {
	time := checkTime(m)
	if strings.Contains(m.Text, BetURL) {
		log.Printf("[%v] - Bet Text Received: %s\n", time, m.Text)
		registerBet(m, b)
		return true
	}
	return false
}

// CheckGreen check if bet is green
func CheckGreen(m *tb.Message, b *tb.Bot) bool {
	time := checkTime(m)
	if strings.Contains((m.Text), GreenEmoji) {
		log.Printf("[%v] - Green!\n", time)
		registerGreen(m, b)
		return true
	}
	return false
}

// CheckRed check if a bet is red
func CheckRed(m *tb.Message, b *tb.Bot) bool {
	time := checkTime(m)
	if strings.Contains((m.Text), RedEmoji) {
		log.Printf("[%v] - Red!\n", time)
		registerRed(m, b)
		return true
	}
	return false
}

// SendStats send stats to client
func SendStats(m *tb.Message, b *tb.Bot, key string) {
	day := GblSts[key]
	response := key + "\n"
	response += "Bets registradas: %d\n"
	response += "Greens registradas: %d\n"
	response += "Reds registradas: %d\n"
	response = fmt.Sprintf(response, day.Bets, day.Green, day.Red)
	b.Send(m.Chat, response)

	response = "Taxa de Green: %.2f %\n"
	response += "Taxa de Red: %.2f %\n"
	response += "Green/Red: %.2f %\n"

	taxGreen := (100.0 * day.Green) / day.Bets
	taxRed := (100.0 * day.Red) / day.Bets
	taxGR := day.Green / day.Red
	response = fmt.Sprintf(response, float32(taxGreen), float32(taxRed), float32(taxGR))

	b.Send(m.Chat, response)
}
