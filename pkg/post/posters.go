package post

import (
	"context"

	"github.com/mizumoto-cn/dailyreport/conf"
	"github.com/mizumoto-cn/dailyreport/pkg/util"
)

type SMTPMailPoster struct {
	formatter util.EmailFormatter
	cfg       *conf.SmtpDialer
}

func (p SMTPMailPoster) Post(ctx context.Context, token ...string) error {
	m := p.formatter((p.cfg.To), token...)
	if err := util.NewSmtpDialer(p.cfg).DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func NewSmtpMailPoster(formatter util.EmailFormatter, cfg *conf.SmtpDialer) *SMTPMailPoster {
	return &SMTPMailPoster{
		formatter: formatter,
		cfg:       cfg,
	}
}
