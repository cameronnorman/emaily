package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/cameronnorman/emaily/pkg/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var Templates = map[string]*template.Template{}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	templatesPath := "./templates"
	if os.Getenv("TEMPLATES_PATH") != "" {
		templatesPath = os.Getenv("TEMPLATES_PATH")
	}

	items, _ := ioutil.ReadDir(templatesPath)
	for _, item := range items {
		filename := strings.TrimSuffix(item.Name(), filepath.Ext(item.Name()))
		path := fmt.Sprintf("%s/%s.html", templatesPath, filename)
		t, err := template.ParseFiles(path)
		if err != nil {
			panic(err)
		}

		Templates[filename] = t
	}

	e.GET("/check", handleHealthCheckRequest)
	e.POST("/send", handleSendEmailRequest)
	e.POST("/send_with_template", handleSendTemplateRequest)
	e.Logger.Fatal(e.Start(":8081"))
}

func handleHealthCheckRequest(c echo.Context) error {
	c.JSON(200, "OK")
	return nil
}

func handleSendEmailRequest(c echo.Context) error {
	r := service.Request{}
	if err := c.Bind(&r); err != nil {
		return err
	}

	go service.SendMail(r)

	return nil
}

func handleSendTemplateRequest(c echo.Context) error {
	r := service.Request{}
	if err := c.Bind(&r); err != nil {
		return err
	}

	t := Templates[r.TemplateName]
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, r.Data); err != nil {
		c.JSON(422, err.Error())
		return err
	}
	r.Details.Body = tpl.String()
	go service.SendMail(r)

	return nil
}
