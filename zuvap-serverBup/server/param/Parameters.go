package param

import (
	"log"
	"os"
)

var (
	ambiente string
)

type Parameters struct {
	//Server
	IP   string
	Port string
	// DB
	DbDriver   string
	DbIP       string
	DbUser     string
	DbPassword string
	DbName     string
	//Pay
	Token             string
	URL               string
	URLPostCreate     string
	URLGetPlanDetails string
	URLPostPay        string
}

func (param *Parameters) SetParameters() (*Parameters, error) {
	if ambiente == "" {
		ambiente = os.Getenv("Ambiente")
		if ambiente == "" {
			log.Fatalln("No se encontro el Ambiente")
			panic("ambiente")
		}
	}
	//Server
	param.IP = os.Getenv("IP")
	param.Port = os.Getenv("Port")

	// DB
	param.DbDriver = os.Getenv("DbDriver")
	param.DbIP = os.Getenv("DbIP")
	param.DbUser = os.Getenv("DbUser")
	param.DbPassword = os.Getenv("DbPassword")
	param.DbName = os.Getenv("DbName")

	// Sistema de Pagos

	param.Token = os.Getenv("Token")
	param.URL = os.Getenv("URL")
	param.URLPostCreate = param.URL + os.Getenv("URLPostCreate")
	param.URLGetPlanDetails = param.URL + os.Getenv("URLGetPlanDetails")
	param.URLPostPay = param.URL + os.Getenv("URLPostPay")

	return param, nil
}
