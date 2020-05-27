package main

import (
	"fmt"
	"net/smtp"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type emailDetails struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type emailConfig struct {
	Server   string `json:"server"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type sendEmailRequest struct {
	Details emailDetails `json:"details"`
	Config  emailConfig  `json:"config"`
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.POST("/", handleSendEmailRequest)
	e.Logger.Fatal(e.Start(":8081"))
}

func handleSendEmailRequest(c echo.Context) error {
	r := sendEmailRequest{}
	if err := c.Bind(&r); err != nil {
		return err
	}
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	fromHeader := "From: " + r.Details.From + "!\n"
	toHeader := "To: " + r.Details.To + "!\n"
	subject := "Subject: " + r.Details.Subject + "!\n"
	auth := smtp.PlainAuth("", r.Config.Username, r.Config.Password, r.Config.Server)
	addr := fmt.Sprintf("%s:%s", r.Config.Server, r.Config.Port)
	msg := []byte(toHeader + fromHeader + subject + mime + "\n" + r.Details.Body)
	err := smtp.SendMail(addr, auth, r.Details.From, []string{r.Details.To}, msg)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
