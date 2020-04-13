package models

import "encoding/json"

func UnmarshalException(data []byte) (Exception, error) {
	var r Exception
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Exception) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Exception struct {
	Code     int64  `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int64  `json:"status"`
}
