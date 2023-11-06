package mail

import "explore_address/pkg/server"

func init() {
	sesCfg.Sender = server.Config.GetString("MAIL_SENDER")
	sesCfg.AccessKey = server.Config.GetString("MAIL_ACCESS_KEY")
	sesCfg.SecretKey = server.Config.GetString("MAIL_SECRET_KEY")
	sesCfg.Region = server.Config.GetString("MAIL_REGION")

	sesClient = sesConnect()
}

type Email struct {
	Recipient string
	CharSet   string
	Subject   string
	HtmlBody  string
	TextBody  string
}
