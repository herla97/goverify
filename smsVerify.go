package goverify

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

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

	// err = json.NewDecoder(resp.Body).Decode(&smsResp)
	// if err != nil {
	// 	return smsResp, exception, err
	// }

	if resp.StatusCode != http.StatusCreated {
		exception, err = UnmarshalException(respBody)
		return smsResp, exception, err
	}

	smsResp, err := UnmarshalVerifySMSResponse(respBody)

	return smsResp, exception, err
}

// func VerifyCheckSMS() {

// }
