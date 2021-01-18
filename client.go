package talend

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const defaultRestURL string = "https://api.eu.cloud.talend.com/tmc/v2.2"

// Client contains the configuration for the API client
type Client struct {
	HTTPClient *http.Client
	APIKey     string
	Host       string
	Base       string
	Proxy      string
}

// NewClient creates a new API client
func NewClient(apiKey string, proxy string) (*Client, error) {

	if len(proxy) > 0 {

		proxyURL, err := url.Parse(proxy)
		if err != nil {
			return nil, err
		}

		transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
		client := &http.Client{Transport: transport}

		return &Client{
			HTTPClient: client,
			APIKey:     apiKey,
			Proxy:      proxy,
		}, nil
	}
	return &Client{
		HTTPClient: http.DefaultClient,
		APIKey:     apiKey,
	}, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	res, err := c.HTTPClient.Do(req)
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
	}

	return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
}
