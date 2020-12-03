package data

import (
	"database/sql"
	"time"
)

type Delegacion struct {
	ID_DELEGACION   int64     `json:"idDelegacion"`
	CODIGO          string    `json:"codigo"`
	DESCRIPCION     string    `json:"descripcion"`
	FECHA_ALTA      time.Time `json:"fechaAlta"`
	ID_LOCALIZACION int64     `json:"idLocalizacion"`
	ID_ENTIDAD      int64     `json:"idEntidad"`
	ID_BANCO        int64     `json:"idBanco"`
	UTILIZAR        int       `json:"utilizar"`
}

func (n *Delegacion) CleanDelegacion() {
	n.ID_DELEGACION = 0
	n.CODIGO = ""
	n.DESCRIPCION = ""
	n.FECHA_ALTA = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	n.ID_LOCALIZACION = 0
	n.ID_ENTIDAD = 0
	n.ID_BANCO = 0
	n.UTILIZAR = 0
}

func (n *Delegacion) GetDelegacion(db *sql.DB) (Delegacion, error) {
	var ent Delegacion
	q := `select * from DELEGACION where ID_DELEGACION=?`
	rows, err := db.Query(q, n.ID_DELEGACION)
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
	return ent, nil
}

func (n *Delegacion) UpdateDelegacion(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("update DELEGACION set CODIGO=?, DESCRIPCION=?, FECHA_ALTA=?, ID_LOCALIZACION=?, ID_ENTIDAD=?, ID_BANCO=?, UTILIZAR=? where ID_DELEGACION=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.CODIGO, n.DESCRIPCION, n.FECHA_ALTA, n.ID_LOCALIZACION, n.ID_ENTIDAD, n.ID_BANCO, n.UTILIZAR, n.ID_DELEGACION)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}

func (n *Delegacion) DeleteDelegacion(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("delete from DELEGACION where ID_DELEGACION=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.ID_DELEGACION)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}

func (n *Delegacion) CreateDelegacion(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("INSERT DELEGACION set ID_DELEGACION=?, CODIGO=?, DESCRIPCION=?, FECHA_ALTA=?, ID_LOCALIZACION=?, ID_ENTIDAD=?, ID_BANCO=?, UTILIZAR=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.ID_DELEGACION, n.CODIGO, n.DESCRIPCION, n.FECHA_ALTA, n.ID_LOCALIZACION, n.ID_ENTIDAD, n.ID_BANCO, n.UTILIZAR)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err == nil {
		n.ID_DELEGACION = id
	}

	err = tx.Commit()
	return err
}

//func GetDelegacions(db *sql.DB, start, count int) ([]Delegacion, error) {
func (n *Delegacion) GetDelegaciones(db *sql.DB) ([]Delegacion, error) {
	q := `select * from DELEGACION`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ent Delegacion
	e := []Delegacion{} //array
	for rows.Next() {
		ent, err = n.parsear(rows)
		if err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
func (n *Delegacion) parsear(rows *sql.Rows) (Delegacion, error) {
	var ent Delegacion
	err := rows.Scan(&ent.ID_DELEGACION, &ent.CODIGO, &ent.DESCRIPCION, &ent.FECHA_ALTA, &ent.ID_LOCALIZACION, &ent.ID_ENTIDAD, &ent.ID_BANCO, &ent.UTILIZAR)
	return ent, err
}
