package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/sakurtek/goserver/bookingremyconcept/pkg/config"
	"github.com/sakurtek/goserver/bookingremyconcept/pkg/model"
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
		"templates/header.layout.gohtml",
		"templates/footer.layout.gohtml",
	)

	return mytemplate, err
}

func (m *Repository) HandleHome(w http.ResponseWriter, r *http.Request) {
	//get session
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	m.App.Session.Put(r.Context(), "userID", "sakur.stendy@gmail.com")

	MyIndexData := model.DataStruk{
		MyID:       model.SessionProperty{}, // kosoong
		MyKaryawan: modelproc.GetDataKaryawan(),
		MyNews:     modelproc.GetOnlineNews(),
	}

	mytemplate, _ = ViewTemplate("templates/home.page.gohtml")

	err := mytemplate.Execute(w, MyIndexData)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) HandleAbout(w http.ResponseWriter, r *http.Request) {
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	userID := m.App.Session.GetString(r.Context(), "userID")

	myStringMap := model.SessionProperty{
		LabelName: "Hello, again.",
		Status:    "Data untuk About page",
		RemoteIP:  remoteIP,
		UserID:    userID,
	}

	mytemplate, _ = ViewTemplate("templates/about.page.gohtml")

	err := mytemplate.Execute(w, myStringMap)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) HandleNewsDetail(w http.ResponseWriter, r *http.Request) {

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	userID := m.App.Session.GetString(r.Context(), "userID")

	myStringMap := model.SessionProperty{
		LabelName: "Hello, again.",
		Status:    "Data untuk About page",
		RemoteIP:  remoteIP,
		UserID:    userID,
	}

	var IDExist bool // buat status data yang dicari

	// get id dari url paramter
	q := r.URL.Query()
	idstr := q.Get("id")
	idint, _ := strconv.Atoi(idstr) // conversi string ke int

	mytemplate, _ = ViewTemplate("templates/detail.page.gohtml")

	// membuat variabel dari tipe data struk NewsDetail
	// untuk memasukkan ID
	mynewdetail := model.NewsDetail{
		IDParam: idint,
		MyNews:  modelproc.GetOnlineNews(),
		MyID:    myStringMap,
	}

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
		err := mytemplate.Execute(w, mynewdetail)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("Redirect")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
