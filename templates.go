package mailing

import "os"

type Template string

const (
	PasswordRecovery = Template("password_recovery.html")
	UserInvitation   = Template("invitation.html")
	TestTemplate     = Template("test_template.html")
)

func (t Template) GetPath() string {
	return os.Getenv("GOPATH") + "/src/github.com/epa-datos/mailing/templates/" + string(t)
}
