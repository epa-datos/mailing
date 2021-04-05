package mailing

import "github.com/epa-datos/mailing/templates"

type baseTemplate string

type Template interface {
	String() string
}

const (
	PasswordRecovery = baseTemplate("password_recovery.html")
	UserInvitation   = baseTemplate("invitation.html")
	TestTemplate     = baseTemplate("test_template.html")
)

var baseTemplates = map[Template]string{
	TestTemplate:     templates.TestTemplate,
	UserInvitation:   templates.UserInvitation,
	PasswordRecovery: templates.PasswordRecovery,
}

func (t baseTemplate) String() string {
	return baseTemplates[t]
}
