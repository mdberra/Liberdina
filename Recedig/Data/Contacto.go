package Data

import (
	"database/sql"
	"time"
)

type Contacto struct {
	IdContacto    int64     `json:"idContacto"`
	Nombre        string    `json:"nombre"`
	Apellido      string    `json:"apellido"`
	Email         string    `json:"email"`
	Telefono      string    `json:"telefono"`
	Dni           int64     `json:"dni"`
	Monto         float32   `json:"monto"`
	Plazo         int64     `json:"plazo"`
	Mensaje       string    `json:"mensaje"`
	Cbu           string    `json:"cbu"`
	IdImagen      int64     `json:"idImagen"`
	FechaIngreso  time.Time `json:"fechaIngreso"`
	Estado        int64     `json:"estado"`
	FechaEstado   time.Time `json:"fechaEstado"`
	EstadoDescrip string    `json:"estadoDescrip"`
}

func (n *Contacto) CleanContacto() {
	n.IdContacto = 0
	n.Nombre = ""
	n.Apellido = ""
	n.Email = ""
	n.Telefono = ""
	n.Dni = 0
	n.Monto = 0.0
	n.Plazo = 0
	n.Mensaje = ""
	n.Cbu = ""
	n.IdImagen = 0
	n.FechaIngreso = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	n.Estado = 0
	n.FechaEstado = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	n.EstadoDescrip = ""
}
func (n *Contacto) GetContacto(db *sql.DB) (Contacto, error) {
	var ent Contacto
	q := `select c.idContacto, c.nombre, c.apellido, c.email, c.telefono, c.dni, c.monto, c.plazo,
	             c.mensaje, c.cbu, c.idImagen, c.fechaIngreso, c.estado, c.fechaEstado, k.descripcion
			from Contacto AS c    inner join  KeyValue AS k  
			      on k.entidad = "contacto" and k.atributo = "estado" and k.idEstado = c.estado
			where c.idContacto =?`
	rows, err := db.Query(q, n.IdContacto)
	if err != nil {
		return ent, err
	}
	defer rows.Close()

	for rows.Next() {
		//		log.Println("rows ", rows)
		if err := rows.Scan(&ent.IdContacto, &ent.Nombre, &ent.Apellido, &ent.Email, &ent.Telefono, &ent.Dni, &ent.Monto,
			&ent.Plazo, &ent.Mensaje, &ent.Cbu, &ent.IdImagen, &ent.FechaIngreso, &ent.Estado, &ent.FechaEstado, &ent.EstadoDescrip); err != nil {
			return ent, err
		}
	}
	if n.IdContacto != ent.IdContacto {
		ent.IdContacto = -1
	}
	return ent, err
}
func (n *Contacto) GetContactosDni(db *sql.DB) ([]Contacto, error) {
	e := []Contacto{} //array
	q := `select c.idContacto, c.nombre, c.apellido, c.email, c.telefono, c.dni, c.monto, c.plazo,
	             c.mensaje, c.cbu, c.idImagen, c.fechaIngreso, c.estado, c.fechaEstado, k.descripcion
			from Contacto AS c    inner join  KeyValue AS k  
			      on k.entidad = "contacto" and k.atributo = "estado" and k.idEstado = c.estado
			where c.dni =?`
	rows, err := db.Query(q, n.Dni)
	if err != nil {
		return e, err
	}
	defer rows.Close()

	for rows.Next() {
		var ent Contacto
		if err := rows.Scan(&ent.IdContacto, &ent.Nombre, &ent.Apellido, &ent.Email, &ent.Telefono, &ent.Dni, &ent.Monto,
			&ent.Plazo, &ent.Mensaje, &ent.Cbu, &ent.IdImagen, &ent.FechaIngreso, &ent.Estado, &ent.FechaEstado, &ent.EstadoDescrip); err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
func (n *Contacto) GetContactos(db *sql.DB) ([]Contacto, error) {
	// preparar el query para ejecutarlo muchas veces,  ahorra tiempo
	q := `select c.idContacto, c.nombre, c.apellido, c.email, c.telefono, c.dni, c.monto, c.plazo,
				    c.mensaje, c.cbu, c.idImagen, c.fechaIngreso, c.estado, c.fechaEstado, k.descripcion
			from Contacto AS c    inner join  KeyValue AS k  
  					on k.entidad = "contacto" and k.atributo = "estado" and k.idEstado = c.estado`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	e := []Contacto{} //array
	for rows.Next() {
		var ent Contacto
		if err := rows.Scan(&ent.IdContacto, &ent.Nombre, &ent.Apellido, &ent.Email, &ent.Telefono, &ent.Dni, &ent.Monto,
			&ent.Plazo, &ent.Mensaje, &ent.Cbu, &ent.IdImagen, &ent.FechaIngreso, &ent.Estado, &ent.FechaEstado, &ent.EstadoDescrip); err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
func (n *Contacto) CreateContacto(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("INSERT Contacto SET NOMBRE=?, APELLIDO=?, EMAIL=?, TELEFONO=?, DNI=?, MONTO=?,PLAZO=?, MENSAJE=?, CBU=?, fechaIngreso=?, estado=?, fechaEstado=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.Nombre, n.Apellido, n.Email, n.Telefono, n.Dni, n.Monto, n.Plazo, n.Mensaje, n.Cbu, n.FechaIngreso, n.Estado, n.FechaEstado)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err == nil {
		n.IdContacto = id
		n.IdImagen = id
		stmt, err = db.Prepare("update Contacto set idImagen=? where idContacto=?")
		if err != nil {
			return err
		}
		res, err = stmt.Exec(n.IdImagen, n.IdContacto)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	return err
}
func (n *Contacto) UpdateContacto(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	stmt, err := db.Prepare("update Contacto set nombre=?, apellido=?, email=?, telefono=?, dni=?, monto=?,plazo=?, mensaje=?, cbu=?, fechaIngreso=?, estado=?, fechaEstado=? where idContacto=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.Nombre, n.Apellido, n.Email, n.Telefono, n.Dni, n.Monto, n.Plazo, n.Mensaje, n.Cbu, n.FechaIngreso, n.Estado, n.FechaEstado, n.IdContacto)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}
func (n *Contacto) DeleteContacto(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("delete from Contacto where dni=?")
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
