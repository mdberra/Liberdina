package Data

import (
	"database/sql"
	"time"
)

type Farmacia struct {
	IdFarmacia    int64     `json:"idFarmacia"`
	Nombre        string    `json:"nombre"`
	Direccion     string    `json:"direccion"`
	Telefono      string    `json:"telefono"`
	Email         string    `json:"email"`
	FechaIngreso  time.Time `json:"fechaIngreso"`
	Estado        int64     `json:"estado"`
	FechaEstado   time.Time `json:"fechaEstado"`
	EstadoDescrip string    `json:"estadoDescrip"`
}

func (n *Farmacia) CleanFarmacia() {
	n.IdFarmacia = 0
	n.Nombre = ""
	n.Direccion = ""
	n.Telefono = ""
	n.Email = ""
	n.FechaIngreso = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	n.Estado = 0
	n.FechaEstado = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	n.EstadoDescrip = ""
}
func (n *Farmacia) GetFarmacia(db *sql.DB) (Farmacia, error) {
	var ent Farmacia
	q := `select c.idFarmacia, c.nombre, c.direccion, c.telefono, c.email,
	             c.fechaIngreso, c.estado, c.fechaEstado, k.descripcion
			from Farmacia AS c    inner join  KeyValue AS k  
			     on k.entidad = "Farmacia" and k.atributo = "estado" and k.idEstado = c.estado
			where c.idFarmacia =?`
	rows, err := db.Query(q, n.IdFarmacia)
	if err != nil {
		return ent, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&ent.IdFarmacia, &ent.Nombre, &ent.Direccion, &ent.Telefono, &ent.Email,
			&ent.FechaIngreso, &ent.Estado, &ent.FechaEstado, &ent.EstadoDescrip); err != nil {
			return ent, err
		}
	}
	if n.IdFarmacia != ent.IdFarmacia {
		ent.IdFarmacia = -1
	}
	return ent, err
}
func (n *Farmacia) GetFarmacias(db *sql.DB) ([]Farmacia, error) {
	// preparar el query para ejecutarlo muchas veces,  ahorra tiempo
	q := `select c.idFarmacia, c.nombre, c.direccion, c.telefono, c.email,
				    c.fechaIngreso, c.estado, c.fechaEstado, k.descripcion
			from Farmacia AS c    inner join  KeyValue AS k  
  					on k.entidad = "Farmacia" and k.atributo = "estado" and k.idEstado = c.estado`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	e := []Farmacia{} //array
	for rows.Next() {
		var ent Farmacia
		if err := rows.Scan(&ent.IdFarmacia, &ent.Nombre, &ent.Direccion, &ent.Telefono, &ent.Email,
			&ent.FechaIngreso, &ent.Estado, &ent.FechaEstado, &ent.EstadoDescrip); err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
func (n *Farmacia) CreateFarmacia(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("INSERT Farmacia SET NOMBRE=?, DIRECCION=?, TELEFONO=?, EMAIL=?, fechaIngreso=?, estado=?, fechaEstado=?")
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
		n.IdFarmacia = id
	}

	err = tx.Commit()
	return err
}
func (n *Farmacia) UpdateFarmacia(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	stmt, err := db.Prepare("update Farmacia set nombre=?, direccion=?, telefono=?, email=?, fechaIngreso=?, estado=?, fechaEstado=? where idFarmacia=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.Nombre, n.Direccion, n.Telefono, n.Email, n.FechaIngreso, n.Estado, n.FechaEstado, n.IdFarmacia)
	if err != nil {
		return 0, err
	}
	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}
func (n *Farmacia) DeleteFarmacia(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("delete from Farmacia where idFarmacia=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.IdFarmacia)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}