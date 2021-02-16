package goverify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// VerifySend Sending verification
func (t *TwilioClient) VerifySend(i *VerifyInput) (*VerifySendResponse, error) {
	data := url.Values{}
	url := t.VerifyURL + t.TwilioService + "/Verifications"

	data.Set("To", i.To)
	data.Set("Channel", i.Channel)

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(t.AccountSid, t.AuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		exception := new(Exception)
		if err := json.Unmarshal(body, exception); err != nil {
			return nil, err
		}
		return nil, exception
	}

	out := new(VerifySendResponse)
	if err := json.Unmarshal(body, out); err != nil {
		return nil, err
	}
	return out, err
}
