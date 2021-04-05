package mailgunService

import (
	"context"
	"os"
	"time"

	"github.com/epa-datos/mailing"
	"github.com/mailgun/mailgun-go/v3"
)

type mailgunService struct {
	mgService *mailgun.MailgunImpl
}

func New() *mailgunService {
	mg := mailgun.NewMailgun(os.Getenv("EMAIL_DOMAIN"), os.Getenv("EMAIL_API_KEY"))
	return &mailgunService{
		mgService: mg,
	}
}

func (m *mailgunService) SendSingleEmail(headers *mailing.Headers, template mailing.Template, params map[string]interface{}) error {
	emailBody, err := mailing.ParseMailTemplate(template, params)
	if err != nil {
		return err
	}

	message := m.mgService.NewMessage(
		headers.Sender,
		headers.Subject,
		"",
		headers.Recipient,
	)
	message.SetHtml(emailBody)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err = m.mgService.Send(ctx, message)
	return err

}
