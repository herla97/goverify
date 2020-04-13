package goverify

const (
	verifyURL = "https://verify.twilio.com/v2/Services/"
)

// TwilioClient Structure to receive the parameters of a Twilio account.
type TwilioClient struct {
	AccountSid    string
	AuthToken     string
	TwilioService string
	VerifyURL     string
}

// VerifySMSResponse Twilio Verify API response structure
// See more: https://www.twilio.com/docs/verify/api/verification
type VerifySMSResponse struct {
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

// Lookup Twilio Verify API response structure
type Lookup struct {
	Carrier Carrier `json:"carrier"`
}

// Carrier Twilio Verify API response structure
type Carrier struct {
	ErrorCode         int    `json:"error_code,omitempty"`
	MobileCountryCode string `json:"mobile_country_code"`
	MobileNetworkCode string `json:"mobile_network_code"`
	Name              string `json:"name"`
	Type              string `json:"type"`
}

// SendCodeAttempt Twilio Verify API response structure
type SendCodeAttempt struct {
	Channel string `json:"channel"`
	Time    string `json:"time"`
}

// Exception Twilio Verify API errors structure
// See more: https://www.twilio.com/docs/usage/troubleshooting/debugging-your-application
type Exception struct {
	Code     int64  `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int64  `json:"status"`
}

// NewVerify Enable the use of goverify functions.
func NewVerify(accountSid, authToken, twilioService string) *TwilioClient {
	return &TwilioClient{
		AccountSid:    accountSid,
		AuthToken:     authToken,
		TwilioService: twilioService,
		VerifyURL:     verifyURL,
	}
}
