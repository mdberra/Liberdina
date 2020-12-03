package data

import (
	"database/sql"
)

//----------------------------------------------------------------------------------------------
type General struct {
	DELEGACION       string  `json:"DELEGACION"`
	BANCO_RECAUDADOR string  `json:"BANCO_RECAUDADOR"`
	ESTADO           string  `json:"ESTADO"`
	CODIGO_RECHAZO   string  `json:"CODIGO_RECHAZO"`
	PERIODO          string  `json:"PERIODO"`
	IMPORTE          float32 `json:"IMPORTE"`
}

func (n *General) CleanGeneral() {
	n.DELEGACION = ""
	n.BANCO_RECAUDADOR = ""
	n.ESTADO = ""
	n.CODIGO_RECHAZO = ""
	n.PERIODO = ""
	n.IMPORTE = 0
}

func (n *General) GetGeneral(db *sql.DB) ([]General, error) {
	// preparar el query para ejecutarlo muchas veces,  ahorra tiempo
	q := `
	select d.descripcion as DELEGACION,
		if(m.fecha < 20180700, "Banco Provincia",
			case d.id_banco
				when 22 then "Banco Itau"
				when 23 then "Banco Bica"
				else "Banco Provincia"
			end
		) as BANCO_RECAUDADOR,
		case e.id_estado
			when 1 then "1"
			when 2 then "Disparado"
			when 3 then "Acreditado"
			when 5 then "Reversion"
			when 6 then "Devuelto"
			else "Rechazado"
		end as ESTADO,
		case e.id_estado
			when 1 then " "
			when 2 then " "
			when 3 then " "
			when 5 then " "
			when 6 then " "
			else e.cod_rechazo
		end as CODIGO_RECHAZO,
		DATE_FORMAT(m.fecha, '%m/%Y') as PERIODO,
		sum(m.importe) as IMPORTE
	from CLIENTE AS c     inner join DELEGACION AS d   on c.id_delegacion = d.id_delegacion,
		MOVIMIENTO AS m  inner join Estado_Mov AS e   on m.id_estado = e.id_estado
	where m.id_cliente = c.id_cliente and
		fecha > 20180300
	group by DELEGACION, BANCO_RECAUDADOR, ESTADO, CODIGO_RECHAZO, PERIODO
	`

	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	e := []General{} //array
	for rows.Next() {
		var ent General
		if err := rows.Scan(&ent.DELEGACION, &ent.BANCO_RECAUDADOR, &ent.ESTADO, &ent.CODIGO_RECHAZO, &ent.PERIODO, &ent.IMPORTE); err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}

//----------------------------------------------------------------------------------------------
type General1 struct {
	DELEGACION       string  `json:"DELEGACION"`
	BANCO_RECAUDADOR string  `json:"BANCO_RECAUDADOR"`
	ESTADO           string  `json:"ESTADO"`
	CODIGO_RECHAZO   string  `json:"CODIGO_RECHAZO"`
	PERIODO          string  `json:"PERIODO"`
	IMPORTE          float32 `json:"IMPORTE"`
}

func (n *General1) CleanGeneral1() {
	n.DELEGACION = ""
	n.BANCO_RECAUDADOR = ""
	n.ESTADO = ""
	n.CODIGO_RECHAZO = ""
	n.PERIODO = ""
	n.IMPORTE = 0
}

func (n *General1) GetGeneral1(db *sql.DB) ([]General1, error) {
	// preparar el query para ejecutarlo muchas veces,  ahorra tiempo
	q := `
	select d.descripcion as DELEGACION,
		if(m.fecha < 20180700, "Banco Provincia",
			case d.id_banco
				when 22 then "Banco Itau"
				when 23 then "Banco Bica"
				else "Banco Provincia"
			end
		) as BANCO_RECAUDADOR,
		case e.id_estado
			when 1 then "1"
			when 2 then "Disparado"
			when 3 then "Acreditado"
			when 5 then "Reversion"
			when 6 then "Devuelto"
			else "Rechazado"
		end as ESTADO,
		case e.id_estado
			when 1 then " "
			when 2 then " "
			when 3 then " "
			when 5 then " "
			when 6 then " "
			else e.cod_rechazo
		end as CODIGO_RECHAZO,
		DATE_FORMAT(m.fecha, '%m/%Y') as PERIODO,
		sum(m.importe) as IMPORTE
	from CLIENTE AS c     inner join DELEGACION AS d   on c.id_delegacion = d.id_delegacion,
		MOVIMIENTO AS m  inner join Estado_Mov AS e   on m.id_estado = e.id_estado
	where m.id_cliente = c.id_cliente and
		fecha > 20180300
	group by DELEGACION, BANCO_RECAUDADOR, ESTADO, CODIGO_RECHAZO, PERIODO
	`

	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	e := []General1{} //array
	for rows.Next() {
		var ent General1
		if err := rows.Scan(&ent.DELEGACION, &ent.BANCO_RECAUDADOR, &ent.ESTADO, &ent.CODIGO_RECHAZO, &ent.PERIODO, &ent.IMPORTE); err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
