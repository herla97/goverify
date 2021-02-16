package goverify

const (
	verifyURL = "https://verify.twilio.com/v2/Services/"
)

type (
	// VerifyInput Struture for set input to VerifySend func
	VerifyInput struct {
		To      string `json:"to"`
		Channel string `json:"channel"`
		Code    string `json:"code"`
	}

	// TwilioClient Structure to receive the parameters of a Twilio account
	TwilioClient struct {
		AccountSid    string
		AuthToken     string
		TwilioService string
		VerifyURL     string
	}

	// VerifySendResponse Twilio Verify API send response structure
	// See more: https://www.twilio.com/docs/verify/api/verification
	VerifySendResponse struct {
		AccountSid       string            `json:"account_sid"`
		Amount           string            `json:"amount,omitempty"`
		Channel          string            `json:"channel"`
		DateCreated      string            `json:"date_created"`
		DateUpdated      string            `json:"date_updated"`
		Lookup           Lookup            `json:"lookup"`
		Payee            string            `json:"payee"`
		SendCodeAttempts []SendCodeAttempt `json:"send_code_attempts"`
		ServiceSid       string            `json:"service_sid"`
		Sid              string            `json:"sid"`
		Status           string            `json:"status"`
		To               string            `json:"to"`
		URL              string            `json:"url"`
		Valid            bool              `json:"valid"`
	}

	// VerifyCheckResponse Twilio Verify API check response structure
	// See more: https://www.twilio.com/docs/verify/api/verification-check
	VerifyCheckResponse struct {
		AccountSid  string `json:"account_sid"`
		Amount      string `json:"amount,omitempty"`
		Channel     string `json:"channel"`
		DateCreated string `json:"date_created"`
		DateUpdated string `json:"date_updated"`
		Payee       string `json:"payee,omitempty"`
		ServiceSid  string `json:"service_sid"`
		Sid         string `json:"sid"`
		Status      string `json:"status"`
		To          string `json:"to"`
		Valid       bool   `json:"valid"`
	}

	// Lookup Twilio Verify API response structure
	Lookup struct {
		Carrier Carrier `json:"carrier"`
	}

	// Carrier Twilio Verify API response structure
	Carrier struct {
		ErrorCode         int    `json:"error_code,omitempty"`
		MobileCountryCode string `json:"mobile_country_code"`
		MobileNetworkCode string `json:"mobile_network_code"`
		Name              string `json:"name"`
		Type              string `json:"type"`
	}

	// SendCodeAttempt Twilio Verify API response structure
	SendCodeAttempt struct {
		Channel string `json:"channel"`
		Time    string `json:"time"`
	}
)

// NewVerify Enable the use of goverify functions.
func NewVerify(accountSid, authToken, twilioService string) *TwilioClient {
	return &TwilioClient{
		AccountSid:    accountSid,
		AuthToken:     authToken,
		TwilioService: twilioService,
		VerifyURL:     verifyURL,
	}
}
