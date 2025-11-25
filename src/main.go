package main

import (
	"os"
	"io"
	"slices"
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
	adminsSlice []string  //held in memory to prevent creation without restart
	logLevel = log.DebugLevel
)

func main() {
	s := &ssh.Server{
		Addr:            ":"+port,
		Handler:         client,
	};if !noLogin {
		s.PasswordHandler = authClient
	}

	ssh.Handle(client)
	log.Infof("listening on port %s", port)
	log.Fatal(s.ListenAndServe())
}

func client(s ssh.Session) {
 io.WriteString(s, "foo\n")
}

func authClient(ctx ssh.Context, passIn string) bool {
	if noLogin { return true }
	if passIn == "" { return false }

	user := ctx.User()
	pass := getUserPass(user)
	var isValid bool
	if passIn == pass {	isValid = true }

	if slices.Contains(adminsSlice, user) {
		log.Warnf("admin login attempt; valid? %v.", !isValid)
	}

	return isValid
}

func getUserPass(user string) string {
	var passFile, pass string
	var logins gomn.Map
	var err error
	var ok bool

	if slices.Contains(adminsSlice, user) {
		passFile = adminsFile
	} else { passFile = loginFile }

	var loginsBytes []byte
	if loginsBytes, err = os.ReadFile(passFile); err != nil {
		log.Errorf("failed to read logins file:  %v", err)
	} else { log.Debug("read logins file") }

	if logins, err = gomn.Parse(string(loginsBytes)); err != nil {
		log.Errorf("failed to parse logins file:  %v", err)
		return ranPass(int64(16))
	} else {
		log.Debug("parsed logins file")
		passFile = adminsFile
	}

	if pass, ok = logins[user].(string); !ok {
		return ranPass(int64(16))		
	} else { log.Debug("got real password") }

	return pass
}

