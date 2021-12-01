package wechat

import "encoding/json"

type ErrorResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func isError(body []byte) bool {
	r, _ := checkError(body)
	if r.Errcode != 0 {
		return true
	}
	return false
}

func checkError(body []byte) (*ErrorResponse, error) {
	errresp := &ErrorResponse{}
	if err := json.Unmarshal(body, errresp); err != nil {
		return nil, err
	}
	return errresp, nil
}
