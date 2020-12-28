package connect

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/Liberdina/protobuffers/zuvap/pagopb"
	"github.com/Liberdina/zuvap-server/server/param"
)

// MySQL               PostgreSQL            Oracle
// =====               ==========            ======
// WHERE col = ?       WHERE col = $1        WHERE col = :col
// VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)**/

var (
	err error
)

type Conexion struct {
	Db             *sql.DB
	param          param.Parameters
	DataSourceName string
}

func (c *Conexion) ConectToDB(param *param.Parameters) error {
	c.param = *param
	err = nil
	log.Println("Conectando a la DB...")

	if err = c.Open(); err != nil {
		log.Println("error Open")
		return err
	}
	err = c.Db.Ping()
	if err != nil {
		log.Println("No se pudo conectar")
		log.Println(err)
		return err
	}
	log.Println("Conexion exitosa")
	return err
}

// conectar a la DB
func (c *Conexion) Open() error {
	err = nil
	switch c.param.DbDriver {
	case "mysql":
		c.DataSourceName = c.param.DbUser + ":" + c.param.DbPassword + "@tcp(" + c.param.DbIP + ":3306)/" + c.param.DbName + "?parseTime=true"

	case "postgres":
		c.DataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.param.DbIP, "5432", c.param.DbUser, c.param.DbPassword, c.param.DbName)

	default:
		err = errors.New("Driver no definido")
		return err
	}
	c.Db, err = sql.Open(c.param.DbDriver, c.DataSourceName)
	if err != nil {
		log.Println("No se puede conectar a la DabaBase" + c.DataSourceName)
		panic(err)
	}
	return err
}

// cerrar la Conexion
func (c *Conexion) Close() {
	c.Db.Close()
}
func (c *Conexion) LoguearInsert(req *pagopb.PagoRequest) (int, error) {
	var idPago int
	err := c.Db.QueryRow(`CALL public.newpayment($1, $2, $3, $4, $5, $6, $7, $8)`,
		1, req.GetPago().GetIdUsuario(), 1, req.GetPago().GetPaymentType(),
		req.GetPago().GetCardNumber()[0:4], req.GetPago().GetCardHolderName(),
		req.GetPago().GetCurrency(), req.GetPago().GetAmount()).Scan(&idPago)
	//	req.GetPago().GetCardExpirationDateMonth(),
	//	req.GetPago().GetCardExpirationDateYear(),
	//	req.GetPago().GetCardSecurityCode(),
	fmt.Println(err)
	fmt.Println(idPago)
	return idPago, err
}
func (c *Conexion) LoguearUpdate(idPago int, resp *pagopb.PagoResponse) error {
	_, err := c.Db.Exec(`CALL public.updatepayment($1, $2, $3, $4, $5, $6, $7)`,
		resp.GetPago().PaymentId,
		resp.GetPago().InstallmentPlanDetailId,
		resp.GetPago().ErrorCode,
		resp.GetPago().ErrorMessage,
		resp.GetPago().PagoId,
		resp.GetPago().PagoDate,
		idPago)
	fmt.Println(err)
	return err
}
