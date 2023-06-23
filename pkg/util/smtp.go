package util

import (
	"time"

	"github.com/mizumoto-cn/dailyreport/conf"
	"gopkg.in/gomail.v2"
)

func NewSmtpDialer(conf *conf.SmtpDialer) *gomail.Dialer {
	d := gomail.NewDialer(conf.Host, int(conf.Port), conf.Username, conf.Password)
	return d
}

type EmailFormatter func([]string, string) *gomail.Message

// . NewEmailFormatter
func NewEmailFormatter(conf *conf.SmtpDialer) EmailFormatter {
	return func(to []string, token string) *gomail.Message {
		currentDate := time.Now().Format("20060102")
		body := `<!DOCTYPE html>
		<html>
		<head>
			<meta charset="utf-8">
			<title>進捗報告 ` + currentDate + `</title>
		</head>
		<body>
			<p>本日は以下の仕事をしました。</p>
			<br />
			<p>` + token + `</p>
			<br />
			<p>--</p>
			<p>株式会社  イー・ビジネス　　</p>
			<p>イノベーション事業部　</p>
			<p>DXソリューション部</p>
			<p><b>徐　瑞元</b></p>
			<p>Mobile: 080-7101-8913</p>
			<p>E-mail: xurb@revolvesys.co.jp</p>
		</body>
		</html>
		`
		m := gomail.NewMessage()
		m.SetHeader("From", conf.Username)
		m.SetHeader("To", to...)
		m.SetHeader("Subject", "Reset Password")
		m.SetBody("text/html", body)
		return m
	}
}
