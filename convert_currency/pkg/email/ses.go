package mail

import (
	"convert-service/pkg/log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

// SES Configuration Struct
type mailSESConfig struct {
	Sender    string
	AccessKey string
	SecretKey string
	Region    string
}

// SES Configuration Variable
var sesCfg mailSESConfig

// SES Variable
var sesClient *ses.SES

func sesConnect() *ses.SES {
	// Initialize Connection
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(sesCfg.AccessKey, sesCfg.SecretKey, ""),
		Region:      aws.String(sesCfg.Region)},
	)
	if err != nil {
		log.Println(log.LogLevelFatal, "mail-ses-connect", err.Error())
	}

	// Return Connection
	return ses.New(sess)
}

func SendEmailWithTemplate(email *Email) error {
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(email.Recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(email.CharSet),
					Data:    aws.String(email.HtmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(email.CharSet),
					Data:    aws.String(email.TextBody),
				},
			},

			Subject: &ses.Content{
				Charset: aws.String(email.CharSet),
				Data:    aws.String(email.Subject),
			},
		},
		Source: aws.String(sesCfg.Sender),
	}

	_, err := sesClient.SendEmail(input)
	return err
}
