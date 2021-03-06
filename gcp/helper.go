package gcp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type apiErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Check HTTP response code, return error if HTTP request is not successed.
func apiResponseChecker(statusCode int, response []byte, funcName string) error {

	if statusCode >= 300 || statusCode < 200 {
		log.Printf("%s request failed", funcName)
		var error_response apiErrorResponse
		response_content := bytes.NewBuffer(response).String()
		if err := json.Unmarshal(response, &error_response); err != nil {
			log.Printf("Failed to unmarshall error response from %s", funcName)
			return fmt.Errorf(response_content)
		}
		return fmt.Errorf("code: %d, message: %s", error_response.Code, error_response.Message)
	}

	return nil

}
