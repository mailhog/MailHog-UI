package main

import (
	"flag"
	"os"

	"github.com/ian-kent/go-log/log"
	gotcha "github.com/ian-kent/gotcha/app"
	"github.com/mailhog/MailHog-Server/config"
	"github.com/mailhog/MailHog-UI/assets"
	"github.com/mailhog/MailHog-UI/web"
	"github.com/mailhog/http"
)

var conf *config.Config
var exitCh chan int
var host string

func configure() {
	config.RegisterFlags()
	flag.StringVar(&host, "api-host", "", "API host for MailHog UI to connect to")
	flag.Parse()
	conf = config.Configure()
	// FIXME hacky
	web.Host = host
}

func main() {
	configure()

	// FIXME need to make API URL configurable

	exitCh = make(chan int)
	cb := func(app *gotcha.App) {
		web.CreateWeb(conf, app)
	}
	go http.Listen(conf, assets.Asset, exitCh, cb)

	for {
		select {
		case <-exitCh:
			log.Printf("Received exit signal")
			os.Exit(0)
		}
	}
}
