package Model

import (
	"encoding/json"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Token struct {
	JwtToken    string    `json:"JwtToken"`
	IpAddress   string    `json:"ipAddress"`
	Fecha       time.Time `json:"fecha"`
	TipoEntidad string    `json:"tipoEntidad"`
	TEntidad    Entidad   `json:"entidad"`
}
type JWT struct {
	Token string `json:"token"`
}

func (t *Token) String() string {
	bytes, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
func (t *Token) New() {
	t.JwtToken = ""
	t.IpAddress = ""
	t.Fecha = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	t.TipoEntidad = ""
	t.TEntidad = nil
}
func (t *Token) GenerarToken() error {
	//	var jwtAux JWT
	var mapClaims = jwt.MapClaims{
		"ipAddress":   t.IpAddress,
		"fecha":       t.Fecha,
		"tipoEntidad": t.TipoEntidad,
		"entidad":     t.TEntidad,

		"iss": "course",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	tokenString, err := token.SignedString([]byte(os.Getenv("Secreta")))
	if err != nil {
		log.Fatal(err)
	}
	//	jwtAux.Token = tokenString
	t.JwtToken = tokenString
	return nil
}
