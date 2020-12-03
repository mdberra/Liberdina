package Base

const (
	NoDataFound string = "Dato no encontrado"
)

type CustomError struct {
	Tipo    string
	Funcion string
	Err     string
}

func (e *CustomError) Error() string {
	aux := "Tipo: " + e.Tipo + " Funcion: " + e.Funcion + " Err: " + e.Err
	return aux
}

func (e *CustomError) IsNoDataFound() bool {
	if e.Tipo == NoDataFound {
		return true
	} else {
		return false
	}
}
