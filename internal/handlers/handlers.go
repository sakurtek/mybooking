package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/justinas/nosurf"
	"github.com/sakurtek/goserver/bookingremyconcept/internal/config"
	"github.com/sakurtek/goserver/bookingremyconcept/internal/model"
	"github.com/sakurtek/goserver/bookingremyconcept/internal/modelproc"
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

// tambahka fungsi untuk view template
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

// tambakhan fungsi untuk menamgahkan data
// secara default untuk di eksekusi
func AddLikeDefaultData(td *model.TemplateData, r *http.Request) *model.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func (m *Repository) HandleHome(w http.ResponseWriter, r *http.Request) {
	//get session
	mytemplate, _ = ViewTemplate("templates/home.page.gohtml")

	//td := AddLikeDefaultData(&model.TemplateData{}, r)
	err := mytemplate.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) HandleAbout(w http.ResponseWriter, r *http.Request) {

	mytemplate, _ = ViewTemplate("templates/about.page.gohtml")

	//td := AddLikeDefaultData(&model.TemplateData{}, r)
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
		//td := AddLikeDefaultData(&model.TemplateData{}, r)
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

	//td := AddLikeDefaultData(&model.TemplateData{}, r)
	err := mytemplate.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) HandleGenerals(w http.ResponseWriter, r *http.Request) {

	mytemplate, _ = ViewTemplate("templates/generals.page.gohtml")

	//td := AddLikeDefaultData(&model.TemplateData{}, r)
	err := mytemplate.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) HandleMajors(w http.ResponseWriter, r *http.Request) {

	mytemplate, _ = ViewTemplate("templates/majors.page.gohtml")

	//td := AddLikeDefaultData(&model.TemplateData{}, r)
	err := mytemplate.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) HandleMakeReservation(w http.ResponseWriter, r *http.Request) {

	mytemplate, _ = ViewTemplate("templates/make-reservation.page.gohtml")

	td := AddLikeDefaultData(&model.TemplateData{}, r)
	err := mytemplate.Execute(w, td)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) HandleSearchAvailability(w http.ResponseWriter, r *http.Request) {

	mytemplate, _ = ViewTemplate("templates/search-availability.page.gohtml")

	td := AddLikeDefaultData(&model.TemplateData{}, r)
	err := mytemplate.Execute(w, td)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) HandlePostSearchAvailability(w http.ResponseWriter, r *http.Request) {

	// cara mengambil form pada GO
	mStart := r.Form.Get("start")
	mEnd := r.Form.Get("end")

	type DataCheckAvailability struct {
		StartDate  string
		EndDate    string
		StatusInfo bool
	}

	dateChceck := DataCheckAvailability{
		StartDate:  mStart,
		EndDate:    mEnd,
		StatusInfo: true,
	}

	log.Println(mStart, mEnd)
	//w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", mStart, mEnd)))

	// ambil data halaman proses untuk tes menampilkan data
	mytemplate, _ := ViewTemplate("templates/proses.check.gohtml")

	err := mytemplate.Execute(w, dateChceck)
	if err != nil {
		fmt.Println(err)
	}
}
