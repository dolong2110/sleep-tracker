package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mindx/pkg/zapx"
	"net"
	"net/http"
)

type ExternalHttpRequestResponse struct {
	Header     http.Header
	Body       []byte
	StatusCode int
}

func Request(method, url, token string, body interface{}) (*ExternalHttpRequestResponse, error) {
	httpRequest, err := createHttpRequest(method, url, token, body)
	if err != nil {
		return nil, err
	}

	_, err = net.Dial("tcp", "localhost:8000")
	if err != nil {
		zapx.Error(context.TODO(), "Failed to connect to dial tcp", err)
		return nil, err
	}

	resp, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		zapx.Error(context.TODO(), "Failed to get the response", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			zapx.Error(context.TODO(), "Failed to sen request", err)
		}
	}(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err = errors.New(fmt.Sprintf("Auth0 request error: %s, message: %s", resp.Status, resp.Body))
		zapx.Error(context.TODO(), "Unsuccessful request", err)
		return nil, err
	}

	respHeader := resp.Header
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		zapx.Error(context.TODO(), "Failed to parse request body", err)
		return nil, err
	}

	respStatusCode := resp.StatusCode

	result := ExternalHttpRequestResponse{
		Header:     respHeader,
		Body:       respBody,
		StatusCode: respStatusCode,
	}

	return &result, nil
}

func createHttpRequest(method, url, token string, body interface{}) (*http.Request, error) {
	payloadBuf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(payloadBuf).Encode(body)
		if err != nil {
			zapx.Error(context.TODO(), "Failed to encode the request body in json format", err)
			return nil, err
		}
	}

	httpRequest, err := http.NewRequest(method, url, payloadBuf)
	if err != nil {
		zapx.Error(context.TODO(), "Failed to make the request", err)
		return nil, err
	}

	if body != nil {
		httpRequest.Header.Add(HeaderContentType, ContentTypeApplicationJSON)
	}
	if token != "" {
		httpRequest.Header.Set(HeaderAuthorization, TokenTypeBearer+" "+token)
	}

	return httpRequest, nil
}
