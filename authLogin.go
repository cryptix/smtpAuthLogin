package smtpAuthLogin

import (
	"fmt"
	"net/smtp"
)

// mailing helpers
type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", nil, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			buf := []byte(fmt.Sprintf("%s", a.username))
			return buf, nil
		case "Password:":
			buf := []byte(fmt.Sprintf("%s", a.password))
			return buf, nil
		default:
			return nil, fmt.Errorf("Unknown next from server")
		}
	}

	return nil, nil
}
