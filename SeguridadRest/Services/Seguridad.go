package Services

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Liberdina/Recedig/Data"
	"github.com/Liberdina/Seguridad/Model"
	jwt "github.com/dgrijalva/jwt-go"
	qrcode "github.com/skip2/go-qrcode"
	"github.com/subosito/gotenv"
	"golang.org/x/crypto/bcrypt"
)

var (
	conexion Model.Conexion
)

type Seguridad struct {
}

func init() {
	gotenv.Load()
}

/**

BuscarEntidad: Token  -->  Tipo, Entidad

ActualizarEntidad; Token, Estado --> CerrarToken
*/
/**
Enrolar: Actor(IpAddress, Tipo, Dni, Pin, Email)    Paciente/Medico
*/
func (s *Seguridad) Enrolar(actor Model.Actor) (Model.Actor, error) {
	if actor.IpAddress == "" {
		return actor, errors.New("IpAddress sin informar")
	}
	if actor.Tipo == "" {
		return actor, errors.New("Tipo de Actor sin informar")
	}
	if actor.Dni == 0 {
		return actor, errors.New("Dni sin informar")
	}
	if actor.Pin == "" {
		return actor, errors.New("IpAddress sin informar")
	}
	if actor.Email == "" {
		return actor, errors.New("Email sin informar")
	}
	//----------------------------------------------------
	if err := conexion.Open(conexion.GetEsquemaSeguridad()); err != nil {
		panic(err)
	} else {
		defer conexion.Close()
	}
	//----------------------------------------------------
	// Valido que no exista un Actor con este Dni
	a, err := actor.GetActorDni(conexion.Db)
	if err != nil {
		panic(err)
	}
	if a.Id > 0 {
		return actor, errors.New(fmt.Sprintf("Ya existe un Actor con Dni %v", actor.Dni))
	}
	//----------------------------------------------------
	// Valido que no exista un Actor con este Email
	a, err = actor.GetActorEmail(conexion.Db)
	if err != nil {
		panic(err)
	}
	if a.Id > 0 {
		return actor, errors.New(fmt.Sprintf("Ya existe un Actor con Email %v", actor.Email))
	}
	//----------------------------------------------------
	// Encriptamos el PIN
	hashPin, err := bcrypt.GenerateFromPassword([]byte(actor.Pin), 10)
	if err != nil {
		log.Panic(err)
	}
	actor.Pin = string(hashPin)
	actor.FechaEnrolar = time.Now().Add(time.Hour * -3)

	if err := actor.Create(conexion.Db); err != nil {
		panic(err)
	}
	return actor, nil
}

/**
El Medico solicita generar una Entidad (receta) que se genera, encripta devolviendo un token con valides 1 mes
. verificar el pin del Paciente
. verificar el pin del Medico
. verificar el medicamento
. generarToken

Medico [Actor(Dni Pin)] --> GenerarEntidad: [Entidad] --> CompareHashandPassword, GenerarToken(Ip, Time, Actor(Dni, Pin), Tipo, Entidad)

*/
func (s *Seguridad) GenerarEntidad(ipAddress string, pacienteIn Model.Actor, medicoIn Model.Actor, medicamentoIn Data.Medicamento) ([]byte, error) {
	var png []byte
	//	var tokenJWT Model.JWT
	var err error

	var recetaDigital Model.RecetaDigital
	recetaDigital.New()

	var token Model.Token
	token.New()
	//----------------------------------------------------
	if ipAddress == "" {
		return png, errors.New("IpAddress sin informar")
	}
	if pacienteIn.Id == 0 {
		return png, errors.New("Paciente sin informar")
	}
	if medicoIn.Id == 0 {
		return png, errors.New("Medico sin informar")
	}
	if medicamentoIn.Id == 0 {
		return png, errors.New("Medicamento sin informar")
	}
	//----------------------------------------------------
	if err := conexion.Open(conexion.GetEsquemaSeguridad()); err != nil {
		panic(err)
	} else {
		defer conexion.Close()
	}
	//----------------------------------------------------
	paciente, err := pacienteIn.GetActorDni(conexion.Db)
	if err != nil {
		return png, errors.New(fmt.Sprintf("El Paciente con Dni %v no existe", pacienteIn.Dni))
	}
	medico, err := medicoIn.GetActorDni(conexion.Db)
	if err != nil {
		return png, errors.New(fmt.Sprintf("El Medico con Dni %v no existe", medicoIn.Dni))
	}
	//----------------------------------------------------
	// verificar Pin
	err = bcrypt.CompareHashAndPassword([]byte(paciente.Pin), []byte(pacienteIn.Pin))
	if err != nil {
		return png, errors.New("El PIN del Paciente es incorrecto")
	}
	err = bcrypt.CompareHashAndPassword([]byte(medico.Pin), []byte(medicoIn.Pin))
	if err != nil {
		return png, errors.New("El PIN del Medico es incorrecto")
	}
	//----------------------------------------------------
	recetaDigital.Tipo = Model.ConstRecetaDigital
	recetaDigital.IdMedico = medico.Id
	recetaDigital.IdPaciente = paciente.Id
	recetaDigital.IdMedicamento = medicamentoIn.Id
	recetaDigital.Fecha = time.Now().Add(time.Hour * -3)

	token.IpAddress = ipAddress
	token.Fecha = recetaDigital.Fecha
	token.TipoEntidad = recetaDigital.Tipo
	token.TEntidad = recetaDigital
	err = token.GenerarToken()
	if err != nil {
		return png, errors.New(fmt.Sprintf("Al generar el token %v se produjo el error %v", token, err))
	}
	// con el payload se calcula el token que se colocan en el atributo token dentro del struct
	// luego se genera el qrcode
	// entonces cuando se reciba un QR, se valida el atributo token para ver si el resto
	// de los atributo son validos
	log.Println(token.String())
	png, err = qrcode.Encode(token.String(), qrcode.Medium, 256)

	log.Println("QRCode(Token(payload)")
	log.Println(png)

	return png, nil
}

/**
A JWT token is simply a signed JSON object. It can be used anywhere such a thing is useful.
*/
func (s *Seguridad) ValidarToken(authToken Model.JWT) error {
	token, err := jwt.Parse(authToken.Token, funcion)
	if err != nil {
		return errors.New("Token no valido")
	}
	if !token.Valid {
		return errors.New("Token no valido")
	}
	return nil
}
func funcion(token *jwt.Token) (interface{}, error) {
	//	spew.Dump(token)
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("SigningMethodHMAC invalido")
	}
	return []byte(os.Getenv("Secreta")), nil
}

func (s *Seguridad) ActualizarEntidad() {
}
