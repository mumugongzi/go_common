package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type HttpClient struct {
	client        *http.Client
	isHttps       bool // 默认https
}

func NewHttpClient() *HttpClient {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	return &HttpClient{client: client}
}

func (c *HttpClient) Get(ctx context.Context, url string) (string, error) {
	return c.Request(ctx, url, nil, "", "", "GET", nil)
}

func (c *HttpClient) GetWithHeader(ctx context.Context, url string, header map[string]string) (string, error) {
	return c.Request(ctx, url, header, "", "", "GET", nil)
}

func (c *HttpClient) GetWithBasicAuth(ctx context.Context, url, username, password string) (string, error) {
	return c.Request(ctx, url, nil, username, password, "GET", nil)
}

func (c *HttpClient) Post(ctx context.Context, url string, postObject interface{}, header map[string]string) (string, error) {
	bodyBytes, err := json.Marshal(postObject)
	if err != nil {
		return "", err
	}
	bodyReader := bytes.NewReader(bodyBytes)

	return c.Request(ctx, url, header, "", "", "POST", bodyReader)
}

func (c *HttpClient) Put(ctx context.Context, url string, header map[string]string) (string, error) {
	return c.Request(ctx, url, header, "", "", "PUT", nil)
}

func (c *HttpClient) Patch(ctx context.Context, url string, header map[string]string) (string, error) {
	return c.Request(ctx, url, header, "", "", "PATCH", nil)
}

func (c *HttpClient) Delete(ctx context.Context, url string, header map[string]string) (string, error) {
	return c.Request(ctx, url, header, "", "", "DELETE", nil)
}

func (c *HttpClient) PutWithBody(ctx context.Context, url string, header map[string]string, bodyObject interface{}) (string, error) {
	bodyBytes, err := json.Marshal(bodyObject)
	if err != nil {
		return "", err
	}
	bodyReader := bytes.NewReader(bodyBytes)
	return c.Request(ctx, url, header, "", "", "PUT", bodyReader)
}

func (c *HttpClient) PatchWithBody(ctx context.Context, url string, header map[string]string, bodyObject interface{}) (string, error) {
	bodyBytes, err := json.Marshal(bodyObject)
	if err != nil {
		return "", err
	}
	bodyReader := bytes.NewReader(bodyBytes)
	return c.Request(ctx, url, header, "", "", "PATCH", bodyReader)
}

func (c *HttpClient) Request(ctx context.Context, url string, header map[string]string, username, password, method string, body io.Reader) (string, error) {

	response, err := c.RequestGetResponse(ctx, url, header, username, password, method, body)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	s := strings.TrimSpace(string(content))

	return s, nil
}

func (c *HttpClient) RequestGetResponse(ctx context.Context, url string, header map[string]string, username, password, method string, body io.Reader) (*http.Response, error) {
	req, err := c.NewRequest(ctx, method, url, header, body)
	if err != nil {
		return nil, err
	}

	// 增加basic http auth
	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	response, err := c.client.Do(req)
	return response, err
}

func (c *HttpClient) NewRequest(ctx context.Context, method, url string, header map[string]string, body io.Reader) (*http.Request, error) {
	rq, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	rq.Header.Set("Content-Type", "application/json")

	addHeader := func(data map[string]string) {
		if len(data) == 0 {
			return
		}
		for key, value := range data {
			rq.Header.Set(key, value)
		}
	}
	// 增加header选项
	addHeader(header)
	return rq, err
}
