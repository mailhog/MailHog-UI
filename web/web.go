package web

import (
	"html/template"

	gotcha "github.com/ian-kent/gotcha/app"
	"github.com/ian-kent/gotcha/events"
	"github.com/ian-kent/gotcha/http"
	"github.com/mailhog/MailHog-UI/config"
)

var APIHost string

type Web struct {
	config *config.Config
	app    *gotcha.App
}

func CreateWeb(cfg *config.Config, app *gotcha.App) *Web {
	app.On(events.BeforeHandler, func(session *http.Session, next func()) {
		session.Stash["config"] = cfg
		next()
	})

	r := app.Router

	r.Get("/images/(?P<file>.*)", r.Static("assets/images/{{file}}"))
	r.Get("/js/(?P<file>.*)", r.Static("assets/js/{{file}}"))
	r.Get("/", Index)

	app.Config.LeftDelim = "[:"
	app.Config.RightDelim = ":]"

	return &Web{
		config: cfg,
		app:    app,
	}
}

func Index(session *http.Session) {
	session.Stash["APIHost"] = APIHost
	html, _ := session.RenderTemplate("index.html")

	session.Stash["Page"] = "Browse"
	session.Stash["Content"] = template.HTML(html)
	session.Render("layout.html")
}
