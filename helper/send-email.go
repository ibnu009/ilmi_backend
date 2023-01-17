package helper

import (
	"fmt"
	"net/smtp"
	"os"
	"strconv"
	"strings"

	"github.com/xlzd/gotp"
)

// config to v2 Gmail
func ConfigSendGmail(configAuthEmail, configAuthPassword, configSmtpHost, configSmtpPort string, toEmail []string, body []byte) error {
	auth := smtp.PlainAuth("", configAuthEmail, configAuthPassword, configSmtpHost)
	smtpAddr := fmt.Sprintf("%s:%v", configSmtpHost, configSmtpPort)

	err := smtp.SendMail(smtpAddr, auth, configAuthEmail, toEmail, body)
	if err != nil {
		return err
	}
	return nil
}

// body send email otp
func SendBodyOtpToEmail(toEmail []string, senderName, subject, message string) error {
	body := "From: " + senderName + "\n" +
		"To: " + strings.Join(toEmail, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	err := ConfigSendGmail(os.Getenv("CONFIG_AUTH_EMAIL"), os.Getenv("CONFIG_AUTH_PASSWORD"), os.Getenv("CONFIG_SMTP_HOST"), os.Getenv("CONFIG_SMTP_PORT"), toEmail, []byte(body))
	if err != nil {
		return err
	}
	return nil
}

func AutoGenerateOtp(secretLength int) uint32 {
	otpAutoGenerate := gotp.RandomSecret(secretLength)
	newOtp := gotp.NewDefaultTOTP(otpAutoGenerate).Now()
	//str to uint64
	convertOtpInt, _ := strconv.ParseInt(newOtp, 10, 32)

	return uint32(convertOtpInt)
}
