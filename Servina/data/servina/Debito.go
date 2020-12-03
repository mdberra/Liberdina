package data

import (
	"database/sql"
	"time"
)

type Debito struct {
	ID_CLIENTE       int64     `json:"idCliente"`
	ID_SERVICIO      int64     `json:"idServicio"`
	PERIODO          string    `json:"periodo"`
	FECHA            time.Time `json:"fecha"`
	IMPORTE          float32   `json:"importe"`
	BANCO_RECAUDADOR string    `json:"bancoRecaudador"`
	ESTADO           string    `json:"estado"`
	CODIGO_RECHAZO   string    `json:"codigoRechazo"`
}

func (n *Debito) CleanDebito() {
	n.ID_CLIENTE = 0
	n.ID_SERVICIO = 0
	n.PERIODO = ""
	n.FECHA = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	n.IMPORTE = 0
	n.BANCO_RECAUDADOR = ""
	n.ESTADO = ""
	n.CODIGO_RECHAZO = ""
}

/**
if(m.fecha < 20180700, "Banco Provincia",
case d.id_banco
	when 22 then "Banco Itau"
	when 23 then "Banco Bica"
	else "Banco Provincia"
end
)
*/
func (n *Debito) GetDebito(db *sql.DB) ([]Debito, error) {
	q := `
	select m.id_cliente, m.id_servicio,
		DATE_FORMAT(m.fecha, '%m/%Y') as periodo,
		m.fecha as fecha,
		m.importe as importe,
 		m.descripcion as bancoRecaudador,
		case e.id_estado
			when 1 then "1"
			when 2 then "Disparado"
			when 3 then "Acreditado"
			when 5 then "Reversion"
			when 6 then "Devuelto"
			else "Rechazado"
		end as estado,
		case e.id_estado
			when 1 then " "
			when 2 then " "
			when 3 then " "
			when 5 then " "
			when 6 then " "
			else e.cod_rechazo
		end
	from CLIENTE AS c    inner join DELEGACION AS d   on c.id_delegacion = d.id_delegacion,
		MOVIMIENTO AS m  inner join Estado_Mov AS e   on m.id_estado = e.id_estado
	where m.id_cliente = c.id_cliente and
		c.id_cliente = ?
	order by m.id_servicio, m.fecha desc;
	`
	rows, err := db.Query(q, n.ID_CLIENTE)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	e := []Debito{} //array
	for rows.Next() {
		var ent Debito
		if err := rows.Scan(&ent.ID_CLIENTE, &ent.ID_SERVICIO, &ent.PERIODO, &ent.FECHA, &ent.IMPORTE, &ent.BANCO_RECAUDADOR, &ent.ESTADO, &ent.CODIGO_RECHAZO); err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
