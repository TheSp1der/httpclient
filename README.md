# RemoteHTTP

## Table of Contents

1. [Information](#information)
1. [Requirements](#requirements)
1. [Go doc](#go-documentation)

## Information

This library is provides a simple method for interacting with remote http servers.

## Requirements

This library does not have any non-standard requirements.

## Usage

``` go
import "github.com/TheSp1der/httpclient"
import "log"

func main() {
     x := httpclient.DefaultClient()
     x.SetHeader("Accept", "application/json")
     x.SetHeader("Content-Type", "application/json")
     x.SetBasicAuth("username", "password")
     x.SetPostData("{}")

     output, err := x.Get("http://example.com")
     if err != nil {
          log.Println(err.Error())
     }

     log.Println(string(output))
}
```

## Go Documentation

`import "github.com/TheSp1der/httpclient"`

* type HTTPClient
* func DefaultClient() *HTTPClient
* func NewClient(httpClient *http.Client) *HTTPClient
* func (c *HTTPClient) Get(url string) ([]byte, error)
* func (c *HTTPClient) Patch(url string) ([]byte, error)
* func (c *HTTPClient) Post(url string) ([]byte, error)
* func (c *HTTPClient) Put(url string) ([]byte, error)
* func (c *HTTPClient) SetBasicAuth(username, password string) *HTTPClient
* func (c *HTTPClient) SetHeader(label string, value string) *HTTPClient
* func (c *HTTPClient) SetPostData(data string) *HTTPClient

### Package files
[remotehttp.go](github.com/TheSp1der/httpclient/blob/master/remotehttp.go) 
### type HTTPClient

``` go
type HTTPClient struct {
    Client  *http.Client
    Data    *bytes.Buffer
    Headers map[string]string

    Username string
    Password string
}

```

HTTPClient is an interface for initializing the http client library.

### func DefaultClient

``` go
func DefaultClient() *HTTPClient
```

DefaultClient is a function for defining a basic HTTP client with standard timeouts.

### func NewClient

``` go
func NewClient(httpClient *http.Client) *HTTPClient
```

Create an HTTPClient with a user-provided net/http.Client

### func (\*HTTPClient) Get

``` go
func (c *HTTPClient) Get(url string) ([]byte, error)
```
Get calls the net.http GET operation

### func (\*HTTPClient) Patch

``` go
func (c *HTTPClient) Patch(url string) ([]byte, error)
```

Patch calls the net.http PATCH operation

### func (\*HTTPClient) Post

``` go
func (c *HTTPClient) Post(url string) ([]byte, error)
```

Post calls the net.http POST operation

### func (\*HTTPClient) Put

``` go
func (c *HTTPClient) Put(url string) ([]byte, error)
```

Put calls the net.http PUT operation

### func (\*HTTPClient) SetBasicAuth

``` go
func (c *HTTPClient) SetBasicAuth(username, password string) *HTTPClient
```

SetBasicAuth is a chaining function to set the username and password for basic authentication

### func (\*HTTPClient) SetHeader

``` go
func (c *HTTPClient) SetHeader(label string, value string) *HTTPClient
```

SetHeader is a chaining function to set arbitrary HTTP Headers

### func (\*HTTPClient) SetPostData

``` go
func (c *HTTPClient) SetPostData(data string) *HTTPClient
```

SetPostData is a chaining function to set POST/PUT/PATCH data
