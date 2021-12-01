package qq

import "encoding/json"

type ErrorResponse struct {
	Error int `json:"error"`
	ErrorDescription string `json:"error_description"`
}


func isError(body []byte) bool {
	r, _ := checkError(body)
	if r.Error != 0 {
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