package util

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/go-kratos/kratos/v2/log"
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
		template += line + "\n"
	}
	poweredBy := `
    <div class="signature" style="margin-top: 20px; text-align: right; background-color: #eee;">
      <p>
      Powered by <a href="https://github.com/mizumoto-cn/dailyreport">@mizumoto-cn/dailyreport</a>
      <img src="https://github.com/mizumoto-cn/mizumoto-cn/raw/main/image.png" width="16" height="16" style="margin-top: 4px"/>
      </p>
    </div>
`
	return func(to []string, token ...string) *gomail.Message {
		currentDate := time.Now().Format("20060102")
		body := fmt.Sprintf(template, convertToAnySlice(token)...)
		addPoweredBy(&body, poweredBy)
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

func addPoweredBy(body *string, poweredBy string) {
	// if contains </div> then add poweredBy before last </div>
	// else if contains </body> then add poweredBy before </body>
	// else add poweredBy before </html>
	// else add poweredBy at the end of body
	if i := lastIndexOf(*body, "</div>"); i != -1 {
		*body = (*body)[:i] + poweredBy + (*body)[i:]
	} else if i := lastIndexOf(*body, "</body>"); i != -1 {
		*body = (*body)[:i] + poweredBy + (*body)[i:]
	} else if i := lastIndexOf(*body, "</html>"); i != -1 {
		*body = (*body)[:i] + poweredBy + (*body)[i:]
	} else {
		*body += poweredBy
	}
	log.Debugf("body: %s", *body)
}

func lastIndexOf(s string, substr string) int {
	index := -1
	for i := 0; i < len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			index = i
		}
	}
	return index
}
