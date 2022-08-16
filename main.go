package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Could not get env values")
	}
	from := os.Getenv("username")
	password := os.Getenv("password")

	SendingInfos := []string{"vijaydinesh77vj@gmail.com","nelsonct7@gmail.com","abhijithv108@gmail.com"}

	host := "smtp.gmail.com"

	port := "587"

	msg := "testing message for checking the working of smtp server"

	body := []byte(msg)

	auth := smtp.PlainAuth("", from, password, host)

	err = smtp.SendMail(host+":"+port, auth, from, SendingInfos, body)

	if err != nil {
		fmt.Println("Error occured at while sending mail.", err)
		os.Exit(1)
	}

	fmt.Println("Email send successfully.")

}
