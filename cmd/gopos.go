package main

import (
	"github.com/matejkramny/gopos"
)

func main() {
	// your printer IP & Port
	// In this case, the printer has Ethernet interface and assigned a static ip
	// TM-T20II Ethernet
	printer := gopos.POS("192.168.0.66", 9100)

	gopos.ConnectRedis()
	defer (*gopos.Redis).Close()

	gopos.RedisListen(printer)

	// testing
	// pos.Print("I am NOT emphesized")
	// pos.LineFeed()

	// pos.SetEmphesized(gopos.MODE_ON)
	// pos.Print("I am emphesized")
	// pos.SetEmphesized(gopos.MODE_OFF)
	// pos.LineFeed()

	// pos.SetFont(gopos.MODE_ON)
	// pos.Print("I am in Font B")
	// pos.LineFeed()
	// pos.SetFont(gopos.MODE_OFF)

	// pos.SetFont(gopos.MODE_ON2)
	// pos.Print("I am in Font C")
	// pos.LineFeed()
	// pos.SetFont(gopos.MODE_OFF)

	// pos.SetReversePrint(gopos.MODE_ON)
	// pos.Print("I am reverse white/black thing")
	// pos.LineFeed()
	// pos.SetReversePrint(gopos.MODE_OFF)

	// pos.SetUnderline(gopos.MODE_ON)
	// pos.Print("I am in Underline mode")
	// pos.LineFeed()
	// pos.SetUnderline(gopos.MODE_OFF)

	// pos.SetUnderline(gopos.MODE_ON2)
	// pos.Print("I am in 2 Underline mode")
	// pos.LineFeed()
	// pos.SetUnderline(gopos.MODE_OFF)

	// pos.SetDoubleStrike(gopos.MODE_ON)
	// pos.Print("I am in Double strike")
	// pos.LineFeed()
	// pos.SetDoubleStrike(gopos.MODE_OFF)

	// pos.LineFeed()
	// pos.LineFeed()
	// pos.LineFeed()
	// pos.LineFeed()

	// pos.Cut()
}
