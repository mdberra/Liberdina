package data

import (
	"database/sql"
	"time"
)

type FindCliente struct {
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
	DELEG_DESCRIP     string    `json:"delegDescrip"`
	ESTADO            int64     `json:"idEstado"`
	EST_DESCRIP       string    `json:"estadoDescrip"`
	F_ESTADO          time.Time `json:"fechaEstado"`
	ESTADO_ANTERIOR   int64     `json:"estadoAnterior"`
	COMENTARIOS       string    `json:"comentarios"`
}

func (n *FindCliente) CleanFindCliente() {
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
	n.DELEG_DESCRIP = ""
	n.ESTADO = 0
	n.EST_DESCRIP = ""
	n.F_ESTADO = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	n.ESTADO_ANTERIOR = 0
	n.COMENTARIOS = ""
}
func (n *FindCliente) GetFindCliente(db *sql.DB) ([]FindCliente, error) {
	q := `
	select c.id_cliente as idCliente, c.tipo_doc as tipoDoc, c.nro_doc as nroDoc,
		c.cbu as cbu, c.ca_sucursal as caSucursal, c.ca_nro as caNro, c.nombre as nombre,
		c.apellido as apellido, c.f_ingreso as fIngreso, c.id_loc_particular as idLocParticular,
		c.id_loc_laboral as idLocLaboral, c.id_loc_informado as idLocInformado,
		c.id_delegacion as idDelegacion, d.descripcion as descripcion, c.estado as estado,
		case c.estado
			when 0 then "Activo"
			when 1 then "Embargo"
			when 2 then "Finalizado"
			when 3 then "EnviadoJuan"
			when 4 then "Incobrable"
			else "Sin Estado"
		end as estDescrip,
		c.f_estado as fEstado, c.estado_anterior as estadoAnterior, c.comentarios as comentarios
	from CLIENTE AS c  inner join DELEGACION AS d  on c.id_delegacion = d.id_delegacion
	order by c.estado, c.id_cliente;
	`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	e := []FindCliente{} //array
	for rows.Next() {
		var ent FindCliente
		if err := rows.Scan(&ent.ID_CLIENTE, &ent.TIPO_DOC, &ent.NRO_DOC, &ent.CBU,
			&ent.CA_SUCURSAL, &ent.CA_NRO, &ent.NOMBRE, &ent.APELLIDO, &ent.F_INGRESO,
			&ent.ID_LOC_PARTICULAR, &ent.ID_LOC_LABORAL, &ent.ID_LOC_INFORMADO, &ent.ID_DELEGACION,
			&ent.DELEG_DESCRIP, &ent.ESTADO, &ent.EST_DESCRIP, &ent.F_ESTADO,
			&ent.ESTADO_ANTERIOR, &ent.COMENTARIOS); err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
