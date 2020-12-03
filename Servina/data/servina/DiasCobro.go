package data

import (
	"database/sql"
	"time"
)

type DiasCobro struct {
	ID_DIASCOBRO  int64     `json:"idDiasCobro"`
	ID_DELEGACION int64     `json:"idDelegacion"`
	FECHA_DISPARO time.Time `json:"fechaDisparo"`
}

func (n *DiasCobro) CleanDiasCobro() {
	n.ID_DIASCOBRO = 0
	n.ID_DELEGACION = 0
	n.FECHA_DISPARO = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
}

func (n *DiasCobro) GetDiasCobro(db *sql.DB) (DiasCobro, error) {
	var ent DiasCobro
	q := `select * from DIAS_COBRO where ID_DIASCOBRO=?`
	rows, err := db.Query(q, n.ID_DIASCOBRO)
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

func (n *DiasCobro) UpdateDiasCobro(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("update DIAS_COBRO set ID_DELEGACION=?, FECHA_DISPARO=? where ID_DIASCOBRO=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.ID_DELEGACION, n.FECHA_DISPARO, n.ID_DIASCOBRO)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}

func (n *DiasCobro) DeleteDiasCobro(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("delete from DIAS_COBRO where ID_DIASCOBRO=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.ID_DIASCOBRO)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}

func (n *DiasCobro) CreateDiasCobro(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("INSERT DIAS_COBRO set ID_DIASCOBRO=?, ID_DELEGACION=?, FECHA_DISPARO=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.ID_DIASCOBRO, n.ID_DELEGACION, n.FECHA_DISPARO)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err == nil {
		n.ID_DIASCOBRO = id
	}

	err = tx.Commit()
	return err
}

//func GetDiasCobros(db *sql.DB, start, count int) ([]DiasCobro, error) {
func (n *DiasCobro) GetDiasCobroDelegacion(db *sql.DB) ([]DiasCobro, error) {
	q := `select * from DIAS_COBRO where ID_DELEGACION=?`
	rows, err := db.Query(q, n.ID_DELEGACION)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ent DiasCobro
	e := []DiasCobro{} //array
	for rows.Next() {
		ent, err = n.parsear(rows)
		if err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
func (n *DiasCobro) parsear(rows *sql.Rows) (DiasCobro, error) {
	var ent DiasCobro
	err := rows.Scan(&ent.ID_DIASCOBRO, &ent.ID_DELEGACION, &ent.FECHA_DISPARO)
	return ent, err
}
