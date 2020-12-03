package data

import (
	"database/sql"
	"time"
)

type Cliente struct {
	ID_CLIENTE        int64     `json:"idCliente"`
	TIPO_DOC          string    `json:"tipoDoc"`
	NRO_DOC           int64     `json:"nroDoc"`
	CBU               string    `json:"cbu"`
	CA_SUCURSAL       string    `json:"caSucursal"`
	CA_NRO            string    `json:"caNro"`
	NOMBRE            string    `json:"nombre"`
	APELLIDO          string    `json:"apellido"`
	F_INGRESO         time.Time `json:"fechaIngreso"`
	ID_LOC_PARTICULAR int64     `json:"idLocParticular"`
	ID_LOC_LABORAL    int64     `json:"idLocLaboral"`
	ID_LOC_INFORMADO  int64     `json:"idLocInformado"`
	ID_DELEGACION     int64     `json:"idDelegacion"`
	ESTADO            int64     `json:"estado"`
	F_ESTADO          time.Time `json:"fechaEstado"`
	ESTADO_ANTERIOR   int64     `json:"estadoAnterior"`
	COMENTARIOS       string    `json:"comentarios"`
}

func (n *Cliente) CleanCliente() {
	n.ID_CLIENTE = 0
	n.TIPO_DOC = ""
	n.NRO_DOC = 0
	n.CBU = ""
	n.CA_SUCURSAL = ""
	n.CA_NRO = ""
	n.NOMBRE = ""
	n.APELLIDO = ""
	n.F_INGRESO = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	n.ID_LOC_PARTICULAR = 0
	n.ID_LOC_LABORAL = 0
	n.ID_LOC_INFORMADO = 0
	n.ID_DELEGACION = 0
	n.ESTADO = 0
	n.F_ESTADO = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	n.ESTADO_ANTERIOR = 0
	n.COMENTARIOS = ""
}

func (n *Cliente) GetCliente(db *sql.DB) (Cliente, error) {
	var ent Cliente
	q := `select * from CLIENTE where nro_doc=?`
	rows, err := db.Query(q, n.NRO_DOC)
	if err != nil {
		return ent, err
	}
	defer rows.Close()

	for rows.Next() {
		//		log.Println("rows ", rows)
		if err := rows.Scan(&ent.ID_CLIENTE, &ent.TIPO_DOC, &ent.NRO_DOC, &ent.CBU, &ent.CA_SUCURSAL, &ent.CA_NRO, &ent.NOMBRE, &ent.APELLIDO, &ent.F_INGRESO, &ent.ID_LOC_PARTICULAR, &ent.ID_LOC_LABORAL, &ent.ID_LOC_INFORMADO, &ent.ID_DELEGACION, &ent.ESTADO, &ent.F_ESTADO, &ent.ESTADO_ANTERIOR, &ent.COMENTARIOS); err != nil {
			return ent, err
		}
	}
	return ent, nil
}

func (n *Cliente) UpdateCliente(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("update CLIENTE set TIPO_DOC=?, NRO_DOC=?, CBU=?, CA_SUCURSAL=?, CA_NRO=?, NOMBRE=?, APELLIDO=?, F_INGRESO=?, ID_LOC_PARTICULAR=?, ID_LOC_LABORAL=?, iD_LOC_INFORMADO=?, ID_DELEGACION=?, ESTADO=?, F_ESTADO=?, ESTADO_ANTERIOR=?, COMENTARIOS=? where ID_CLIENTE=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.TIPO_DOC, n.NRO_DOC, n.CBU, n.CA_SUCURSAL, n.CA_NRO, n.NOMBRE, n.APELLIDO, n.F_INGRESO, n.ID_LOC_PARTICULAR, n.ID_LOC_LABORAL, n.ID_LOC_INFORMADO, n.ID_DELEGACION, n.ESTADO, n.F_ESTADO, n.ESTADO_ANTERIOR, n.COMENTARIOS, n.ID_CLIENTE)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}

func (n *Cliente) DeleteCliente(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("delete from CLIENTE where ID_CLIENTE=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.ID_CLIENTE)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}

func (n *Cliente) CreateCliente(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("INSERT CLIENTE SET TIPO_DOC=?, NRO_DOC=?, CBU=?, CA_SUCURSAL=?, CA_NRO=?, NOMBRE=?, APELLIDO=?, F_INGRESO=?, ID_LOC_PARTICULAR=?, ID_LOC_LABORAL=?, iD_LOC_INFORMADO=?, ID_DELEGACION=?, ESTADO=?, F_ESTADO=?, ESTADO_ANTERIOR=?, COMENTARIOS=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.TIPO_DOC, n.NRO_DOC, n.CBU, n.CA_SUCURSAL, n.CA_NRO, n.NOMBRE, n.APELLIDO, n.F_INGRESO, n.ID_LOC_PARTICULAR, n.ID_LOC_LABORAL, n.ID_LOC_INFORMADO, n.ID_DELEGACION, n.ESTADO, n.F_ESTADO, n.ESTADO_ANTERIOR, n.COMENTARIOS)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err == nil {
		n.ID_CLIENTE = id
	}

	err = tx.Commit()
	return err
}

//func GetClientes(db *sql.DB, start, count int) ([]Cliente, error) {
func (n *Cliente) GetClientes(db *sql.DB) ([]Cliente, error) {
	// preparar el query para ejecutarlo muchas veces,  ahorra tiempo
	stmt, err := db.Prepare("select ID_CLIENTE, TIPO_DOC, NRO_DOC, CBU, CA_SUCURSAL, CA_NRO, NOMBRE, APELLIDO from CLIENTE")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	e := []Cliente{} //array
	for rows.Next() {
		var ent Cliente
		if err := rows.Scan(&ent.ID_CLIENTE, &ent.TIPO_DOC, &ent.NRO_DOC, &ent.CBU, &ent.CA_SUCURSAL, &ent.CA_NRO, &ent.NOMBRE, &ent.APELLIDO); err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
