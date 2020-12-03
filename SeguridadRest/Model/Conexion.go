package Model

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

/*
MySQL               PostgreSQL            Oracle
=====               ==========            ======
WHERE col = ?       WHERE col = $1        WHERE col = :col
VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
*/
const (
	Local         string = "LOCAL"
	ComputeEngine string = "COMPUTE ENGINE"
)

var (
	err     error
	entorno string = Local //
)

func (r *Conexion) GetEsquemaLiberdina() string {
	return "liberdina"
}
func (r *Conexion) GetEsquemaSeguridad() string {
	return "seguridad"
}

type Conexion struct {
	Db       *sql.DB
	Ip       string
	User     string
	Password string
	DbName   string
}

func (r *Conexion) CleanConexion() {
	r.Db = nil
	r.Ip = ""
	r.User = ""
	r.Password = ""
	r.DbName = ""
}

func (r *Conexion) GetDataSourceName() string {
	dataSourceName := r.User + ":" + r.Password + "@tcp(" + r.Ip + ":3306)/" + r.DbName + "?parseTime=true"
	return dataSourceName
}

// conectar a la DB
func (r *Conexion) Open(esquema string) error {
	var err error
	if r.Db == nil {
		if err = r.getEntorno(); err != nil {
			return err
		}
	}
	driverName := "mysql"
	r.DbName = esquema
	r.Db, err = sql.Open(driverName, r.GetDataSourceName())
	if err != nil {
		log.Println("No se puede conectar a la DabaBase" + r.GetDataSourceName())
		panic(err)
	}
	return err
}

// cerrar la conexion
func (r *Conexion) Close() {
	r.Db.Close()
}
func (r *Conexion) getEntorno() error {
	switch os.Getenv("Entorno") {
	case os.Getenv("EntornoLocal"):
		r.Ip = os.Getenv("IpLocal")
		r.User = os.Getenv("UserLocal")
		r.Password = os.Getenv("PasswordLocal")

	case os.Getenv("EntornoComputeEngine"):
		r.Ip = os.Getenv("IpComputeEngine")
		r.User = os.Getenv("UserComputeEngine")
		r.Password = os.Getenv("PasswordComputeEngine")

	default:
		err := errors.New("Entorno no definido")
		return err
	}
	return nil
}
func (r *Conexion) ConectToDB() error {
	err = nil
	log.Println("Conectando a la DB...")
	r.CleanConexion()
	if err = r.getEntorno(); err != nil {
		return err
	}
	log.Println("DataSourceName : " + r.GetDataSourceName())

	if err = r.Open(r.GetEsquemaLiberdina()); err != nil {
		log.Println("error Open")
		return err
	}
	err = r.Db.Ping()
	if err != nil {
		log.Println("No se pudo conectar")
		log.Println(err)
		return err
	}
	log.Println("Conexion exitosa")
	return err
}
