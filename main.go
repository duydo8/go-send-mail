package main

import (
	"fmt"
	"net/smtp"
)

func SendMailSimple() {
	auth := smtp.PlainAuth("", "duydvph09704@fpt.edu.vn", "iepmdsgpxwfclimh", "smtp.gmail.com")
	msg := "Subject : My Special subject\n this is the body of email "
	err := smtp.SendMail("smtp.gmail.com:587", auth, "duydvph09704@fpt.edu.vn", []string{"duydogch123456@gmail.com"}, []byte(msg))
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	SendMailSimple()
}
