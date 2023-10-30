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
	fmt.Printf("***** Access lewat browser dengan localhost%v *****\n", port)
	fmt.Println("")
}
