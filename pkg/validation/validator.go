package validation

import (
	"errors"
	"github.com/dlclark/regexp2"
)

type (
	Config struct {
		MessageRegexp string
	}

	Validator struct {
		messageRegexp *regexp2.Regexp
	}
)

func NewValidator(config Config) (*Validator, error) {
	messageRegex, err := regexp2.Compile(config.MessageRegexp, 0)
	if err != nil {
		return nil, err
	}
	return &Validator{
		messageRegexp: messageRegex,
	}, nil
}

func (v *Validator) Message(msg string) error {
	ok, _ := v.messageRegexp.MatchString(msg)
	if !ok {
		return errors.New("invalid message")
	}
	return nil
}
