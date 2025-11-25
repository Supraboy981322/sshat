package main

import (
	"os"
	"strings"
	"strconv"
	"github.com/charmbracelet/log"
	"github.com/Supraboy981322/gomn"
)

func init() {
	log.Info("Using debug log level until parsed config") 
	log.SetLevel(logLevel)
	var err error
	var ok bool
	confBytes, err := os.ReadFile("conf.gomn")
	if err != nil {
		log.Fatalf("problem reading config:  %v", err)
	} else { log.Debug("read config") }

	if conf, err = gomn.Parse(string(confBytes)); err != nil {
		log.Fatalf("err parsing conf:  %v", err)
	} else { log.Debug("parsed config") }

	if lgLvl, ok := conf["log level"].(string); ok {
		switch strings.ToLower(lgLvl) {
		case "debug":
			logLevel = log.DebugLevel
		case "info":
			logLevel = log.DebugLevel
		case "warn":
			logLevel = log.DebugLevel
		case "error":
			logLevel = log.DebugLevel
		case "fatal":
			logLevel = log.DebugLevel
		default:
			log.Fatalf("invalid log level:  %s", lgLvl)
		}
		log.SetLevel(logLevel)
		log.Infof("log level set to %s", lgLvl)
	} else { log.Fatal("err checking log level") }

	if logConf, ok := conf["logins"].(gomn.Map); ok {
		log.Debug("found logins config")

		if adminsFile, ok = logConf["admins file"].(string); ok {
			log.Debug("admins file set")
			var adminsBytes []byte
			if adminsBytes, err = os.ReadFile(adminsFile); err != nil {
				log.Fatalf("err reading admins file:  %v", err)
			} else { log.Debug("read admins file") }

			if admins, err = gomn.Parse(string(adminsBytes)); err == nil {
				log.Debug("parsed admins file")
				for usernameRaw, _ := range admins {
					if username, ok := usernameRaw.(string); ok {
						adminsSlice = append(adminsSlice, username)
					} else { log.Fatal("invalid admin username") }
				}
			} else { log.Fatal("problem parsing admins file") }
		} else { log.Warn("no admins file set") }

		if loginFile, ok = logConf["file"].(string); !ok {
			log.Error("logins file not set")
		} else { log.Debug("logins file set") }

		if noLogin, ok = logConf["no login"].(bool); ok && noLogin {
			log.Warn("logins disabled, therefore no auth will be used")
		} else {
			if !noLogin {
				log.Debug(`["no login"] is set to false (good)`)
			} else {
				log.Debug(`["no login"] not configured; defaulting to false (good)`) 
			}
		}

	} else {
		log.Warn("didn't find login config")
		if noLogin, ok = conf["no login"].(bool); ok {
			log.Warn("logins disabled, therefore no auth will be used")
		} else { log.Fatal("logins not configured") }
	}

	var portInt int
	if portInt, ok = conf["port"].(int); !ok {
		portInt = 2222
		log.Warnf("port not defined, defaulting to %d", portInt)
	};port = strconv.Itoa(portInt)
 
	if bannerText, ok = conf["banner"].(string); !ok {
		bannerText = "sshat"
		log.Warnf("banner text not defined, defaulting to %s", bannerText)
	}
}
