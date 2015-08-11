package gopos

import (
	"fmt"
	"net"
	"strconv"
)

// pretty good documentation here:
// http://content.epson.de/fileadmin/content/files/RSD/downloads/escpos.pdf

const (
	LF = "\x0a"
	ESC = "\x1b"
	CR = "\x0d"
	GS = "\x1d"

	MODE_FONT_A = 0
	MODE_FONT_B = 1
	MODE_EMPHESIZED = 8
	MODE_DOUBLE_HEIGHT = 16
	MODE_DOUBLE_WIDTH = 32
	MODE_UNDERLINE = 128

	MODE_OFF = 0
	MODE_ON = 1
	MODE_ON2 = 2

	JUSTIFICATION_LEFT = 0
	JUSTIFICATION_CENTER = 1
	JUSTIFICATION_RIGHT = 2
)

type ESCPOS struct {
	IP string
	Port int

	Connection net.Conn
}

func (pos *ESCPOS) connect() {
	c, err := net.Dial("tcp", pos.IP + ":" + strconv.Itoa(pos.Port))
	if err != nil {
		panic(err)
	}

	pos.Connection = c
	// Init the printer
	pos.Raw(ESC + "@")
}

func (pos *ESCPOS) Raw(cmd string) {
	fmt.Fprint(pos.Connection, cmd)
}

func (pos *ESCPOS) Print(text string) {
	pos.Raw(text)
}

func POS(ip string, port int) *ESCPOS {
	pos := ESCPOS{IP: ip, Port: port}
	pos.connect()

	return &pos
}

// line feed
func LineFeed() string {
	return LF
}

// n = lines
func Feed(n int) string {
	if n == 1 || n == 0 {
		return LineFeed()
	}

	return ESC + "d" + strconv.Itoa(n)
}

// Reverse print buffer & feed n lines
func ReverseFeed(n int) string {
	return ESC + "e" + strconv.Itoa(n)
}

func Cut() string {
	return GS + "V0"
}

// Wrong possibly
func SetMode(mode int) string {
	return ESC + "!" + strconv.Itoa(mode)
}

func Underline(on bool) string {
	cmd := ESC + "-"
	if on {
		return cmd + "1"
	}

	return cmd + "0"
}

func Emphesize(on bool) string {
	cmd := ESC + "E"
	if on {
		return cmd + "1"
	}

	return cmd + "0"
}

func DoubleStrike(on bool) string {
	cmd := ESC + "G"
	if on {
		return cmd + "1"
	}

	return cmd + "0"
}

// 0-2
func Font(font int) string {
	return ESC + "M" + strconv.Itoa(font)
}

// 0 = left, 1 = center, 2 = right
func Justify(mode int) string {
	return ESC + "a" + strconv.Itoa(mode)
}

// 0 = paper roll near-end sensor disabled
// 1 = paper roll near-end sensor enabled
// 2 = paper roll near-end sensor enabled
// 4 = paper roll end sensor enabled
// 8 = paper roll end sensor enabled
func PaperSensor(mode int) string {
	return ESC + "c3" + strconv.Itoa(mode)
}

// pin = 0, 1
func GeneratePulse(pin int) string {
	return ESC + "p" + strconv.Itoa(pin) + "\x32\x32"
}

// Reverse prints
// 0 = off
// 1 = reverse white/black on
func ReversePrint(on bool) string {
	cmd := GS + "B"
	if on {
		return cmd + "1"
	}

	return cmd + "0"
}
