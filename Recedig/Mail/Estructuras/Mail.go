package Estructuras

import (
	"fmt"
	"strings"
)

type Request struct {
	From    string
	To      []string
	Subject string
	Body    string
}

func (m *Mail) Clean() {
	m.SenderId = ""
	m.ToIds = []string{""}
	m.Subject = ""
	m.Body = ""
}

func (m *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", m.SenderId)
	if len(m.ToIds) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(m.ToIds, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", m.Subject)
	message += "\r\n" + m.Body

	return message
}
