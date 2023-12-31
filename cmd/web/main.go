package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sakurtek/goserver/bookingremyconcept/pkg/config"
	"github.com/sakurtek/goserver/bookingremyconcept/pkg/handlers"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var sessionmanager *scs.SessionManager

func main() {

	// UBAH INI MENJADI -TRUE- APABILA PRODUKSI ATAU SUDAH MAU DI DEPLOY KE SERVER
	app.InProduction = false

	sessionmanager = scs.New()
	sessionmanager.Lifetime = 24 * time.Hour
	sessionmanager.Cookie.Persist = true
	sessionmanager.Cookie.SameSite = http.SameSiteLaxMode
	sessionmanager.Cookie.Secure = app.InProduction

	/* SIMPAN sessionmanager KE config.AppConfig */
	app.Session = sessionmanager

	/* pelajari ini dengan baik karena dengan dua fungsi ini
	Penggunaan SESSION berhasil
	*/
	// Penting ini----
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)
	// Penting ini----

	fmt.Println("Server running in localhost port: ", portNumber)

	myserver := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err := myserver.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
