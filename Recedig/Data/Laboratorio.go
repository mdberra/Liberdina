package Data

import (
	"database/sql"
	"time"
)

type Laboratorio struct {
	IdLaboratorio int64     `json:"idLaboratorio"`
	Nombre        string    `json:"nombre"`
	Direccion     string    `json:"direccion"`
	Telefono      string    `json:"telefono"`
	Email         string    `json:"email"`
	FechaIngreso  time.Time `json:"fechaIngreso"`
	Estado        int64     `json:"estado"`
	FechaEstado   time.Time `json:"fechaEstado"`
	EstadoDescrip string    `json:"estadoDescrip"`
}

func (n *Laboratorio) CleanLaboratorio() {
	n.IdLaboratorio = 0
	n.Nombre = ""
	n.Direccion = ""
	n.Telefono = ""
	n.Email = ""
	n.FechaIngreso = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	n.Estado = 0
	n.FechaEstado = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	n.EstadoDescrip = ""
}
func (n *Laboratorio) GetLaboratorio(db *sql.DB) (Laboratorio, error) {
	var ent Laboratorio
	q := `select c.idLaboratorio, c.nombre, c.direccion, c.telefono, c.email,
	             c.fechaIngreso, c.estado, c.fechaEstado, k.descripcion
			from Laboratorio AS c    inner join  KeyValue AS k  
			     on k.entidad = "Farmacia" and k.atributo = "estado" and k.idEstado = c.estado
			where c.idLaboratorio =?`
	rows, err := db.Query(q, n.IdLaboratorio)
	if err != nil {
		return ent, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&ent.IdLaboratorio, &ent.Nombre, &ent.Direccion, &ent.Telefono, &ent.Email,
			&ent.FechaIngreso, &ent.Estado, &ent.FechaEstado, &ent.EstadoDescrip); err != nil {
			return ent, err
		}
	}
	if n.IdLaboratorio != ent.IdLaboratorio {
		ent.IdLaboratorio = -1
	}
	return ent, err
}
func (n *Laboratorio) GetLaboratorios(db *sql.DB) ([]Laboratorio, error) {
	// preparar el query para ejecutarlo muchas veces,  ahorra tiempo
	q := `select c.idLaboratorio, c.nombre, c.direccion, c.telefono, c.email,
				    c.fechaIngreso, c.estado, c.fechaEstado, k.descripcion
			from Laboratorio AS c    inner join  KeyValue AS k  
  					on k.entidad = "Farmacia" and k.atributo = "estado" and k.idEstado = c.estado`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	e := []Laboratorio{} //array
	for rows.Next() {
		var ent Laboratorio
		if err := rows.Scan(&ent.IdLaboratorio, &ent.Nombre, &ent.Direccion, &ent.Telefono, &ent.Email,
			&ent.FechaIngreso, &ent.Estado, &ent.FechaEstado, &ent.EstadoDescrip); err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
func (n *Laboratorio) CreateLaboratorio(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("INSERT Laboratorio SET NOMBRE=?, DIRECCION=?, TELEFONO=?, EMAIL=?, fechaIngreso=?, estado=?, fechaEstado=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.Nombre, n.Direccion, n.Telefono, n.Email, n.FechaIngreso, n.Estado, n.FechaEstado)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err == nil {
		n.IdLaboratorio = id
	}

	err = tx.Commit()
	return err
}
func (n *Laboratorio) UpdateLaboratorio(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	stmt, err := db.Prepare("update Laboratorio set nombre=?, direccion=?, telefono=?, email=?, fechaIngreso=?, estado=?, fechaEstado=? where idLaboratorio=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.Nombre, n.Direccion, n.Telefono, n.Email, n.FechaIngreso, n.Estado, n.FechaEstado, n.IdLaboratorio)
	if err != nil {
		return 0, err
	}
	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}
func (n *Laboratorio) DeleteLaboratorio(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("delete from Laboratorio where idLaboratorio=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.IdLaboratorio)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}
