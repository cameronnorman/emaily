package service

import (
	"fmt"
	"net/smtp"

	"github.com/labstack/gommon/log"
)

type details struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type config struct {
	Server   string `json:"server"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Request struct {
	Details      details                `json:"details"`
	Config       config                 `json:"config"`
	TemplateName string                 `json:"template_name"`
	Data         map[string]interface{} `json:"data"`
}

func SendMail(r Request) error {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	fromHeader := "From: " + r.Details.From + "!\n"
	toHeader := "To: " + r.Details.To + "!\n"
	subject := "Subject: " + r.Details.Subject + "!\n"
	auth := smtp.PlainAuth("", r.Config.Username, r.Config.Password, r.Config.Server)
	addr := fmt.Sprintf("%s:%s", r.Config.Server, r.Config.Port)
	msg := []byte(toHeader + fromHeader + subject + mime + "\n" + r.Details.Body)
	err := smtp.SendMail(addr, auth, r.Details.From, []string{r.Details.To}, msg)

	if err != nil {
		log.Error("Email failed to send to:" + r.Details.To + " - " + err.Error())
		return nil
	}

	log.Info("Email successfully sent to:" + r.Details.To)

	return nil
}
