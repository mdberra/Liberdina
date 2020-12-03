package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"net/mail"
	"net/smtp"
)

type Dest struct {
	Nombre string
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func (rest *PepeYaRestFul) EnviarMail(nya string, email string, telefono string, dni string, monto string, plazo string) {
	from := mail.Address{
		Name:    nya,
		Address: email,
	}
	to := mail.Address{
		Name:    "Marcelo ,
		Address: "marcelo.berra@gmail.com",
	}
	subject := "Solicita Prestamo de $" + monto + " en " + plazo " meses "
	dest := Dest{
		Nombre: to.Address,
	}

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["Content-Type"] = `text/html; charset="UTF-8"`
//	headers["Body"] = mensaje

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	t, err := template.ParseFiles("Templates/mailing.html")
	checkErr(err)

	buf := new(bytes.Buffer)
	err = t.Execute(buf, dest)
	checkErr(err)

	message += buf.String()

	serverName := "smtp.gmail.com:465"
	host := "smtp.gmail.com"
	auth := smtp.PlainAuth("", "marcelo.berra@gmail.com", "Newdestino2018", host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", serverName, tlsConfig)
	checkErr(err)

	client, err := smtp.NewClient(conn, host)
	checkErr(err)

	err = client.Auth(auth)
	checkErr(err)

	client.Mail(from.Address)
	checkErr(err)

	err = client.Rcpt(to.Address)
	checkErr(err)

	w, err := client.Data()
	checkErr(err)

	_, err = w.Write([]byte(message))
	checkErr(err)

	err = w.Close()
	checkErr(err)

	client.Quit()
}