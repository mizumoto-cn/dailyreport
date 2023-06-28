package mock_gomail

import "gopkg.in/gomail.v2"

type Dialer interface {
	Dial() (gomail.SendCloser, error)
	DialAndSend(m ...*gomail.Message) error
}
