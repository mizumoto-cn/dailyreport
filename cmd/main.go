package main

import (
	"context"
	"flag"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/mizumoto-cn/dailyreport/conf"
	"github.com/mizumoto-cn/dailyreport/pkg/post"
	"github.com/mizumoto-cn/dailyreport/pkg/util"
)

var (
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "..", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()
	if err := c.Load(); err != nil {
		panic(err)
	}
	var sdc conf.SmtpDialer
	if err := c.Scan(&sdc); err != nil {
		panic(err)
	}
	post.DefaultPosterFactoryMap().Register(
		"smtp",
		post.PosterFactoryFunc(func(ctx context.Context, pc any) (post.Poster, error) {
			return post.NewSmtpMailPoster(util.NewEmailFormatter(pc.(*conf.SmtpDialer)), pc.(*conf.SmtpDialer)), nil
		}),
		&sdc,
	)
	var pc conf.Path
	if err := c.Scan(&pc); err != nil {
		panic(err)
	}
	token, err := util.ReadFile(pc.Path)
	if err != nil {
		panic(err)
	}
	if err := post.Post(context.Background(), "smtp", token); err != nil {
		panic(err)
	}
}
