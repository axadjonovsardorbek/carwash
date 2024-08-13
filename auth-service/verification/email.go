package verification

import (
	"html/template"
	"log"
	"net/smtp"
	"os"
	"strings"
)

type Params struct {
	From     string
	Password string
	To       string
	Message  string
	Code     string
	UserName string
}

func SendVerificationCode(params Params) error {
	htmlFile, err := os.ReadFile("verification/format.html")
	if err != nil {
		log.Println("Cannot read html file", err.Error())
		return err
	}
	temp, err := template.New("email").Parse(string(htmlFile))
	if err != nil {
		log.Println("Cannot parse html file")
		return err
	}

	var Builder strings.Builder
	err = temp.Execute(&Builder, params)
	if err != nil {
		log.Println("Cannot execute HTML", err.Error())
		return err
	}

	message := "From: " + params.From + "\n" + "To: " + params.To + "\n" + "Message: " + params.Message + "\n" + "MIME-Version: 1.0\n" + "Content-type: text/html; charset=\"UTF-8\"\n" + "\n" + Builder.String()

	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("smtp", params.From, params.Password, "smtp.gmail.com"),
		params.From, []string{params.To}, []byte(message),
	)

	if err != nil {
		log.Println("Could not send an email", err.Error())
		return err
	}

	return nil
}
