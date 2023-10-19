package model

import (
	"strings"
)

type Personal struct {
	Nama   string
	Status string
}

type Newsline struct {
	ID      int
	MyTitle string
	Summery string
	Desc    string
	Writer  Personal
}

func (n *Newsline) SetConstructor(id int, t, s, d, nm, sts string) {
	n.SetTitle(t)
	n.SetSummary(s)
	n.SetDesc(d)
	n.SetWriter(nm, sts)
	n.SetID(id)
}

func (n *Newsline) SetID(id int) {
	if id > 0 {
		n.ID = id
	}
}

func (n *Newsline) SetTitle(t string) {
	n.MyTitle = t
}

func (n *Newsline) SetSummary(s string) {
	n.Summery = s
}

func (n *Newsline) SetDesc(d string) {
	n.Desc = d
}

func (n *Newsline) SetWriter(nm, sts string) {
	n.Writer.Nama = nm
	n.Writer.Status = sts
}

func (n *Newsline) GetID() int {
	return n.ID
}

func (n *Newsline) GetTitle() string {
	return strings.ToUpper(n.MyTitle)
}

func (n *Newsline) GetSummery() string {
	return n.Summery
}

func (n *Newsline) GetDesc() string {
	return n.Desc
}

func (n *Newsline) GetNama() string {
	nm := n.Writer.Nama
	return strings.ToUpper(nm)
}

func (n *Newsline) GetStatus() string {
	return strings.ToUpper(n.Writer.Status)
}

/* STRUKTUR DATA UNTUK KARYAWAN */
type Karyawan struct {
	NamaKaryawan  string
	UmurKaryawan  int
	EmailKaryawan string
	GajiBulanan   float32
}

func (k *Karyawan) SetConstructor(n, e string, u int, g float32) {
	k.SetNama(n)
	k.SetEmail(e)
	k.SetUmur(u)
	k.SetGaji(g)
}

func (k *Karyawan) SetNama(n string) {
	k.NamaKaryawan = n
}

func (k *Karyawan) SetUmur(u int) {
	k.UmurKaryawan = u
}

func (k *Karyawan) SetEmail(e string) {
	k.EmailKaryawan = e
}

func (k *Karyawan) SetGaji(g float32) {
	k.GajiBulanan = g
}

/* STRUKTUR DATA UNTUK HOME PAGE */
type DataStruk struct {
	MyID       SessionProperty
	MyKaryawan []Karyawan
	MyNews     []Newsline
}

type NewsDetail struct {
	IDParam int
	MyNews  []Newsline
	MyID    SessionProperty
}

/* DIGUNAKAN UNTUK HALAMAN YANG MENGGUANAKAN SESSION */
type SessionProperty struct {
	LabelName string
	Status    string
	RemoteIP  string
	UserID    string
}

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
