package mailing

type MailingService interface {
	SendSingleEmail(headers *Headers, template Template, params map[string]interface{}) error
}
