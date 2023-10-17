package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/sakurtek/goserver/bookingremyconcept/pkg/config"
	"github.com/sakurtek/goserver/bookingremyconcept/pkg/modelproc"
)

var Repo *Repository
var mytemplate *template.Template

// var session manager *scs.SessionManger
// untuk handling saja
// untuk type structur lainnya sebaiknya dimasukkan ke dalam
// modul model
type Repository struct {
	App *config.AppConfig
}

// buat new repo
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func ViewTemplate(maintemplate string) (*template.Template, error) {
	var mytemplate *template.Template

	mytemplate, err := template.ParseFiles(
		maintemplate,
		"templates/base.layout.gohtml",
		//"templates/header.layout.gohtml",
		//"templates/footer.layout.gohtml",
	)

	return mytemplate, err
}

func (m *Repository) HandleHome(w http.ResponseWriter, r *http.Request) {
	//get session
	mytemplate, _ = ViewTemplate("templates/home.page.gohtml")

	err := mytemplate.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) HandleAbout(w http.ResponseWriter, r *http.Request) {

	mytemplate, _ = ViewTemplate("templates/about.page.gohtml")

	err := mytemplate.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) HandleNewsDetail(w http.ResponseWriter, r *http.Request) {

	var IDExist bool // buat status data yang dicari

	// get id dari url paramter
	q := r.URL.Query()
	idstr := q.Get("id")
	idint, _ := strconv.Atoi(idstr) // conversi string ke int

	mytemplate, _ = ViewTemplate("templates/detail.page.gohtml")

	/*
		INI CONTOH MENCARI ID TAPI INI MASIH MENGGUNAKAN DATA
		STATIC TIDAK MENGGUAKAN DATABASE
		JIKA MENGGUNAKAN DATABASE MAKA CUKUP MENGGUNAKAN SQL SAJA
	*/

	if idstr != "" {
		totalExist := 0
		for _, IDParam := range modelproc.GetOnlineNews() {
			if IDParam.ID == idint {
				totalExist++
			}
		}

		if totalExist > 0 {
			IDExist = true
		} else {
			IDExist = false
		}
	}

	if IDExist {
		err := mytemplate.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("Redirect")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (m *Repository) HandleContact(w http.ResponseWriter, r *http.Request) {

	mytemplate, _ = ViewTemplate("templates/contact.page.gohtml")

	err := mytemplate.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) HandleGenerals(w http.ResponseWriter, r *http.Request) {

	mytemplate, _ = ViewTemplate("templates/generals.page.gohtml")

	err := mytemplate.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) HandleMajors(w http.ResponseWriter, r *http.Request) {

	mytemplate, _ = ViewTemplate("templates/majors.page.gohtml")

	err := mytemplate.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) HandleMakeReservation(w http.ResponseWriter, r *http.Request) {

	mytemplate, _ = ViewTemplate("templates/make-reservation.page.gohtml")

	err := mytemplate.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) HandleSearchAvailability(w http.ResponseWriter, r *http.Request) {

	mytemplate, _ = ViewTemplate("templates/search-availability.page.gohtml")

	err := mytemplate.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
