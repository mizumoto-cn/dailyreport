package util

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/mizumoto-cn/dailyreport/conf"
	"gopkg.in/gomail.v2"
)

func NewSmtpDialer(conf *conf.SmtpDialer) *gomail.Dialer {
	d := gomail.NewDialer(conf.Host, int(conf.Port), conf.Username, conf.Password)
	return d
}

type EmailFormatter func([]string, ...string) *gomail.Message

// . NewEmailFormatter
func NewEmailFormatter(conf *conf.SmtpDialer) EmailFormatter {
	tPath := conf.TemplatePath
	f, err := os.Open(tPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	template := ""
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		template += line
	}
	return func(to []string, token ...string) *gomail.Message {
		currentDate := time.Now().Format("20060102")
		body := fmt.Sprintf(template, convertToAnySlice(token)...)
		m := gomail.NewMessage()
		m.SetHeader("From", conf.Username)
		m.SetHeader("To", to...)
		m.SetHeader("Subject", "進捗報告 "+currentDate)
		m.SetBody("text/html", body)
		return m
	}
}

func convertToAnySlice(slice []string) []any {
	is := make([]any, len(slice))
	for i, s := range slice {
		is[i] = s
	}
	return is
}
