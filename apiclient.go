package rebotapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var kJson = jsoniter.ConfigCompatibleWithStandardLibrary

type apiClient struct {
	Hostport string
	Client   *http.Client
}

func NewApiclient(hostport string) (c *apiClient) {
	c = &apiClient{
		Hostport: hostport,
		Client: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 64,
			},
			Timeout: 1 * time.Second,
		},
	}
	return
}

func (c *apiClient) Post(category, model, action string, param interface{}, resp interface{}) (err error) {
	body, _ := kJson.Marshal(param)
	url := c.buildUrl(category, model, action)
	httpReq, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	httpReq.Header.Add("Content-Type", "application/json; charset=utf-8")
	httpReq.Header.Add("Accept", "application/json")
	httpReq.Header.Add("Connection", "keep-alive")

	httpResp, err := c.Client.Do(httpReq)
	if err != nil {
		return
	}
	if httpResp.Body == nil {
		err = fmt.Errorf("resp body is empty")
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != 200 {
		err = fmt.Errorf("http status code: %v", httpResp.StatusCode)
		return
	}

	payload, err := ioutil.ReadAll(httpResp.Body)
	kJson.Unmarshal(payload, resp)

	return
}

func (c *apiClient) Get(category, model, action string, resp interface{}) (err error) {
	url := c.buildUrl(category, model, action)
	httpReq, err := http.NewRequest("GET", url, nil)
	httpReq.Header.Add("Content-Type", "application/json; charset=utf-8")
	httpReq.Header.Add("Accept", "application/json")
	httpReq.Header.Add("Connection", "keep-alive")

	httpResp, err := c.Client.Do(httpReq)
	if err != nil {
		return
	}
	if httpResp.Body == nil {
		err = fmt.Errorf("resp body is empty")
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != 200 {
		err = fmt.Errorf("http status code: %v", httpResp.StatusCode)
		return
	}

	payload, err := ioutil.ReadAll(httpResp.Body)
	kJson.Unmarshal(payload, resp)

	return
}

func (c *apiClient) buildUrl(category, model, action string) (url string) {
	url = c.Hostport
	if category != "" {
		url = url + "/" + category
	}
	if model != "" {
		url = url + "/" + model
	}
	if action != "" {
		url = url + "/" + action
	}
	return
}
