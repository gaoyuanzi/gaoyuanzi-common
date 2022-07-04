package mail

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

type Mail struct {
	From   string
	Host   string
	Port   int
	Usr    string
	Passwd string
}

func NewMail(host string, port int, usr string, passwd string) (*Mail, error) {
	ma := &Mail{
		From:   usr,
		Host:   host,
		Port:   port,
		Usr:    usr,
		Passwd: passwd,
	}
	return ma, nil
}

func (m *Mail) SendMsg(tousr string, subject string, context string, attach string) error {
	m1 := gomail.NewMessage()
	m1.SetHeader("From", m.From)
	m1.SetHeader("To", tousr)
	//m1.SetAddressHeader("Cc", m.From, "lhzd863")
	m1.SetHeader("Subject", subject)
	m1.SetBody("text/html", context)
	if len(attach) != 0 {
		_, err := os.Stat(attach)
		if err == nil {
			m1.Attach(attach)
		} else {
			log.Println(err)
		}
	}

	d := gomail.NewDialer(m.Host, m.Port, m.Usr, m.Passwd)

	if err := d.DialAndSend(m1); err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("mail send success...")
	return nil
}
