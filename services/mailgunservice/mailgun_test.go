package mailgunService

import (
	"testing"

	"github.com/epa-datos/mailing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSendSingleEmail(t *testing.T) {
	Convey("Given a mailgun service", t, func() {
		mgService := New()
		Convey("when calling SendSingleEmail", func() {
			err := mgService.SendSingleEmail(&mailing.Headers{
				Recipient: "kgonzalez@epa.digital",
				Sender:    "test@test.com",
				Subject:   "Test email",
			}, mailing.TestTemplate, map[string]interface{}{
				"Name": "Test user",
			})
			So(err, ShouldBeNil)
		})
	})
}
