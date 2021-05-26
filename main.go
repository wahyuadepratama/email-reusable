package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	gomail "gopkg.in/mail.v2"
)

func main() {
	router := gin.Default()
	router.POST("api/email/send", SendEmail)
	router.Run(":8000")
}

// DataEmail to receive data
type DataEmail struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Subject  string `json:"subject"`
	Message  string `json:"message"`
	SMTP     string `json:"smtp"`
	Port     int    `json:"port"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SendEmail to send from smtp
func SendEmail(c *gin.Context) {
	m := gomail.NewMessage()

	var input DataEmail
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set E-Mail sender
	m.SetHeader("From", input.From)

	// Set E-Mail receivers
	m.SetHeader("To", input.To)

	// Set E-Mail subject
	m.SetHeader("Subject", input.Subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", input.Message)

	// Settings for SMTP server
	d := gomail.NewDialer(input.SMTP, input.Port, input.Email, input.Password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully Send Email"),
	})
}
