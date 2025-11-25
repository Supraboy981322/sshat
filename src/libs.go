package main

import (
	"io"
//	"fmt"
//	"strconv"
	"math/big"
	"crypto/rand"
	randFallback "math/rand"
	"github.com/gliderlabs/ssh"
	"github.com/charmbracelet/log"
)

func ranPass(l int64) string {
	var res string
	var chars = []string{
		"a", "b",	"c", "d", "e", "f", "g",
		"h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u",
		"v", "w", "x", "y", "z", "A", "B",
		"C", "D", "E", "F", "G", "H", "I",
		"J", "K", "L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U", "V", "W",
		"X", "Y", "Z", "0", "9", "8", "7",
		"6", "5", "4", "3", "2", "1", "!",
		"@", "#", "$", "%", "^", "&", "*",
		"(", ")", "-", "_", "=", "+", "[",
		"]", "{", "}", "|", "\\", ";", ":",
		"'", "\"", "<", ">", "/", "?", ".",
		",",
	}

	var i int64
	for i = 0; i < l; i++ {
		ranDig := randInt64(len(chars))
		res += chars[ranDig]
	}

	return res
}

func randInt64(m int) int {
	bigInt := big.NewInt(int64(m))
	i, err := rand.Int(rand.Reader, bigInt)
	if err != nil {
		log.Errorf("err in crypto/rand:  %v", err)
		log.Warn("using fallback random number generator")
		i = big.NewInt(int64(randFallback.Intn(m)))
	}

	return int(i.Int64())
}


func send(s ssh.Session, str string) {
	io.WriteString(s, str+"\n")
}
func sendLn(s ssh.Session, str string) {
	send(s, str+"\n")
}
func cls(s ssh.Session) {
	send(s, "\033[0m\033[2J\033[H")
}
func toStart(s ssh.Session) {
	send(s, "\033[2J")
}
func cll(s ssh.Session) {
	send(s, "\033[0K")
}
func banner(s ssh.Session) {
//	ctx := s.Context()
	pty, _, _ := s.Pty()
	window := pty.Window
	width := window.Width

	//define banner
	type bStruct struct {
		Text     string
		Res      string
		txtColor string
		bgColor  string
		Padding  int
	};b := bStruct {
		Text:    "sshat",
		Res:     colors["bgBlue"],
		txtColor:"\033[30;44m",
		bgColor: colors["bgBlue"],
		Padding: (width/2)-(len(bannerText)/2),
	}
	 
	//construct banner
	for i := 0; i < b.Padding; i++ {
		b.Res += " "
	};b.Res += colors["off"]+b.txtColor+b.Text+b.bgColor
	for i := 0; i < b.Padding; i++ {
		b.Res += " "
	};b.Res += colors["off"]

	cls(s)
	send(s, b.Res)
}
