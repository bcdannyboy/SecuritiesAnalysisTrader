package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
)

func SendHTTPPOSTRequest(url string, headers map[string]string, cookies map[string]string, bodyData interface{}, ignoreSSLErrors bool) (int, []byte, error) {
	/*
	 * Sends an HTTP POST request to the specified URL.
	 *
	 * Parameters:
	 * 	url (string): The URL to send the request to.
	 * 	headers (map[string]string): The headers to send with the request.
	 * 	cookies (map[string]string): The cookies to send with the request.
	 * 	bodyData (interface{}): The body data to send with the request.
	 * 	ignoreSSLErrors (bool): Whether or not to ignore SSL errors.
	 *
	 * Returns:
	 * 	int: The status code of the response.
	 * 	[]byte: The response object.
	 * 	error: The error object.
	 */

	StatusCode := 0
	ResponseObject := []byte{}

	jar, _ := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: nil,
	})

	client := &http.Client{
		Jar: jar,
	}

	if ignoreSSLErrors {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	jsonBody, err := json.Marshal(bodyData)
	if err != nil {
		return StatusCode, ResponseObject, fmt.Errorf("failed to marshal body data: %s", err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return StatusCode, ResponseObject, fmt.Errorf("failed to initiate POST request: %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	for name, value := range cookies {
		req.AddCookie(&http.Cookie{Name: name, Value: value})
	}

	resp, err := client.Do(req)
	if err != nil {
		return StatusCode, ResponseObject, fmt.Errorf("failed to send POST request: %s", err.Error())
	}
	defer resp.Body.Close()

	StatusCode = resp.StatusCode

	// Reading the response body using io.ReadAll
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return StatusCode, ResponseObject, fmt.Errorf("failed to read response body: %s", err.Error())
	}
	ResponseObject = responseBody

	return StatusCode, ResponseObject, nil
}

func SendHTTPGETRequest(url string, headers map[string]string, cookies map[string]string, ignoreSSLErrors bool) (int, []byte, error) {
	/*
	 * Sends an HTTP GET request to the specified URL.
	 *
	 * Parameters:
	 * 	url (string): The URL to send the request to.
	 * 	headers (map[string]string): The headers to send with the request.
	 * 	cookies (map[string]string): The cookies to send with the request.
	 * 	ignoreSSLErrors (bool): Whether or not to ignore SSL errors.
	 *
	 * Returns:
	 * 	int: The status code of the response.
	 * 	[]byte: The response object.
	 * 	error: The error object.
	 */

	StatusCode := 0
	ResponseObject := []byte{}

	jar, _ := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: nil,
	})

	client := &http.Client{
		Jar: jar,
	}

	if ignoreSSLErrors {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return StatusCode, ResponseObject, fmt.Errorf("failed to initiate GET request: %s", err.Error())
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	for name, value := range cookies {
		req.AddCookie(&http.Cookie{Name: name, Value: value})
	}

	resp, err := client.Do(req)
	if err != nil {
		return StatusCode, ResponseObject, fmt.Errorf("failed to send GET request: %s", err.Error())
	}
	defer resp.Body.Close()

	StatusCode = resp.StatusCode

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return StatusCode, ResponseObject, fmt.Errorf("failed to read response body: %s", err.Error())
	}
	ResponseObject = responseBody

	return StatusCode, ResponseObject, nil
}
