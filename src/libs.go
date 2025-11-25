package main

import (
	"math/big"
	"crypto/rand"
	randFallback "math/rand"
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
