package Data

import (
	"database/sql"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Usuario struct {
	IdUsuario int64     `json:"idUsuario"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Fecha     time.Time `json:"fecha"`
}
type JWT struct {
	Token string `json:"token"`
}

func (n *Usuario) CleanUsuario() {
	n.IdUsuario = 0
	n.Email = ""
	n.Password = ""
	n.Fecha = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
}
func (n *Usuario) GenerarToken() (JWT, error) {
	var jwtAux JWT
	var mapClaims = jwt.MapClaims{
		"email": n.Email,
		"iss":   "course",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	tokenString, err := token.SignedString([]byte(os.Getenv("Secreta")))
	if err != nil {
		log.Fatal(err)
	}
	jwtAux.Token = tokenString
	return jwtAux, nil
}
func (n *Usuario) CreateUsuario(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("INSERT Usuario SET email=?, password=?, fecha=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.Email, n.Password, n.Fecha)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err == nil {
		n.IdUsuario = id
	}
	err = tx.Commit()
	return err
}
func (n *Usuario) GetUsuarioEmail(db *sql.DB) (Usuario, error) {
	var ent Usuario
	q := `select c.idUsuario, c.email, c.password, c.fecha
			from Usuario AS c
			where c.email =?`
	rows, err := db.Query(q, n.Email)
	if err != nil {
		return ent, err
	}
	defer rows.Close()

	for rows.Next() {
		ent, err = n.parsear(rows)
		if err != nil {
			return ent, err
		}
	}
	return ent, err
}
func (n *Usuario) GetUsuarios(db *sql.DB) ([]Usuario, error) {
	// preparar el query para ejecutarlo muchas veces,  ahorra tiempo
	q := `select c.idUsuario, c.email, c.password, c.fecha
			from Usuario AS c`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ent Usuario
	e := []Usuario{}
	for rows.Next() {
		ent, err = n.parsear(rows)
		if err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
func (n *Usuario) GetUsuario(db *sql.DB) (Usuario, error) {
	var ent Usuario
	q := `select c.idUsuario, c.email, c.password, c.fecha
			from Usuario AS c
			where c.idUsuario =?`
	rows, err := db.Query(q, n.IdUsuario)
	if err != nil {
		return ent, err
	}
	defer rows.Close()

	for rows.Next() {
		ent, err = n.parsear(rows)
		if err != nil {
			return ent, err
		}
	}
	return ent, err
}
func (n *Usuario) parsear(rows *sql.Rows) (Usuario, error) {
	var ent Usuario
	err := rows.Scan(&ent.IdUsuario, &ent.Email, &ent.Password, &ent.Fecha)
	return ent, err
}
