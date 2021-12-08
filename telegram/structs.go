package telegram

import "time"

type bet struct {
	Time time.Time
	Link string
}

type Stats struct {
	Bets  int
	Green int
	Red   int
}

// GreenEmoji is a hexadecimal bytes in green emoji
const GreenEmoji = "\xE2\x9C\x85"

// RedEmoji is a hexadecimal bytes to red cross
const RedEmoji = "\xE2\x9D\x8C"

// BetURL contains part of bet365 url
const BetURL = "www.bet365.com/"
