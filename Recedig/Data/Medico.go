package Data

import (
	"database/sql"
	"time"
)

type Medico struct {
	IdMedico      int64     `json:"idMedico"`
	Nombre        string    `json:"nombre"`
	Apellido      string    `json:"apellido"`
	Email         string    `json:"email"`
	Telefono      string    `json:"telefono"`
	Dni           int64     `json:"dni"`
	Matricula     string    `json:"matricula"`
	IdImagen      int64     `json:"idImagen"`
	FechaIngreso  time.Time `json:"fechaIngreso"`
	Estado        int64     `json:"estado"`
	FechaEstado   time.Time `json:"fechaEstado"`
	EstadoDescrip string    `json:"estadoDescrip"`
}

func (n *Medico) CleanMedico() {
	n.IdMedico = 0
	n.Nombre = ""
	n.Apellido = ""
	n.Email = ""
	n.Telefono = ""
	n.Dni = 0
	n.Matricula = ""
	n.IdImagen = 0
	n.FechaIngreso = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	n.Estado = 0
	n.FechaEstado = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	n.EstadoDescrip = ""
}
func (n *Medico) GetMedico(db *sql.DB) (Medico, error) {
	var ent Medico
	q := `select c.idMedico, c.nombre, c.apellido, c.email, c.telefono, c.dni, c.matricula,
	            c.idImagen, c.fechaIngreso, c.estado, c.fechaEstado, k.descripcion
			from Medico AS c    inner join  KeyValue AS k  
			     on k.entidad = "Medico" and k.atributo = "estado" and k.idEstado = c.estado
			where c.idMedico =?`
	rows, err := db.Query(q, n.IdMedico)
	if err != nil {
		return ent, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&ent.IdMedico, &ent.Nombre, &ent.Apellido, &ent.Email, &ent.Telefono, &ent.Dni,
			&ent.Matricula, &ent.IdImagen, &ent.FechaIngreso, &ent.Estado, &ent.FechaEstado, &ent.EstadoDescrip); err != nil {
			return ent, err
		}
	}
	if n.IdMedico != ent.IdMedico {
		ent.IdMedico = -1
	}
	return ent, err
}
func (n *Medico) GetMedicosDni(db *sql.DB) ([]Medico, error) {
	e := []Medico{} //array
	q := `select c.idMedico, c.nombre, c.apellido, c.email, c.telefono, c.dni, c.matricula,
	             c.idImagen, c.fechaIngreso, c.estado, c.fechaEstado, k.descripcion
			from Medico AS c    inner join  KeyValue AS k  
			      on k.entidad = "Medico" and k.atributo = "estado" and k.idEstado = c.estado
			where c.dni =?`
	rows, err := db.Query(q, n.Dni)
	if err != nil {
		return e, err
	}
	defer rows.Close()

	for rows.Next() {
		var ent Medico
		if err := rows.Scan(&ent.IdMedico, &ent.Nombre, &ent.Apellido, &ent.Email, &ent.Telefono, &ent.Dni,
			&ent.Matricula, &ent.IdImagen, &ent.FechaIngreso, &ent.Estado, &ent.FechaEstado, &ent.EstadoDescrip); err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
func (n *Medico) GetMedicos(db *sql.DB) ([]Medico, error) {
	// preparar el query para ejecutarlo muchas veces,  ahorra tiempo
	q := `select c.idMedico, c.nombre, c.apellido, c.email, c.telefono, c.dni, c.matricula,
				    c.idImagen, c.fechaIngreso, c.estado, c.fechaEstado, k.descripcion
			from Medico AS c    inner join  KeyValue AS k  
  					on k.entidad = "Medico" and k.atributo = "estado" and k.idEstado = c.estado`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	e := []Medico{} //array
	for rows.Next() {
		var ent Medico
		if err := rows.Scan(&ent.IdMedico, &ent.Nombre, &ent.Apellido, &ent.Email, &ent.Telefono, &ent.Dni,
			&ent.Matricula, &ent.IdImagen, &ent.FechaIngreso, &ent.Estado, &ent.FechaEstado, &ent.EstadoDescrip); err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
func (n *Medico) CreateMedico(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("INSERT Medico SET NOMBRE=?, APELLIDO=?, EMAIL=?, TELEFONO=?, DNI=?, MATRICULA=?, fechaIngreso=?, estado=?, fechaEstado=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.Nombre, n.Apellido, n.Email, n.Telefono, n.Dni, n.Matricula, n.FechaIngreso, n.Estado, n.FechaEstado)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err == nil {
		n.IdMedico = id
		n.IdImagen = id
		stmt, err = db.Prepare("update Medico set idImagen=? where idMedico=?")
		if err != nil {
			return err
		}
		res, err = stmt.Exec(n.IdImagen, n.IdMedico)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	return err
}
func (n *Medico) UpdateMedico(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	stmt, err := db.Prepare("update Medico set nombre=?, apellido=?, email=?, telefono=?, dni=?, matricula, fechaIngreso=?, estado=?, fechaEstado=? where idMedico=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.Nombre, n.Apellido, n.Email, n.Telefono, n.Dni, n.Matricula, n.FechaIngreso, n.Estado, n.FechaEstado, n.IdMedico)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}
func (n *Medico) DeleteMedico(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("delete from Medico where dni=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.Dni)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}