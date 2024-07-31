package backend

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func Genotp() string {
	var digits = "0123456789"
	var result string
	for i := 0; i < 6; i++ {
		index := rand.Intn(len(digits))
		result += string(digits[index])
	}
	return result
}

func SendOtp(msg string, to []string) error {
	message := []byte(msg)
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASS")

	smtpServer := "smtp.gmail.com"
	port := "587"
	auth := smtp.PlainAuth("", from, password, smtpServer)
	err := smtp.SendMail(smtpServer+":"+port, auth, from, to, message)
	if err != nil {
		return err
	}
	fmt.Println("Email Sent")
	return nil
}

func CheckOtp(inputOtp string, generatedOtp string) bool {
	return inputOtp == generatedOtp
}
