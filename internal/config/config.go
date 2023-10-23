package config

import (
	"fmt"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache     bool
	InfoLog      *log.Logger
	InProduction bool
	Session      *scs.SessionManager
}

func GetInfoProgram(port string) {
	fmt.Println("")
	fmt.Println("Copyright(c) 2021 By Sakur Technologies")
	fmt.Println("Posyandu Lansia Versi 1.0 (PosLan)")
	fmt.Println("Versi Gratis (Freeware) dengan AKTIVASI")
	fmt.Println("Using: Microsoft Access (Daya tampung 2 Gb maks.)")
	fmt.Println("       Aplikasi Berbasis Web")
	fmt.Println("Penulis Kode Komputer oleh Stendy B. Sakur")
	fmt.Println("********************** <!> ****************************")
	fmt.Println("Untuk Full version please call us: ")
	fmt.Println("Email 	 : sakur.stendy@gmail.com")
	fmt.Println("Telegram: https://t.me/stendysakur")
	fmt.Println("Full Version: ")
	fmt.Println("{Using: MySQL / PostgreSQL {WebBase Or Qt/C++/QML}}")
	fmt.Println("")
	fmt.Printf("***** Access lewat browser dengan localhost%v *****\n", port)
	fmt.Println("")
}
