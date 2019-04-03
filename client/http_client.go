package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func NewHTTP(target string, method string) Client {
	return &httpClient{target: target, method: method}
}

type httpClient struct {
	target string
	method string
}

func (c httpClient) GetVersion(ctx context.Context) (*GetVersionReply, error) {
	var netTransport = &http.Transport{
		TLSHandshakeTimeout: 10 * time.Second,
	}
	client := &http.Client{
		Timeout:   time.Second * 60,
		Transport: netTransport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.target, c.method), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	out := new(GetVersionReply)
	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return nil, err
	}

	return out, nil
}
