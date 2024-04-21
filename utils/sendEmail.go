package utils

import (
	"bytes"
	"html/template"
	"log"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/config"
	"gopkg.in/gomail.v2"
)

var Email = config.Environments().Email

type DataEmail struct {
	Email      string
	First_Name string
	Username   string
	Password   string
}

type SendEmail interface {
	EmployeeRegistered(email, first_name, username, password string)
}

type sendEmail struct{}

func NewSendEmail() SendEmail {
	return &sendEmail{}
}

func (s *sendEmail) EmployeeRegistered(email, first_name, username, password string) {

	dataEmail := DataEmail{
		Email:      email,
		First_Name: first_name,
		Username:   username,
		Password:   password,
	}

	t, err := template.ParseFiles("./templates/welcome-new-employee.html")
	if err != nil {
		log.Println(err)
	}

	body := new(bytes.Buffer)
	if err := t.Execute(body, dataEmail); err != nil {
		log.Println(err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "nljstore.ecommerce@gmail.com")
	m.SetHeader("To", dataEmail.Email)
	m.SetHeader("Subject", "¡Bienvenido al equipo NLJStore!")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, Email.Email, Email.Password)

	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
	}

	log.Println("Email enviado con exito.!!")

}
