package models

import "encoding/json"

func UnmarshalVerifySMSResponse(data []byte) (VerifySMSResponse, error) {
	var r VerifySMSResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *VerifySMSResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

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

type Lookup struct {
	Carrier Carrier `json:"carrier"`
}

type Carrier struct {
	ErrorCode         int    `json:"error_code,omitempty"`
	MobileCountryCode string `json:"mobile_country_code"`
	MobileNetworkCode string `json:"mobile_network_code"`
	Name              string `json:"name"`
	Type              string `json:"type"`
}

type SendCodeAttempt struct {
	Channel string `json:"channel"`
	Time    string `json:"time"`
}
