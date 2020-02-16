package client

import (
	"compress/flate"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const eftDomain string = "prod.escapefromtarkov.com"

var eftHeaders http.Header = func() http.Header {
	headers := make(http.Header)
	headers.Add("Host", eftDomain)
	headers.Add("Accept", "*/*")
	headers.Add("Content-Type", "application/json")
	headers.Add("User-Agent", "UnityPlayer/5.6.5p3")
	headers.Add("Connection", "Keep-Alive")
	headers.Add("X-Unity-Version", "5.6.5p3")

	return headers
}()

type TarkovClient struct {
	client    *http.Client
	sessionID string
}

func NewClient(client *http.Client, sessionID string) TarkovClient {
	if client == nil {
		client = http.DefaultClient
	}

	return TarkovClient{
		client,
		sessionID,
	}
}

func (client *TarkovClient) Get(path string, ret interface{}) error {
	endpoint, err := url.Parse(fmt.Sprintf("https://%s/%s", eftDomain, path))
	if err != nil {
		return err
	}

	request := http.Request{
		Method: "GET",
		URL:    endpoint,
		Header: eftHeaders,
	}
	addAuthCookie(&request, client.sessionID)

	response, err := client.client.Do(&request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// ! Apparently, the responses are compressed
	// ! However, no header is sent to let us know that
	// ! Therefore, we must manually initiate the deflate
	deflatedBody := flate.NewReader(response.Body)
	defer deflatedBody.Close()

	data, err := ioutil.ReadAll(deflatedBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, ret)
}

func addAuthCookie(request *http.Request, sessionID string) {
	cookie := http.Cookie{
		Name:  "PHPSESSID",
		Value: sessionID,
	}
	request.AddCookie(&cookie)
}
