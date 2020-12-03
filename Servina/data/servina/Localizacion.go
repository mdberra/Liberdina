package data

import (
	"database/sql"
)

type Localizacion struct {
	ID_LOCALIZACION int64  `json:"idLocalizacion"`
	CALLE           string `json:"calle"`
	NRO             string `json:"nro"`
	PISO            string `json:"piso"`
	DEPTO           string `json:"depto"`
	TELEF_LINEA     string `json:"telefLinea"`
	TELEF_CELULAR   string `json:"telefCelular"`
	COD_POSTAL      string `json:"codPostal"`
	BARRIO          string `json:"barrio"`
	LOCALIDAD       string `json:"localidad"`
	PROVINCIA       string `json:"provincia"`
	PAIS            string `json:"pais"`
}

func (n *Localizacion) CleanLocalizacion() {
	n.ID_LOCALIZACION = 0
	n.CALLE = ""
	n.NRO = ""
	n.PISO = ""
	n.DEPTO = ""
	n.TELEF_LINEA = ""
	n.TELEF_CELULAR = ""
	n.COD_POSTAL = ""
	n.BARRIO = ""
	n.LOCALIDAD = ""
	n.PROVINCIA = ""
	n.PAIS = ""
}

func (n *Localizacion) GetLocalizacion(db *sql.DB) (Localizacion, error) {
	var ent Localizacion
	q := `select * LOCALIZACION where ID_LOCALIZACION=?`
	rows, err := db.Query(q, n.ID_LOCALIZACION)
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

func (n *Localizacion) UpdateLocalizacion(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("update LOCALIZACION set CALLE=?, NRO=?, PISO=?, DEPTO=?, TELEF_LINEA=?, TELEF_CELULAR=?, COD_POSTAL=?, BARRIO=?, LOCALIDAD=?, PROVINCIA=?, PAIS=? where ID_LOCALIZACION=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.CALLE, n.NRO, n.PISO, n.DEPTO, n.TELEF_LINEA, n.TELEF_CELULAR, n.COD_POSTAL, n.BARRIO, n.LOCALIDAD, n.PROVINCIA, n.PAIS, n.ID_LOCALIZACION)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}

func (n *Localizacion) DeleteLocalizacion(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("delete LOCALIZACION where ID_LOCALIZACION=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.ID_LOCALIZACION)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}

func (n *Localizacion) CreateLocalizacion(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("INSERT LOCALIZACION set ID_LOCALIZACION=?, CALLE=?, NRO=?, PISO=?, DEPTO=?, TELEF_LINEA=?, TELEF_CELULAR=?, COD_POSTAL=?, BARRIO=?, LOCALIDAD=?, PROVINCIA=?, PAIS=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.ID_LOCALIZACION, n.CALLE, n.NRO, n.PISO, n.DEPTO, n.TELEF_LINEA, n.TELEF_CELULAR, n.COD_POSTAL, n.BARRIO, n.LOCALIDAD, n.PROVINCIA, n.PAIS)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err == nil {
		n.ID_LOCALIZACION = id
	}

	err = tx.Commit()
	return err
}

//func GetDelegacions(db *sql.DB, start, count int) ([]Localizacion, error) {
func (n *Localizacion) GetDelegacions(db *sql.DB) ([]Localizacion, error) {
	q := `select * LOCALIZACION`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ent Localizacion
	e := []Localizacion{} //array
	for rows.Next() {
		ent, err = n.parsear(rows)
		if err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
func (n *Localizacion) parsear(rows *sql.Rows) (Localizacion, error) {
	var ent Localizacion
	err := rows.Scan(&ent.ID_LOCALIZACION, &ent.CALLE, &ent.NRO, &ent.PISO, &ent.DEPTO, &ent.TELEF_LINEA, &ent.TELEF_CELULAR, &ent.COD_POSTAL, &ent.BARRIO, &ent.LOCALIDAD, &ent.PROVINCIA, &ent.PAIS)
	return ent, err
}
