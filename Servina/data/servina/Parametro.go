package data

import (
	"database/sql"
)

type Parametro struct {
	ID_PARAMETRO   int64  `json:"idParametro"`
	INDICE         int    `json:"indice"`
	CUIT_EMPRESA   string `json:"cuitEmpresa"`
	NOMBRE_EMPRESA string `json:"nombreEmpresa"`
	MONEDA         string `json:"moneda"`
	ULT_SERVICIO   int    `json:"ultServicio"`
}

func (n *Parametro) CleanParametro() {
	n.ID_PARAMETRO = 0
	n.INDICE = 0
	n.CUIT_EMPRESA = ""
	n.NOMBRE_EMPRESA = ""
	n.MONEDA = ""
	n.ULT_SERVICIO = 0
}

func (n *Parametro) GetParametro(db *sql.DB) (Parametro, error) {
	var ent Parametro
	q := `select * PARAMETRO where ID_PARAMETRO=?`
	rows, err := db.Query(q, n.ID_PARAMETRO)
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

func (n *Parametro) UpdateParametro(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("update PARAMETRO set INDICE=?, CUIT_EMPRESA=?, NOMBRE_EMPRESA=?, MONEDA=?, ULT_SERVICIO=? where ID_PARAMETRO=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.INDICE, n.CUIT_EMPRESA, n.NOMBRE_EMPRESA, n.MONEDA, n.ULT_SERVICIO, n.ID_PARAMETRO)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}

func (n *Parametro) DeleteParametro(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("delete PARAMETRO where ID_PARAMETRO=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.ID_PARAMETRO)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}

func (n *Parametro) CreateParametro(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("INSERT PARAMETRO set ID_PARAMETRO=?, INDICE=?, CUIT_EMPRESA=?, NOMBRE_EMPRESA=?, MONEDA=?, ULT_SERVICIO=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.ID_PARAMETRO, n.INDICE, n.CUIT_EMPRESA, n.NOMBRE_EMPRESA, n.MONEDA, n.ULT_SERVICIO)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err == nil {
		n.ID_PARAMETRO = id
	}

	err = tx.Commit()
	return err
}

//func GetParametros(db *sql.DB, start, count int) ([]Parametro, error) {
func (n *Parametro) GetParametros(db *sql.DB) ([]Parametro, error) {
	q := `select * from PARAMETRO`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ent Parametro
	e := []Parametro{} //array
	for rows.Next() {
		ent, err = n.parsear(rows)
		if err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
func (n *Parametro) parsear(rows *sql.Rows) (Parametro, error) {
	var ent Parametro
	err := rows.Scan(&ent.ID_PARAMETRO, &ent.INDICE, &ent.CUIT_EMPRESA, &ent.NOMBRE_EMPRESA, &ent.MONEDA, &ent.ULT_SERVICIO)
	return ent, err
}
