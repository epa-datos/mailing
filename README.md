# mailing

## Summary
EPA digital mailing package

## Motivation
* This package aims to provide a simple API for mailing services.
* This package connects with the following mailing providers:
  * Mailgun
* This package allows:
  * Sending single emails

## API 

To start using this package you need to include the following lines in your imports
```golang
import(
	"github.com/epa-datos/mailing"
	"github.com/epa-datos/mailing/services/{MAILING_SERVICE}"
 )
```
### Sending a single email
To send an email you need to use the following method
```golang
  err := mail.SendSingleEmail(headers, template, params)
```
Single emails have three input values:   
**Headers**
```golang
  headers := &mailing.Headers{
	  Sender:    "test@epa.ditigal",
		Recipient: "recipient@test.com",
		Subject:   "test",
		ContentType: "text/html",           //This value is optional
```
**Template**
```golang
  template := mailing.TestTemplate      //This is a constant of type template. Valid templates are defined in mailing/templates.go
```

**Params**
 ```golang
  params := map[string]interface{}{
    "Example":"test",                   //Valid params depend on template placeholders
  }
 ```

### Mailgun specific implementation
```golang
	mailService := mailgunService.New()
```
**Note:** You need to set the EMAIL_DOMAIN and EMAIL_API_KEY as environment variables











