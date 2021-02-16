package goverify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// VerifyCheck Code verification
func (t *TwilioClient) VerifyCheck(i *VerifyInput) (*VerifyCheckResponse, error) {
	data := url.Values{}
	url := t.VerifyURL + t.TwilioService + "/VerificationCheck"

	data.Set("To", i.To)
	data.Set("Code", i.Code)

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

	if resp.StatusCode != http.StatusOK {
		exception := new(Exception)
		if err := json.Unmarshal(body, exception); err != nil {
			return nil, err
		}
		return nil, exception
	}

	out := new(VerifyCheckResponse)
	if err := json.Unmarshal(body, out); err != nil {
		return nil, err
	}
	return out, err
}
