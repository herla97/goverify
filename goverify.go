package goverify

const (
	verifyURL = "https://verify.twilio.com/v2/Services/"
)

type TwilioClient struct {
	AccountSid    string
	AuthToken     string
	TwilioService string
	VerifyURL     string
}

func NewVerify(accountSid, authToken, twilioService string) *TwilioClient {
	return &TwilioClient{
		AccountSid:    accountSid,
		AuthToken:     authToken,
		TwilioService: twilioService,
		VerifyURL:     verifyURL,
	}
}
