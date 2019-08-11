package httpclient

import (
	"fmt"
	"testing"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

type Data struct {
	Greeting string            `json:"greeting"`
	Headers  map[string]string `json:"headers"`
	Method   string            `json:"method"`
	Username string            `json:"username"`
	Password string            `json:"password"`
	PostData string            `json:"postdata"`
}

var (
	greeting    = "Hello world"
	postData    = "Test data"
	authUser    = "testuser"
	authPass    = "testpass"
	headerLabel = "Test-Header"
	headerValue = "Test-Value"
)

func httpTestHandler(w http.ResponseWriter, r *http.Request) {
	var (
		b    []byte
		user string
		pass string
		body []byte
	)

	data := Data{
		Greeting: greeting,
		Headers:  map[string]string{},
		Method:   r.Method,
	}

	user, pass, ok := r.BasicAuth()
	if ok {
		data.Username = user
		data.Password = pass
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "ioutil.ReadAll failed")
	}
	data.PostData = string(body)

	for h := range r.Header {
		data.Headers[h] = r.Header.Get(h)
	}

	b, err = json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Fprint(w, "Json marshal failed somehow")
	}
	fmt.Fprint(w, string(b))
}

func checkMethod(t *testing.T, data Data, method string) {
	if data.Method != method {
		t.Errorf("data.Method(%s) != method(%s)", data.Method, method)
	}
	t.Log("checkMethod() success")
}

func checkGreeting(t *testing.T, data Data) {
	if data.Greeting != greeting {
		t.Errorf("data.Greeting(%s) != greeting(%s)", data.Greeting, greeting)
	}
	t.Log("checkGreeting() success")
}

func checkBasicAuth(t *testing.T, data Data) {
	if data.Username != authUser {
		t.Errorf("data.Username(%s) != authUser(%s)", data.Username, authUser)
	}
	if data.Password != authPass {
		t.Errorf("data.Password(%s) != authPass(%s)", data.Password, authPass)
	}
	t.Log("checkBasicAuth() success")
}

func checkPostData(t *testing.T, data Data) {
	if data.PostData != postData {
		t.Errorf("data.PostData(%s) != postData(%s)", data.PostData, postData)
	}
	t.Log("checkPostData() success")
}

func TestGet(t *testing.T) {
	var data Data

	ts := httptest.NewServer(http.HandlerFunc(httpTestHandler))
	defer ts.Close()

	output, err := DefaultClient().Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if err = json.Unmarshal(output, &data); err != nil {
		t.Error(err)
	}

	checkMethod(t, data, http.MethodGet)
	checkGreeting(t, data)
}

func TestGetAuth(t *testing.T) {
	var data Data

	ts := httptest.NewServer(http.HandlerFunc(httpTestHandler))
	defer ts.Close()

	output, err := DefaultClient().SetBasicAuth(authUser, authPass).Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if err = json.Unmarshal(output, &data); err != nil {
		t.Error(err)
	}

	checkMethod(t, data, http.MethodGet)
	checkGreeting(t, data)
	checkBasicAuth(t, data)
}

func TestPut(t *testing.T) {
	var data Data

	ts := httptest.NewServer(http.HandlerFunc(httpTestHandler))
	defer ts.Close()

	output, err := DefaultClient().SetPostData(postData).Put(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if err = json.Unmarshal(output, &data); err != nil {
		t.Error(err)
	}

	checkMethod(t, data, http.MethodPut)
	checkGreeting(t, data)
	checkPostData(t, data)
}

func TestPutAuth(t *testing.T) {
	var data Data

	ts := httptest.NewServer(http.HandlerFunc(httpTestHandler))
	defer ts.Close()

	output, err := DefaultClient().SetBasicAuth(authUser, authPass).SetPostData(postData).Put(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if err = json.Unmarshal(output, &data); err != nil {
		t.Error(err)
	}

	checkMethod(t, data, http.MethodPut)
	checkGreeting(t, data)
	checkBasicAuth(t, data)
	checkPostData(t, data)
}

func TestPost(t *testing.T) {
	var data Data

	ts := httptest.NewServer(http.HandlerFunc(httpTestHandler))
	defer ts.Close()

	output, err := DefaultClient().SetPostData(postData).Post(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if err = json.Unmarshal(output, &data); err != nil {
		t.Error(err)
	}

	checkMethod(t, data, http.MethodPost)
	checkGreeting(t, data)
	checkPostData(t, data)
}

func TestPostAuth(t *testing.T) {
	var data Data

	ts := httptest.NewServer(http.HandlerFunc(httpTestHandler))
	defer ts.Close()

	output, err := DefaultClient().SetBasicAuth(authUser, authPass).SetPostData(postData).Post(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if err = json.Unmarshal(output, &data); err != nil {
		t.Error(err)
	}

	checkMethod(t, data, http.MethodPost)
	checkGreeting(t, data)
	checkBasicAuth(t, data)
	checkPostData(t, data)
}

func TestPatch(t *testing.T) {
	var data Data

	ts := httptest.NewServer(http.HandlerFunc(httpTestHandler))
	defer ts.Close()

	output, err := DefaultClient().SetPostData(postData).Patch(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if err = json.Unmarshal(output, &data); err != nil {
		t.Error(err)
	}

	checkMethod(t, data, http.MethodPatch)
	checkGreeting(t, data)
	checkPostData(t, data)
}

func TestPatchAuth(t *testing.T) {
	var data Data

	ts := httptest.NewServer(http.HandlerFunc(httpTestHandler))
	defer ts.Close()

	output, err := DefaultClient().SetBasicAuth(authUser, authPass).SetPostData(postData).Patch(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if err = json.Unmarshal(output, &data); err != nil {
		t.Error(err)
	}

	checkMethod(t, data, http.MethodPatch)
	checkGreeting(t, data)
	checkBasicAuth(t, data)
	checkPostData(t, data)
}

func TestSetHeader(t *testing.T) {
	var data Data

	ts := httptest.NewServer(http.HandlerFunc(httpTestHandler))
	defer ts.Close()

	output, err := DefaultClient().SetHeader(headerLabel, headerValue).Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if err = json.Unmarshal(output, &data); err != nil {
		t.Error(err)
	}

	checkMethod(t, data, http.MethodGet)
	checkGreeting(t, data)
	if data.Headers[headerLabel] != headerValue {
		t.Errorf("SetHeader values not set in header: %+v", data.Headers)
	}
}
