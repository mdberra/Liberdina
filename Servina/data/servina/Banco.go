package data

import (
	"database/sql"
)

type Banco struct {
	ID_BANCO           int64  `json:"idBanco"`
	CODIGO             string `json:"codigo"`
	DESCRIPCION        string `json:"descripcion"`
	ID_LOCALIZACION    int64  `json:"idLocalizacion"`
	CODIGO_DEBITO      int    `json:"codigoDebito"`
	DESCRIP_PRESTACION string `json:"descripPrestacion"`
	BANCO_RECAUDADOR   string `json:"bancoRecaudador"`
}

func (n *Banco) CleanBanco() {
	n.ID_BANCO = 0
	n.CODIGO = ""
	n.DESCRIPCION = ""
	n.ID_LOCALIZACION = 0
	n.CODIGO_DEBITO = 0
	n.DESCRIP_PRESTACION = ""
	n.BANCO_RECAUDADOR = ""
}

func (n *Banco) GetBanco(db *sql.DB) (Banco, error) {
	var ent Banco
	q := `select * from BANCO where ID_BANCO=?`
	rows, err := db.Query(q, n.ID_BANCO)
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

func (n *Banco) UpdateBanco(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("update BANCO set CODIGO=?, DESCRIPCION=?, ID_LOCALIZACION=?, CODIGO_DEBITO=?, DESCRIP_PRESTACION=?, BANCO_RECAUDADOR=? where ID_BANCO=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.CODIGO, n.DESCRIPCION, n.ID_LOCALIZACION, n.CODIGO_DEBITO, n.DESCRIP_PRESTACION, n.BANCO_RECAUDADOR, n.ID_BANCO)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}

func (n *Banco) DeleteBanco(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("delete from BANCO where ID_BANCO=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.ID_BANCO)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}

func (n *Banco) CreateBanco(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("INSERT BANCO set ID_BANCO=?, CODIGO=?, DESCRIPCION=?, ID_LOCALIZACION=?, CODIGO_DEBITO=?, DESCRIP_PRESTACION=?, BANCO_RECAUDADOR=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.ID_BANCO, n.CODIGO, n.DESCRIPCION, n.ID_LOCALIZACION, n.CODIGO_DEBITO, n.DESCRIP_PRESTACION, n.BANCO_RECAUDADOR)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err == nil {
		n.ID_BANCO = id
	}

	err = tx.Commit()
	return err
}

//func GetBancos(db *sql.DB, start, count int) ([]Banco, error) {
func (n *Banco) GetBancos(db *sql.DB) ([]Banco, error) {
	q := `select * from BANCO`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ent Banco
	e := []Banco{} //array
	for rows.Next() {
		ent, err = n.parsear(rows)
		if err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
func (n *Banco) parsear(rows *sql.Rows) (Banco, error) {
	var ent Banco
	err := rows.Scan(&ent.ID_BANCO, &ent.CODIGO, &ent.DESCRIPCION, &ent.ID_LOCALIZACION, &ent.CODIGO_DEBITO, &ent.DESCRIP_PRESTACION, &ent.BANCO_RECAUDADOR)
	return ent, err
}
