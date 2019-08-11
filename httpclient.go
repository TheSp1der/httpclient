package httpclient

import (
	"bytes"
	"errors"
	"net"
	"strconv"
	"time"

	"io/ioutil"
	"net/http"
)

// HTTPClient is an interface for initializing the http client library.
type HTTPClient struct {
	Client  *http.Client
	Data    *bytes.Buffer
	Headers map[string]string

	Username string
	Password string
}

// DefaultClient is a function for defining a basic HTTP client with standard timeouts.
func DefaultClient() *HTTPClient {
	return &HTTPClient{
		Client: &http.Client{
			Timeout: 60 * time.Second,
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout: 5 * time.Second,
				}).Dial,
				TLSHandshakeTimeout: 5 * time.Second,
				IdleConnTimeout:     300 * time.Second,
			},
		},
	}
}

// NewClient Create an HTTPClient with a user-provided net/http.Client
func NewClient(httpClient *http.Client) *HTTPClient {
	return &HTTPClient{Client: httpClient}
}

// SetBasicAuth is a chaining function to set the username and password for basic
// authentication
func (c *HTTPClient) SetBasicAuth(username, password string) *HTTPClient {
	c.Username = username
	c.Password = password

	return c
}

// SetPostData is a chaining function to set POST/PUT/PATCH data
func (c *HTTPClient) SetPostData(data string) *HTTPClient {
	c.Data = bytes.NewBufferString(data)

	return c
}

// SetHeader is a chaining function to set arbitrary HTTP Headers
func (c *HTTPClient) SetHeader(label string, value string) *HTTPClient {
	if c.Headers == nil {
		c.Headers = map[string]string{}
	}

	c.Headers[label] = value

	return c
}

// Get calls the net.http GET operation
func (c *HTTPClient) Get(url string) ([]byte, error) {
	return c.do(url, http.MethodGet)
}

// Patch calls the net.http PATCH operation
func (c *HTTPClient) Patch(url string) ([]byte, error) {
	return c.do(url, http.MethodPatch)
}

// Post calls the net.http POST operation
func (c *HTTPClient) Post(url string) ([]byte, error) {
	return c.do(url, http.MethodPost)
}

// Put calls the net.http PUT operation
func (c *HTTPClient) Put(url string) ([]byte, error) {
	return c.do(url, http.MethodPut)
}

func (c *HTTPClient) do(url string, method string) ([]byte, error) {
	var (
		req    *http.Request
		res    *http.Response
		output []byte
		err    error
	)

	// NewRequest knows that c.data is typed *bytes.Buffer and will SEGFAULT
	// if c.data is nil. So we create a request using nil when c.data is nil
	if c.Data != nil {
		req, err = http.NewRequest(method, url, c.Data)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return nil, err
	}

	if (len(c.Username) > 0) && (len(c.Password) > 0) {
		req.SetBasicAuth(c.Username, c.Password)
	}

	if c.Headers != nil {
		for label, value := range c.Headers {
			req.Header.Set(label, value)
		}
	}

	if res, err = c.Client.Do(req); err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if output, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}

	// check status
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, errors.New("non-successful status code received [" + strconv.Itoa(res.StatusCode) + "]")
	}

	return output, nil
}
