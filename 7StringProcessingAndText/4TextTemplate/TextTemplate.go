package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type EmailData struct {
	RecipientName string
	SenderName    string
	Subject       string
	Body          string
	Items         []string
	UnreadCount   int
}

func main() {
	fmt.Println("--- Text template example ---")

	//email template containing placeholders, conditions and loops
	emailTemplate := `
	Hello {{.RecipientName}}

	Subject: {{ .Subject }}
	
	{{.Body}}
	
	{{if .Items}}
		Related Items:
	{{range .Items}}
		- {{.}}
	{{end}}
	{{end}}

	{{if gt .UnreadCount 0}}
	You have {{.UnreadCount}} unreads.
	{{else}}
	You have no messages
	{{end}}


	- Thanks
	{{.SenderName}}
	`

	//parse the template into a reusable template object
	tmpl, err := template.New("email-message").Parse(emailTemplate)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//data used to populate the template
	data := EmailData{
		RecipientName: "Alice",
		SenderName:    "Bob's Auto-Responder",
		Subject:       "Your Weekly Update",
		Body:          "Here is the update you requested. Hope you find it useful.",
		Items:         []string{"Report A", "Document B", "Summary C"},
		UnreadCount:   5,
	}

	//builder stores the rendered template efficiently
	var output strings.Builder

	//render the template using the supplied data
	err = tmpl.Execute(&output, data)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(strings.ToUpper(output.String()))
}
