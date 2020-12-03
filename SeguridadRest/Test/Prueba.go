package main

import (
	"log"

	"github.com/Liberdina/Recedig/Data"
	"github.com/Liberdina/Seguridad/Model"
	"github.com/Liberdina/Seguridad/Services"
)

var (
	paciente    Model.Actor
	medico      Model.Actor
	medicamento Data.Medicamento

	seguridad Services.Seguridad
)

func main() {
	paciente.New()
	paciente.IpAddress = "180.181.182.183"
	paciente.Tipo = Model.ActorTipoPaciente
	paciente.Dni = 12345678
	paciente.Pin = "este es un pinPaciente"
	paciente.Email = "paciente@individuo.com"
	paciente, err := seguridad.Enrolar(paciente)
	if err != nil {
		log.Println(err)
	}

	medico.New()
	medico.IpAddress = "190.191.192.193"
	medico.Tipo = Model.ActorTipoMedico
	medico.Dni = 87654321
	medico.Pin = "este es un pinMedico"
	medico.Email = "medico@individuo.com"
	medico, err := seguridad.Enrolar(medico)
	if err != nil {
		log.Println(err)
	}

	medicamento.Id = 1
	medicamento.Nombre = "Ibuprofeno 200"
	medicamento.Droga = "Ibuprofeno"

	paciente.Pin = "este es un pinPaciente"
	medico.Pin = "este es un pinMedico"
	png, err := seguridad.GenerarEntidad(medico.IpAddress, paciente, medico, medicamento)
	if err != nil {
		log.Println(err)
	}
	log.Println(png)
	//	spew.Dump(png)
}
