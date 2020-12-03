package data

import (
	"database/sql"
)

type Entidad struct {
	ID_ENTIDAD  int64  `json:"idEntidad"`
	CODIGO      string `json:"codigo"`
	DESCRIPCION string `json:"descripcion"`
}

func (n *Entidad) CleanEntidad() {
	n.ID_ENTIDAD = 0
	n.CODIGO = ""
	n.DESCRIPCION = ""
}

func (n *Entidad) GetEntidad(db *sql.DB) (Entidad, error) {
	var ent Entidad
	q := `select * from Entidad where CODIGO=?`
	rows, err := db.Query(q, n.CODIGO)
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

func (n *Entidad) UpdateEntidad(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("update Entidad set CODIGO=?, DESCRIPCION=? where ID_ENTIDAD=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.CODIGO, n.DESCRIPCION, n.ID_ENTIDAD)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}

func (n *Entidad) DeleteEntidad(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("delete from Entidad where ID_Entidad=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.ID_ENTIDAD)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}

func (n *Entidad) CreateEntidad(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("INSERT Entidad set ID_ENTIDAD=?, CODIGO=?, DESCRIPCION=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.ID_ENTIDAD, n.CODIGO, n.DESCRIPCION)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err == nil {
		n.ID_ENTIDAD = id
	}

	err = tx.Commit()
	return err
}

//func GetEntidads(db *sql.DB, start, count int) ([]Entidad, error) {
func (n *Entidad) GetEntidads(db *sql.DB) ([]Entidad, error) {
	q := `select * from ENTIDAD`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ent Entidad
	e := []Entidad{} //array
	for rows.Next() {
		ent, err = n.parsear(rows)
		if err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
func (n *Entidad) parsear(rows *sql.Rows) (Entidad, error) {
	var ent Entidad
	err := rows.Scan(&ent.ID_ENTIDAD, &ent.CODIGO, &ent.DESCRIPCION)
	return ent, err
}
