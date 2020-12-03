package data

import (
	"database/sql"
	"time"
)

type Delegaciones struct {
	ID_DELEGACION      int64     `json:"idDelegacion"`
	CODIGO             string    `json:"codigo"`
	DESCRIPCION        string    `json:"descripcion"`
	FECHA_ALTA         time.Time `json:"fechaAlta"`
	UTILIZAR           int       `json:"utilizar"`
	ESTADO_DESCRIP     string    `json:"estadoDescrip"`
	ID_BANCO           int64     `json:"idBanco"`
	BCODIGO            string    `json:"bcodigo"`
	BDESCRIPCION       string    `json:"bdescripcion"`
	CODIGO_DEBITO      int       `json:"codigoDebito"`
	DESCRIP_PRESTACION string    `json:"descripPrestacion"`
	BANCO_RECAUDADOR   string    `json:"bancoRecaudador"`
}

func (n *Delegaciones) CleanDelegaciones() {
	n.ID_DELEGACION = 0
	n.CODIGO = ""
	n.DESCRIPCION = ""
	n.FECHA_ALTA = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	n.UTILIZAR = 0
	n.ESTADO_DESCRIP = ""
	n.ID_BANCO = 0
	n.BCODIGO = ""
	n.BDESCRIPCION = ""
	n.CODIGO_DEBITO = 0
	n.DESCRIP_PRESTACION = ""
	n.BANCO_RECAUDADOR = ""
}

func (n *Delegaciones) GetDelegaciones(db *sql.DB) ([]Delegaciones, error) {
	q := `
	select	d.ID_DELEGACION, d.CODIGO, d.DESCRIPCION, d.FECHA_ALTA, d.UTILIZAR,
		case d.UTILIZAR
		when 0 then "Activo"
		when 1 then "Finalizado"
		else "Sin Estado"
		end as ESTADO_DESCRIP,
			b.ID_BANCO, b.CODIGO as BCODIGO, b.DESCRIPCION as BDESCRIPCION, b.CODIGO_DEBITO, b.DESCRIP_PRESTACION, b.BANCO_RECAUDADOR
	from DELEGACION as d, BANCO as b
	where d.ID_BANCO = b.ID_BANCO;
	`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ent Delegaciones
	e := []Delegaciones{} //array
	for rows.Next() {
		ent, err = n.parsear(rows)
		if err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}

func (n *Delegaciones) parsear(rows *sql.Rows) (Delegaciones, error) {
	var ent Delegaciones
	err := rows.Scan(&ent.ID_DELEGACION, &ent.CODIGO, &ent.DESCRIPCION, &ent.FECHA_ALTA, &ent.UTILIZAR, &ent.ESTADO_DESCRIP,
		&ent.ID_BANCO, &ent.BCODIGO, &ent.BDESCRIPCION, &ent.CODIGO_DEBITO, &ent.DESCRIP_PRESTACION, &ent.BANCO_RECAUDADOR)
	return ent, err
}
