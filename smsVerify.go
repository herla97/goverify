package goverify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

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

// VerifySendSMS The function of sending SMS receives the recipient's number as a parameter.
func (t *TwilioClient) VerifySendSMS(to string) (smsResp *VerifySMSResponse, exception *Exception, err error) {
	msgData := url.Values{}
	client := &http.Client{}
	urlStr := t.VerifyURL + t.TwilioService + "/Verifications"

	msgData.Set("To", to)
	msgData.Set("Channel", "sms")

	msgDataReader := *strings.NewReader(msgData.Encode())

	req, err := http.NewRequest("POST", urlStr, &msgDataReader)
	if err != nil {
		return smsResp, exception, err
	}

	req.SetBasicAuth(t.AccountSid, t.AuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return smsResp, exception, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return smsResp, exception, err
	}

	if resp.StatusCode != http.StatusCreated {
		exception = new(Exception)
		err = json.Unmarshal(respBody, exception)
		return smsResp, exception, err
	}

	smsResp = new(VerifySMSResponse)
	err = json.Unmarshal(respBody, smsResp)

	return smsResp, exception, err
}

// func VerifyCheckSMS() {

// }
