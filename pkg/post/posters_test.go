package post

import (
	"testing"

	"github.com/mizumoto-cn/dailyreport/conf"
	"github.com/stretchr/testify/assert"
	"gopkg.in/gomail.v2"
)

func TestNewSMTPMailPoster(t *testing.T) {
	formatter := func(to []string, token ...string) *gomail.Message {
		m := gomail.NewMessage()
		m.SetHeader("From", "test from")
		m.SetHeader("To", to...)
		m.SetHeader("Subject", "test subject")
		m.SetBody("text/html", "test body")
		return m
	}
	cfg := &conf.SmtpDialer{
		Host:     "smtp.gmail.com",
		Port:     587,
		Username: "test@gmail.com",
		Password: "testpassword",
	}
	p := NewSmtpMailPoster(formatter, cfg)
	assert.NotNil(t, p)
	assert.Equal(t, "smtp.gmail.com", p.cfg.Host)
	assert.Equal(t, 587, p.cfg.Port)
	assert.Equal(t, "test@gmail.com", p.cfg.Username)
	assert.Equal(t, "testpassword", p.cfg.Password)
	assert.Equal(t, formatter, p.formatter)
}
