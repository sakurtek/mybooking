package render

import (
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/sakurtek/goserver/bookingremyconcept/internal/config"
	"github.com/sakurtek/goserver/bookingremyconcept/internal/model"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// tambakhan fungsi untuk menamgahkan data
// secara default untuk di eksekusi
func AddDefaultData(td *model.TemplateData, r *http.Request) *model.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")

	td.CSRFToken = nosurf.Token(r)
	return td
}
