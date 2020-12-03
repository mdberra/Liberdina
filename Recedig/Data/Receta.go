package Data

import (
	"database/sql"
	"time"
)

type Receta struct {
	IdReceta      int64     `json:"idReceta"`
	FechaCreacion time.Time `json:"fechaCreacion"`
	IdMedico      int64     `json:"idMedico"`
	Estado        int64     `json:"estado"`
	FechaEstado   time.Time `json:"fechaEstado"`
	EstadoDescrip string    `json:"estadoDescrip"`
}
type RecetaItem struct {
	IdRecetaItem  int64 `json:"idRecetaItem"`
	IdReceta      int64 `json:"idReceta"`
	IdMedicamento int64 `json:"idMedicamento"`
}
type RecetaFarmacia struct {
	IdReceta      int64     `json:"idReceta"`
	IdFarmacia    int64     `json:"idFarmacia"`
	Estado        int64     `json:"estado"`
	FechaEstado   time.Time `json:"fechaEstado"`
	EstadoDescrip string    `json:"estadoDescrip"`
}

func (n *Receta) CleanReceta() {
	n.IdReceta = 0
	n.FechaCreacion = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	n.IdMedico = 0
	n.Estado = 0
	n.FechaEstado = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	n.EstadoDescrip = ""
}
func (n *Receta) CleanRecetaItems(recetaItems []RecetaItem) {
	for i := range recetaItems {
		recetaItems[i].IdRecetaItem = 0
		recetaItems[i].IdReceta = 0
		recetaItems[i].IdMedicamento = 0
	}
}
func (n *Receta) GetReceta(db *sql.DB) (Receta, error) {
	var ent Receta
	q := `select c.idReceta, c.fechaCreacion, c.idMedico, c.estado, c.fechaEstado, k.descripcion
			from Receta AS c    inner join  KeyValue AS k  
			     on k.entidad = "Receta" and k.atributo = "estado" and k.idEstado = c.estado
			where c.idReceta =?`
	rows, err := db.Query(q, n.IdReceta)
	if err != nil {
		return ent, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&ent.IdReceta, &ent.FechaCreacion, &ent.IdMedico,
			&ent.Estado, &ent.FechaEstado, &ent.EstadoDescrip); err != nil {
			return ent, err
		}
	}
	if n.IdReceta != ent.IdReceta { //sino la encuentra
		ent.IdReceta = -1
	}
	return ent, err
}
func (n *Receta) GetRecetaItem(db *sql.DB) ([]RecetaItem, error) {
	q := `select c.idRecetaItem, c.idReceta, c.idMedicamento
			from RecetaItem AS c
			where c.idReceta =?`
	rows, err := db.Query(q, n.IdReceta)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	e := []RecetaItem{} //array
	for rows.Next() {
		var ent RecetaItem
		if err := rows.Scan(&ent.IdRecetaItem, &ent.IdReceta, &ent.IdMedicamento); err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
func (n *Receta) GetRecetaFarmacia(db *sql.DB) (RecetaFarmacia, error) {
	var ent RecetaFarmacia
	q := `select c.idReceta, c.idFarmacia, c.estado, c.fechaEstado, k.descripcion
			from RecetaFarmacia AS c    inner join  KeyValue AS k  
			     on k.entidad = "Receta" and k.atributo = "estado" and k.idEstado = c.estado
			where c.idReceta =?`
	rows, err := db.Query(q, n.IdReceta)
	if err != nil {
		return ent, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&ent.IdReceta, &ent.IdFarmacia,
			&ent.Estado, &ent.FechaEstado, &ent.EstadoDescrip); err != nil {
			return ent, err
		}
	}
	if n.IdReceta != ent.IdReceta { //sino la encuentra
		ent.IdReceta = -1
	}
	return ent, err
}
func (n *Receta) GetRecetas(db *sql.DB) ([]Receta, error) {
	// preparar el query para ejecutarlo muchas veces,  ahorra tiempo
	q := `select c.idReceta, c.fechaCreacion, c.idMedico, c.estado, c.fechaEstado, k.descripcion
			from Receta AS c    inner join  KeyValue AS k  
  					on k.entidad = "Receta" and k.atributo = "estado" and k.idEstado = c.estado`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	e := []Receta{} //array
	for rows.Next() {
		var ent Receta
		if err := rows.Scan(&ent.IdReceta, &ent.FechaCreacion, &ent.IdMedico,
			&ent.Estado, &ent.FechaEstado, &ent.EstadoDescrip); err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
func (n *Receta) CreateReceta(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("INSERT Receta SET fechaCreacion=?, idMedico=?, estado=?, fechaEstado=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.FechaCreacion, n.IdMedico, n.Estado, n.FechaEstado)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err == nil {
		n.IdReceta = id
	}

	err = tx.Commit()
	return err
}
func (n *Receta) CreateRecetaItem(db *sql.DB, recetaItem []RecetaItem) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("INSERT RecetaItem SET idReceta=?, idMedicamento=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for i := range recetaItem {
		ent := recetaItem[i]
		res, err := stmt.Exec(ent.IdReceta, ent.IdMedicamento)
		if err != nil {
			break
		}
		id, err := res.LastInsertId()
		if err != nil {
			break
		}
		recetaItem[i].IdRecetaItem = id
	}
	if err != nil {
		return err
	}
	err = tx.Commit()
	return err
}
func (n *Receta) UpdateReceta(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	stmt, err := db.Prepare("update Receta set fechaCreacion=?, idMedico=?, estado=?, fechaEstado=? where idReceta=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.FechaCreacion, n.IdMedico, n.Estado, n.FechaEstado, n.IdReceta)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}
func (n *Receta) DeleteReceta(db *sql.DB) (int64, error) {
	var filasAfectadas int64

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("delete from Receta where idReceta=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(n.IdReceta)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()

	filasAfectadas, err = res.RowsAffected()

	return filasAfectadas, err
}
