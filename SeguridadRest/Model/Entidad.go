package Model

import (
	"database/sql"
	"time"
)

type Entidad interface {
	getTipo() string
	New()
}

const ConstRecetaDigital string = "RecetaDigital"

type RecetaDigital struct {
	Id            int64     `json:"id"`
	Tipo          string    `json:"tipo"`
	IdMedico      int64     `json:"idMedico"`
	IdPaciente    int64     `json:"idPaciente"`
	IdMedicamento int64     `json:"idMedicamento"`
	Fecha         time.Time `json:"fecha"`
}

func (r RecetaDigital) New() {
	r.Id = 0
	r.Tipo = ""
	r.IdMedico = 0
	r.IdPaciente = 0
	r.IdMedicamento = 0
	r.Fecha = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
}
func (r RecetaDigital) getTipo() string {
	return ConstRecetaDigital
}

func (r RecetaDigital) Create(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("INSERT Entidad SET tipo=?, idMedico=?, idPaciente=?, idMedicamento=?, fecha=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(r.Tipo, r.IdMedico, r.IdPaciente, r.IdMedicamento, r.Fecha)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err == nil {
		r.Id = id
	}
	err = tx.Commit()
	return err
}
func (r RecetaDigital) GetEntidad(db *sql.DB) (Entidad, error) {
	var ent RecetaDigital
	q := `select *
			from RecetaDigital AS r
			where r.id =?`
	rows, err := db.Query(q, r.Id)
	if err != nil {
		return ent, err
	}
	defer rows.Close()

	for rows.Next() {
		ent, err = r.parsear(rows)
		if err != nil {
			return ent, err
		}
	}
	return ent, err
}
func (r RecetaDigital) parsear(rows *sql.Rows) (RecetaDigital, error) {
	var ent RecetaDigital
	err := rows.Scan(&ent.Id, &ent.Tipo, &ent.IdMedico, &ent.IdPaciente, &ent.IdMedicamento, &ent.Fecha)
	return ent, err
}
