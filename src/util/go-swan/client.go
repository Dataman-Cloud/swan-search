package swan

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"sync"
)

// Swan is the interface to the Swan API
type Swan interface {
	// -- APPLICATIONS ---
	// create an application in swan
	CreateApplication(version *Version) (*Application, error)
	// get a list of applications from swan
	Applications(url.Values) ([]*Application, error)
	// delete an application in swan
	DeleteApplication(appID string) error
	// get an applications from swan
	GetApplication(appID string) (*Application, error)

	//-- SUBSCRIPTIONS--
	AddEventsListener() (EventsChannel, error)
}

var (
	// ErrInvalidEndpoint is thrown when the swan url specified was invalid
	ErrInvalidEndpoint = errors.New("invalid Swan endpoint specified")
	// ErrInvalidResponse is thrown when swan responds with invalid or error response
	ErrInvalidResponse = errors.New("invalid response from Swan")
	// ErrSwanDown is thrown when all the swan endpoints are down
	ErrSwanDown = errors.New("all the Swan hosts are presently down")
	// ErrTimeoutError is thrown when the operation has timed out
	ErrTimeoutError = errors.New("the operation has timed out")
)

type swanClient struct {
	sync.RWMutex
	// swanAddr
	swanAddr string
	// the http client use for making requests
	httpClient *http.Client
	// a custom logger for debug log messages
	debugLog *log.Logger
	// used for sse
	subscribedToSSE bool
}

// NewClient creates a new swan client
func NewClient(swanAddr string) (Swan, error) {
	debugLogOutput := ioutil.Discard
	return &swanClient{
		swanAddr:   swanAddr,
		httpClient: http.DefaultClient,
		debugLog:   log.New(debugLogOutput, "", 0),
	}, nil
}

func (r *swanClient) apiGet(uri string, post, result interface{}) error {
	return r.apiCall("GET", uri, post, result)
}

func (r *swanClient) apiPut(uri string, post, result interface{}) error {
	return r.apiCall("PUT", uri, post, result)
}

func (r *swanClient) apiPost(uri string, post, result interface{}) error {
	return r.apiCall("POST", uri, post, result)
}

func (r *swanClient) apiDelete(uri string, post, result interface{}) error {
	return r.apiCall("DELETE", uri, post, result)
}

func (r *swanClient) apiCall(method, uri string, body, result interface{}) error {
	var url string

	url = fmt.Sprintf("%s/%s", r.swanAddr, uri)

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}

	// step: create an API request
	request, err := r.apiRequest(method, url, bytes.NewReader(jsonBody))
	if err != nil {
		return err
	}

	response, err := r.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if len(jsonBody) > 0 {
		r.debugLog.Printf("apiCall(): %v %v %s returned %v %s\n", request.Method, request.URL.String(), jsonBody, response.Status, oneLogLine(respBody))
	} else {
		r.debugLog.Printf("apiCall(): %v %v returned %v %s\n", request.Method, request.URL.String(), response.Status, oneLogLine(respBody))
	}

	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		if result != nil {
			if err := json.Unmarshal(respBody, result); err != nil {
				//r.debugLog.Printf("apiCall(): failed to unmarshall the response from marathon, error: %s\n", err)
				fmt.Printf("apiCall(): failed to unmarshall the response from marathon, error: %s\n", err)
				return ErrInvalidResponse
			}
		}
		return nil
	}
	if response.StatusCode >= 400 {
		return errors.New(string(response.StatusCode))
	}
	return nil
	// TODO(xychu): better support for API ERROR
	//return NewAPIError(response.StatusCode, respBody)
}

// apiRequest creates a default API request
func (r *swanClient) apiRequest(method, url string, reader io.Reader) (*http.Request, error) {
	// Make the http request to Marathon
	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")

	return request, nil
}

var oneLogLineRegex = regexp.MustCompile(`(?m)^\s*`)

// oneLogLine removes indentation at the beginning of each line and
// escapes new line characters.
func oneLogLine(in []byte) []byte {
	return bytes.Replace(oneLogLineRegex.ReplaceAll(in, nil), []byte("\n"), []byte("\\n "), -1)
}
