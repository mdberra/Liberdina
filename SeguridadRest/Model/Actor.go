package Model

import (
	"database/sql"
	"time"
)

type Actor struct {
	Id           int64     `json:"id"`
	IpAddress    string    `json:"ipAddress"`
	Tipo         string    `json:"tipo"`
	Dni          int64     `json:"dni"`
	Pin          string    `json:"pin"`
	Email        string    `json:"email"`
	FechaEnrolar time.Time `json:"fechaEnrolar"`
}

const ActorTipoPaciente string = "Paciente"
const ActorTipoMedico string = "Medico"

func (a *Actor) New() {
	a.Id = 0
	a.IpAddress = ""
	a.Tipo = ""
	a.Dni = 0
	a.Pin = ""
	a.Email = ""
	a.FechaEnrolar = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
}
func (a *Actor) GetTipo() string {
	return a.Tipo
}
func (a *Actor) Create(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("INSERT Actor SET ipAddress=?, tipo=?, dni=?, pin=?, email=?, fechaEnrolar=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(a.IpAddress, a.Tipo, a.Dni, a.Pin, a.Email, a.FechaEnrolar)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err == nil {
		a.Id = id
	}
	err = tx.Commit()
	return err
}
func (a *Actor) GetActorDni(db *sql.DB) (Actor, error) {
	var ent Actor
	q := `select *
			from Actor AS a
			where a.dni =?`
	rows, err := db.Query(q, a.Dni)
	if err != nil {
		return ent, err
	}
	defer rows.Close()

	for rows.Next() {
		ent, err = a.parsear(rows)
		if err != nil {
			return ent, err
		}
	}
	return ent, err
}
func (a *Actor) GetActorEmail(db *sql.DB) (Actor, error) {
	var ent Actor
	q := `select *
			from Actor AS a
			where a.email =?`
	rows, err := db.Query(q, a.Email)
	if err != nil {
		return ent, err
	}
	defer rows.Close()

	for rows.Next() {
		ent, err = a.parsear(rows)
		if err != nil {
			return ent, err
		}
	}
	return ent, err
}
func (a *Actor) GetActores(db *sql.DB) ([]Actor, error) {
	// preparar el query para ejecutarlo muchas veces,  ahorra tiempo
	q := `select *
			from Actor`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ent Actor
	e := []Actor{}
	for rows.Next() {
		ent, err = a.parsear(rows)
		if err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
func (a *Actor) GetActor(db *sql.DB) (Actor, error) {
	var ent Actor
	q := `select *
			from Actor AS a
			where a.idActor =?`
	rows, err := db.Query(q, a.Id)
	if err != nil {
		return ent, err
	}
	defer rows.Close()

	for rows.Next() {
		ent, err = a.parsear(rows)
		if err != nil {
			return ent, err
		}
	}
	return ent, err
}
func (a *Actor) parsear(rows *sql.Rows) (Actor, error) {
	var ent Actor
	err := rows.Scan(&ent.Id, &ent.IpAddress, &ent.Tipo, &ent.Dni, &ent.Pin, &ent.Email, &ent.FechaEnrolar)
	return ent, err
}
