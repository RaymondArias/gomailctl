package mailsend

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"time"
)

// Mail contains data about email and user auth
type Mail struct {
	From       string
	Recipients string
	Data       []byte
	Username   string
	Password   string
	SMTPServer string
	SMTPPort   string
}

// SendMail sends an email based on the data in Mail passed in
func SendMail(email Mail) (bool, error) {
	smtpHost := fmt.Sprintf("%s:%s", email.SMTPServer, email.SMTPPort)

	tlsconfig := &tls.Config{
		ServerName: email.SMTPServer,
	}

	dialer := net.Dialer{}
	dialer.Timeout = 10 * time.Second
	conn, err := tls.DialWithDialer(&dialer, "tcp", smtpHost, tlsconfig)
	if err != nil {
		log.Panic(err.Error())
		return false, err
	}

	auth := smtp.PlainAuth("", email.Username, email.Password, smtpHost)
	msg := email.Data
	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		log.Panic(err.Error())
		return false, err
	}

	if err = c.Auth(auth); err != nil {
		log.Panic(err.Error())
		return false, err
	}

	if err = c.Mail(email.From); err != nil {
		log.Panic(err.Error())
		return false, err
	}

	if err = c.Rcpt(email.Recipients); err != nil {
		log.Panic(err.Error())
		return false, err
	}

	w, err := c.Data()
	if err != nil {
		log.Panic(err.Error())
		return false, err
	}

	_, err = w.Write([]byte(msg))
	if err != nil {
		log.Panic(err.Error())
		return false, err
	}

	err = w.Close()
	if err != nil {
		log.Panic(err.Error())
		return false, err
	}

	c.Quit()

	if err != nil {
		log.Panic(err.Error())
		return false, err
	}

	return true, nil
}
