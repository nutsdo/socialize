package weibo

import "encoding/json"

type ErrorResponse struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"error_code"`
	Request   string `json:"request"`
	ErrorUri  string `json:"error_uri"`
	ErrorDescription string `json:"error_description"`
}

func isError(body []byte) bool {
	r, _ := checkError(body)
	if r.ErrorCode != 0 {
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
