package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/sakurtek/goserver/bookingremyconcept/internal/config"
	"github.com/sakurtek/goserver/bookingremyconcept/internal/forms"
	"github.com/sakurtek/goserver/bookingremyconcept/internal/model"
	"github.com/sakurtek/goserver/bookingremyconcept/internal/modelproc"
	"github.com/sakurtek/goserver/bookingremyconcept/internal/render"
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

// buat new handler
func NewHandler(r *Repository) {
	Repo = r
}

// tambahka fungsi untuk view template
func ViewTemplate(maintemplate string) (*template.Template, error) {
	var mytemplate *template.Template

	mytemplate, err := template.ParseFiles(
		maintemplate,
		"templates/base.layout.sakur",
		//"templates/header.layout.gohtml",
		//"templates/footer.layout.gohtml",
	)

	return mytemplate, err
}

func (m *Repository) HandleHome(w http.ResponseWriter, r *http.Request) {
	//get session
	mytemplate, _ = ViewTemplate("templates/home.page.sakur")

	var emptydatacheck model.DataCheckAvailability

	data := make(map[string]interface{})
	data["dataCheck"] = emptydatacheck

	td := render.AddDefaultData(&model.TemplateData{
		Data: data,
	}, r)

	fmt.Println(data)

	err := mytemplate.Execute(w, td) // awalnya sya tidak pake, dan pake nil
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) HandleAbout(w http.ResponseWriter, r *http.Request) {

	mytemplate, _ = ViewTemplate("templates/about.page.sakur")

	td := render.AddDefaultData(&model.TemplateData{}, r)
	err := mytemplate.Execute(w, td) // awlanya saya pake nil
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

	mytemplate, _ = ViewTemplate("templates/detail.page.sakur")

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
		td := render.AddDefaultData(&model.TemplateData{}, r)
		err := mytemplate.Execute(w, td)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("Redirect")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (m *Repository) HandleContact(w http.ResponseWriter, r *http.Request) {

	mytemplate, _ = ViewTemplate("templates/contact.page.sakur")

	td := render.AddDefaultData(&model.TemplateData{}, r)
	err := mytemplate.Execute(w, td)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) HandleGenerals(w http.ResponseWriter, r *http.Request) {

	mytemplate, _ = ViewTemplate("templates/generals.page.sakur")

	td := render.AddDefaultData(&model.TemplateData{}, r)
	err := mytemplate.Execute(w, td)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) HandleMajors(w http.ResponseWriter, r *http.Request) {

	mytemplate, _ = ViewTemplate("templates/majors.page.sakur")

	td := render.AddDefaultData(&model.TemplateData{}, r)
	err := mytemplate.Execute(w, td)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) HandleMakeReservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation model.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	mytemplate, _ = ViewTemplate("templates/make-reservation.page.sakur")

	td := render.AddDefaultData(&model.TemplateData{
		Form: forms.New(nil),
		Data: data,
	}, r)

	err := mytemplate.Execute(w, td)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) HandlePostMakeReservation(w http.ResponseWriter, r *http.Request) {
	errparse := r.ParseForm()

	if errparse != nil {
		log.Println(errparse)
		return
	}

	reservation := model.Reservation{
		FirstName:   r.Form.Get("firstname"),
		LastName:    r.Form.Get("lastname"),
		Email:       r.Form.Get("email"),
		PhoneNumber: r.Form.Get("phonenumber"),
	}

	form := forms.New(r.PostForm)

	form.Required("firstname", "lastname", "email")
	form.MinLength("firstname", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		mytemplate, _ = ViewTemplate("templates/make-reservation.page.sakur")

		td := render.AddDefaultData(&model.TemplateData{
			Form: form,
			Data: data,
		}, r)

		err := mytemplate.Execute(w, td)
		if err != nil {
			log.Println(err)
		}

		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (m *Repository) HandleSearchAvailability(w http.ResponseWriter, r *http.Request) {

	mytemplate, _ = ViewTemplate("templates/search-availability.page.sakur")

	td := render.AddDefaultData(&model.TemplateData{}, r)
	err := mytemplate.Execute(w, td)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) HandlePostSearchAvailability(w http.ResponseWriter, r *http.Request) {

	/*
		// cara mengambil form pada GO
		mStart := r.Form.Get("start")
		mEnd := r.Form.Get("end")

		dataCheck := model.DataCheckAvailability{
			StartDate:    mStart,
			EndDate:      mEnd,
			StatusProses: true,
		}

		data := make(map[string]interface{})
		data["dataCheck"] = dataCheck
	*/

	//fmt.Println(data)

	// jika proses berhasil maka tambahkan ini
	m.App.Session.Put(r.Context(), "flash", "Data BARHASIL ditambahkan.")

	td := render.AddDefaultData(&model.TemplateData{}, r)

	// ambil data halaman proses untuk tes menampilkan data
	//mytemplate, _ := ViewTemplate("templates/proses.check.sakur")
	mytemplate, _ := ViewTemplate("templates/home.page.sakur")

	err := mytemplate.Execute(w, td)
	if err != nil {
		fmt.Println(err)
	}
}

func (m *Repository) HandleReservationSummary(w http.ResponseWriter, r *http.Request) {

	mytemplate, _ = ViewTemplate("templates/reservation-summary.page.sakur")

	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(model.Reservation)

	if !ok {
		//log.Println("Cannot get item from session")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	td := render.AddDefaultData(&model.TemplateData{
		Data: data,
	}, r)

	err := mytemplate.Execute(w, td)

	if err != nil {
		log.Println(err)
	}
}
