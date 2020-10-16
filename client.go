package talend

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const DefaultRestUrl string = "https://api.eu.cloud.talend.com/tmc/v2.2"

type Client struct {
	HttpClient *http.Client
	ApiKey     string
	Host       string
	Base       string
	Proxy      string
}

func NewClient(apiKey string, proxy string) (*Client, error) {

	if len(proxy) > 0 {

		proxyURL, err := url.Parse(proxy)
		if err != nil {
			return nil, err
		}

		transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
		client := &http.Client{Transport: transport}

		return &Client{
			HttpClient: client,
			ApiKey:     apiKey,
			Proxy:      proxy,
		}, nil
	} else {
		return &Client{
			HttpClient: http.DefaultClient,
			ApiKey:     apiKey,
		}, nil
	}
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent || res.StatusCode == http.StatusCreated {
		return body, err
	} else {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
}
