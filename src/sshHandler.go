package main

import (
	"os"
//	"fmt"
	"slices"
	"github.com/gliderlabs/ssh"
	"github.com/charmbracelet/log"
	"github.com/Supraboy981322/gomn"
)

func startSSH() {
	sshServer = &ssh.Server{
		Addr:     ":"+port,
		Handler:  client,
		Banner:   "sshat",
	};if !noLogin {
		sshServer.PasswordHandler = authClient
	}

	ssh.Handle(client)
	log.Infof("listening on port %s", port)
	log.Fatal(sshServer.ListenAndServe())
}

func authClient(ctx ssh.Context, passIn string) bool {
	if noLogin { return true }
	if passIn == "" { return false }
	var isAdmin bool

	user := ctx.User()
	pass := getUserPass(user)
	var isValid bool
	if passIn == pass {	isValid = true }

	if slices.Contains(adminsSlice, user) {
		if isValid { isAdmin = true }
		log.Warnf("admin login attempt; valid? %v.", isValid)
	}
	
	if !isValid {
		log.Debugf("invalid login:  %v", ctx.RemoteAddr())
	} else { ctx.SetValue("admin", isAdmin) }

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
		log.Debug("failed to get real password, returning random string")
		return ranPass(int64(16))
	} else {
		log.Debug("parsed logins file")
		passFile = adminsFile
	}

	if pass, ok = logins[user].(string); !ok {
		log.Debug("failed to get real password, returning random string")
		return ranPass(int64(16))		
	} else { log.Debug("got real password") }

	return pass
}

//Finally, the client handler
func client(s ssh.Session) {
	ctx := s.Context() 
	user := s.User()
	isAdmin, _ := ctx.Value("admin").(bool)
	cls(s)
	banner(s)
	if isAdmin {
		sendLn(s, "welcome, "+colors["iYellow"]+"admin"+colors["reset"])
	}	else {
		sendLn(s, "welcome, "+colors["iBlue"]+user+colors["reset"])
	}
}
