package goverify

import "fmt"

// Exception Twilio Verify API errors structure
// See more: https://www.twilio.com/docs/usage/troubleshooting/debugging-your-application
type Exception struct {
	Code     int64  `json:"code,omitempty"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int64  `json:"status,omitempty"`
}

func (e *Exception) Error() string {
	return fmt.Sprintf("verify: %s", e.Message)
}
