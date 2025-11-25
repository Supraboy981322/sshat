package main

import (
	"github.com/gliderlabs/ssh"
	"github.com/charmbracelet/log"
	"github.com/Supraboy981322/gomn"
)

var (
	port string
	noLogin bool
	conf gomn.Map
	admins gomn.Map
	loginFile string
	adminsFile string
	bannerText string
	adminsSlice []string  //held in memory to prevent creation without restart
	logLevel = log.DebugLevel
	sshServer *ssh.Server
	
	colors = map[string]string {
		"black":      "\033[0;30m",  "bBlack":   "\033[1;30m",
		"red":        "\033[0;31m",  "bRed":     "\033[1;31m",
		"green":      "\033[0;32m",  "bGreen":   "\033[1;32m",
		"yellow":     "\033[0;33m",  "bYellow":  "\033[1;33m",
		"blue":       "\033[0;34m",  "bBlue":    "\033[1;34m",
		"purple":     "\033[0;35m",  "bPurple":  "\033[1;35m",
		"cyan":       "\033[0;36m",  "bCyan":    "\033[1;36m",
		"white":      "\033[0;37m",  "bWhite":   "\033[1;37m",

		"uBlack":     "\033[4;30m",  "bgBlack":  "\033[40m",
		"uRed":       "\033[4;31m",  "bgRed":    "\033[41m",
		"uGreen":     "\033[4;32m",  "bgGreen":  "\033[42m",
		"uYellow":    "\033[4;33m",  "bgYellow": "\033[43m",
		"uBlue":      "\033[4;34m",  "bgBlue":   "\033[44m",
		"uPurple":    "\033[4;35m",  "bgPurple": "\033[45m",
		"uCyan":      "\033[4;36m",  "bgCyan":   "\033[46m",
		"uWhite":     "\033[4;37m",  "bgWhite":  "\033[47m",

		"hiBlack":    "\033[0;90m",  "bhiBlack": "\033[1;90m",
		"hiRed":      "\033[0;91m",  "bhiRed":   "\033[1;91m",
		"hiGreen":    "\033[0;92m",  "bhiGreen": "\033[1;92m",
		"hiYellow":   "\033[0;93m",  "bhiYellow":"\033[1;93m",
		"hiBlue":     "\033[0;94m",  "bhiBlue":  "\033[1;94m",
		"hiPurple":   "\033[0;95m",  "bhiPurple":"\033[1;95m",
		"hiCyan":     "\033[0;96m",  "bhiCyan":  "\033[1;96m",
		"hiWhite":    "\033[0;97m",  "bhiWhite": "\033[1;97m",

		"hibgBlack": "\033[0;100m",  "reset": "\033[0m",
		"hibgRed":   "\033[0;101m",  "off":   "\033[0m",
		"hibgGreen": "\033[0;102m",  "none":  "\033[0m",
		"hibgYellow":"\033[0;103m",  
		"hibgBlue":  "\033[0;104m",  
		"hibgPurple":"\033[0;105m",  
		"hibgCyan":  "\033[0;106m",  
		"hibgWhite": "\033[0;107m",  
	}
)

func main() {
	startSSH()
}
