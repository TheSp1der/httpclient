package remotehttp

import (
	"errors"
	"net"
	"strconv"

	"io/ioutil"
	"net/http"

	e "github.com/TheSp1der/goerror"
)

// GetAuth http get from remote webserver with authentication
func (config WebConfig) GetAuth(url string, headers Headers, user string, pass string) ([]byte, error) {
	var (
		err    error
		client http.Client
		req    *http.Request
		res    *http.Response
		output []byte
	)

	if config.LogLevel >= 75 {
		e.Info(errors.New("http get: " + url))
	}

	// set timeouts
	client = http.Client{
		Timeout: config.RxTimeout,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: config.ConTimeout,
			}).Dial,
			TLSHandshakeTimeout: config.SSLHandshakeTimeout,
		},
	}

	// setup request
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return output, err
	}

	// setup headers
	if len(headers) > 0 {
		for _, header := range headers {
			req.Header.Set(header.Label, header.Value)
		}
	}

	// add credentials to request
	if len(user) == 0 || len(pass) == 0 {
		return output, errors.New("no credentials provided for http get with authentication")
	}
	req.SetBasicAuth(user, pass)

	// perform the request
	if res, err = client.Do(req); err != nil {
		return output, err
	}

	// close the connection upon function closure
	defer res.Body.Close()

	// extract response body
	if output, err = ioutil.ReadAll(res.Body); err != nil {
		return output, err
	}

	// check status
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return output, errors.New("non-successful status code received [" + strconv.Itoa(res.StatusCode) + "]")
	}

	return output, nil
}
