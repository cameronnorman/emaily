package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
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
	Details      emailDetails           `json:"details"`
	Config       emailConfig            `json:"config"`
	TemplateName string                 `json:"template_name"`
	Data         map[string]interface{} `json:"data"`
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/send", handleSendEmailRequest)
	e.POST("/send_with_template", handleSendTemplateRequest)
	e.Logger.Fatal(e.Start(":8081"))
}

func handleSendEmailRequest(c echo.Context) error {
	r := sendEmailRequest{}
	if err := c.Bind(&r); err != nil {
		return err
	}

	go sendMail(r)

	return nil
}

func handleSendTemplateRequest(c echo.Context) error {
	templatesPath := "templates/"
	if os.Getenv("TEMPLATES_PATH") != "" {
		templatesPath = os.Getenv("TEMPLATES_PATH")
	}

	r := sendEmailRequest{}
	if err := c.Bind(&r); err != nil {
		return err
	}

	t, err := template.ParseFiles(templatesPath + r.TemplateName + ".html")
	if err != nil {
		c.JSON(422, err.Error())
		return err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, r.Data); err != nil {
		c.JSON(422, err.Error())
		return err
	}
	r.Details.Body = tpl.String()
	go sendMail(r)

	return nil
}

func sendMail(r sendEmailRequest) error {
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
