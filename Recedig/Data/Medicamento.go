package Data

import (
	"database/sql"
	"time"
)

type Medicamento struct {
	Id            int64     `json:"id"`
	Nombre        string    `json:"nombre"`
	Droga         string    `json:"droga"`
	IdLaboratorio int64     `json:"idLaboratorio"`
	FechaIngreso  time.Time `json:"fechaIngreso"`
	Estado        int64     `json:"estado"`
	FechaEstado   time.Time `json:"fechaEstado"`
	EstadoDescrip string    `json:"estadoDescrip"`
}

func (n *Medicamento) CleanMedicamento() {
	n.Id = 0
	n.Nombre = ""
	n.Droga = ""
	n.IdLaboratorio = 0
	n.FechaIngreso = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	n.Estado = 0
	n.FechaEstado = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	n.EstadoDescrip = ""
}
func (n *Medicamento) GetMedicamento(db *sql.DB) (Medicamento, error) {
	var ent Medicamento
	q := `select c.id, c.nombre, c.droga, c.idLaboratorio,
	             c.fechaIngreso, c.estado, c.fechaEstado, k.descripcion
			from Medicamento AS c    inner join  KeyValue AS k  
			     on k.entidad = "Farmacia" and k.atributo = "estado" and k.idEstado = c.estado
			where c.Id =?`
	rows, err := db.Query(q, n.Id)
	if err != nil {
		return ent, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&ent.Id, &ent.Nombre, &ent.Droga, &ent.IdLaboratorio,
			&ent.FechaIngreso, &ent.Estado, &ent.FechaEstado, &ent.EstadoDescrip); err != nil {
			return ent, err
		}
	}
	if n.Id != ent.Id {
		ent.Id = -1
	}
	return ent, err
}
func (n *Medicamento) GetMedicamentos(db *sql.DB) ([]Medicamento, error) {
	// preparar el query para ejecutarlo muchas veces,  ahorra tiempo
	q := `select c.id, c.nombre, c.droga, c.idLaboratorio,
				    c.fechaIngreso, c.estado, c.fechaEstado, k.descripcion
			from Medicamento AS c    inner join  KeyValue AS k  
  					on k.entidad = "Farmacia" and k.atributo = "estado" and k.idEstado = c.estado`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	e := []Medicamento{} //array
	for rows.Next() {
		var ent Medicamento
		if err := rows.Scan(&ent.Id, &ent.Nombre, &ent.Droga, &ent.IdLaboratorio,
			&ent.FechaIngreso, &ent.Estado, &ent.FechaEstado, &ent.EstadoDescrip); err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
func (n *Medicamento) CreateMedicamento(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("INSERT Medicamento SET NOMBRE=?, DROGA=?, idLaboratorio=? , fechaIngreso=?, estado=?, fechaEstado=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.Nombre, n.Droga, n.IdLaboratorio, n.FechaIngreso, n.Estado, n.FechaEstado)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err == nil {
		n.Id = id
	}

	err = tx.Commit()
	return err
}
func (n *Medicamento) UpdateMedicamento(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	stmt, err := db.Prepare("update Medicamento set nombre=?, droga=?, idLaboratorio=?, fechaIngreso=?, estado=?, fechaEstado=? where Id=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.Nombre, n.Droga, n.IdLaboratorio, n.FechaIngreso, n.Estado, n.FechaEstado, n.Id)
	if err != nil {
		return 0, err
	}
	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}
func (n *Medicamento) DeleteMedicamento(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("delete from Medicamento where Id=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.Id)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}
