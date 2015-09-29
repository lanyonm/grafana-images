package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	post_data = `{
		"imageUrl": "http://httpstat.us/418?panelId=5&width=1000&height=500&from=now-6h&to=now&var-server=test-server"
	}`
	host      = "http://grafana.example.com/saved-images"
	imagePath = os.TempDir()
)

func TestGrafanaImagesHandlerSuccess(t *testing.T) {
	req, _ := http.NewRequest("POST", "/r", bytes.NewBufferString(post_data))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer 1234567543ewsfdgdh432345awdf=")
	resp := httptest.NewRecorder()

	GrafanaImagesHandler(host, imagePath)(resp, req)

	const expected_response_code = 200
	if code := resp.Code; code != expected_response_code {
		t.Errorf("received %v response code, expected %v", code, expected_response_code)
	}
	const expected_response_body = `{"pubImg":"http://grafana.example.com/saved-images/6d3b78cb8bcffe91aa5afaf869f070e3.png"}`
	if ret := resp.Body.String(); ret != expected_response_body {
		t.Errorf("received %v response body, expected %v", ret, expected_response_body)
	}
}

func TestGrafanaImagesHandlerNotPOST(t *testing.T) {
	req, _ := http.NewRequest("GET", "/r", bytes.NewBufferString(post_data))
	resp := httptest.NewRecorder()

	GrafanaImagesHandler(host, imagePath)(resp, req)

	const expected_response_code = 405
	if code := resp.Code; code != expected_response_code {
		t.Errorf("received %v response code, expected %v", code, expected_response_code)
	}
}

func TestGrafanaImagesHandlerInvalidJSON(t *testing.T) {
	req, _ := http.NewRequest("POST", "/r", bytes.NewBufferString(`{invalid:"json"}`))
	resp := httptest.NewRecorder()

	GrafanaImagesHandler(host, imagePath)(resp, req)

	const expected_response_code = 400
	if code := resp.Code; code != expected_response_code {
		t.Errorf("expected %v, but received %v response code", expected_response_code, code)
	}
	const expected_error_message = "Error parsing JSON\n"
	if msg := resp.Body.String(); msg != expected_error_message {
		t.Errorf("expected \"%v\", but found \"%v\" error message", expected_error_message, msg)
	}
}
