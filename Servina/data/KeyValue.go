package data

import "database/sql"

type KeyValue struct {
	IdKeyValue  int64  `json:"idKeyValue"`
	Entidad     string `json:"entidad"`
	Atributo    string `json:"atributo"`
	IdEstado    int64  `json:"idEstado"`
	Descripcion string `json:"descripcion"`
}

func (n *KeyValue) CleanKeyValue() {
	n.IdKeyValue = 0
	n.Entidad = ""
	n.Atributo = ""
	n.IdEstado = 0
	n.Descripcion = ""
}
func (n *KeyValue) GetKeyValue(db *sql.DB) ([]KeyValue, error) {
	q := `select idKeyValue, entidad, atributo, idEstado, descripcion from KeyValue where entidad=?`
	rows, err := db.Query(q, n.Entidad)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	e := []KeyValue{} //array
	for rows.Next() {
		var ent KeyValue
		if err := rows.Scan(&ent.IdKeyValue, &ent.Entidad, &ent.Atributo, &ent.IdEstado, &ent.Descripcion); err != nil {
			return nil, err
		}
		e = append(e, ent)
	}
	return e, nil
}
