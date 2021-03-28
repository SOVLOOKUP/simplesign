package simplesign

import (
	"github.com/gogf/gf/crypto/gaes"
	"github.com/gogf/gf/util/grand"
	mail "github.com/xhit/go-simple-mail/v2"
)

type (
	Simplesign interface {
		Send(to, title string) ([]byte, error)
		Verify(encaptcha []byte, captcha string) (bool, error)
	}

	S struct {
		key    []byte
		Server *mail.SMTPServer
		Client *mail.SMTPClient
	}
)

func NewS(host string, port int, username, password string) (*S, error) {
	server := mail.NewSMTPClient()

	// SMTP Server
	server.Host = host
	server.Port = port
	server.Username = username
	server.Password = password
	server.Encryption = mail.EncryptionSSL

	// SMTP client
	smtpClient, err := server.Connect()
	return &S{
		Server: server,
		Client: smtpClient,
	}, err
}

func (s *S) Send(to, title string) ([]byte, error) {
	if s.key == nil {
		s.key = grand.B(32)
	}
	captcha := []byte(grand.Digits(5))
	encaptcha, _ := gaes.Encrypt(captcha, s.key)
	err := s.newMail(to, title, string(captcha)).Send(s.Client)
	return encaptcha, err
}

func (s *S) Verify(encaptcha []byte, captcha string) (bool, error) {
	decaptcha, err := gaes.Decrypt(encaptcha, s.key)

	if string(decaptcha) == captcha {
		return true, err
	}
	return false, err
}

func (s *S) newMail(to, title, content string) *mail.Email {
	email := mail.NewMSG()
	email.SetFrom(s.Server.Username).
		AddTo(to).
		SetSubject(title)
	email.SetBody(mail.TextHTML, content)
	return email
}
