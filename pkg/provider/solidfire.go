package provider

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

// APIError wrapper
type APIError struct {
	ID    int `json:"id"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Name    string `json:"name"`
	} `json:"error"`
}

// Transport is an interface to the actual http request to the SF endpoint.  We use an interface here so
// that we can easily Mock the SolidFire API server and it's responses in unit tests etc
type Transport interface {
	Post(string, []byte) (*http.Response, error)
}

// HTTP is the standard implementation of Transport, just simply implements Post using
// net/http pkg
type HTTP struct{}

// Request holds all the required info to issue a SolidFire API request.  The Name of the method,
// the URL of the SolidFire endpoint, an ID to use to reconnect to an async request and the
// Params (parameters) for the desired method.
// for the desired method
type Request struct {
	Name   string
	URL    string
	ID     int64
	Params interface{}
	Debug  bool
}

// Post issues the actual request to the SF endpoint.  Requires a url to the endpoint of the form:
// https://admin:admin@10.117.36.101/json-rpc/9.0 and the payload containing the method and it's
// associated parameters
func (t HTTP) Post(url string, payload []byte) (resp *http.Response, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	http := &http.Client{Transport: tr}
	resp, err = http.Post(url, "json-rpc", strings.NewReader(string(payload)))
	return resp, err
}

// NewReqID simply returns a random integer to be used as a request ID
// This ID is a handle to keep track of requests/responses and to help
// with long running commands
func NewReqID() int {
	return rand.Intn(1000-1) + 1
}

// Check is a simple helper function to check for error response.  We're
// using this just for
func Check(f func() error) {
	if err := f(); err != nil {
		log.Printf("error in deferred check:%v\n", err)
	}
}

// DecodeResponse is a helper method to convert the raw response from the SolidFire API
// into the appropriate response data type
func DecodeResponse(response []byte, responseType interface{}) (interface{}, error) {
	if err := json.Unmarshal([]byte(response), &responseType); err != nil {
		log.Printf("error detected unmarshalling response: %v", err)
		return nil, err
	}
	return responseType, nil
}

// IssueRequest does the actual work of sending a request to the SolidFire API server.  We require
// a Request item and a transport.  The Request item contains all of the needed info like the URL
// for the SolidFire API Server, the Name of the method to execute and the Parameters for the call.
func IssueRequest(r Request, t Transport) ([]byte, error) {
	var prettyJSON bytes.Buffer
	data, err := json.Marshal(map[string]interface{}{
		"method": r.Name,
		"id":     r.ID,
		"params": r.Params,
	})

	if err != nil {
		log.Printf("error marshalling api request: %v", err)
		return nil, errors.New("device API error")
	}
	if r.Debug {
		log.Printf("issuing request to SolidFire endpoint:  %+v", string(data))
	}

	_ = json.Indent(&prettyJSON, data, "", "  ")
	resp, err := t.Post(r.URL, data)
	if err != nil {
		if r.Debug {
			log.Printf("error response from SolidFire API request: %v", err)
		}
		return nil, errors.New("device API error")
	}

	if strings.Contains(resp.Status, "Unauthorized") {
		if r.Debug {
			log.Printf("attempted command returned unauthorized, command: %+v, response: %+v", r.Name, resp.Status)
		}
		return nil, errors.New("unauthorized request")
	}

	defer Check(resp.Body.Close)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}

	_ = json.Indent(&prettyJSON, body, "", "  ")
	errresp := APIError{}
	err = json.Unmarshal([]byte(body), &errresp)
	if err != nil {
		return body, err
	}

	if errresp.Error.Code != 0 {
		return body, fmt.Errorf("device API error: %+v", errresp.Error.Name)
	}
	return body, nil
}
