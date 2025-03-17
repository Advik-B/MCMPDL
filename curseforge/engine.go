package curseforge

import (
	"io"
	"net/http"
	"net/url"
)

type API struct {
	apiKey     string
	user_agent string
	baseUrl    string
}

func NewAPI(apiKey string) *API {
	return &API{
		apiKey:     apiKey,
		user_agent: "",
		baseUrl:    "https://api.curseforge.com/",
	}
}

func (self *API) SetUserAgent(user_agent string) {
	self.user_agent = user_agent
}

func (self *API) Fetch(path string) (data []byte, err error) {
	url_to_fetch, err := url.JoinPath(self.baseUrl, path)
	if err != nil {
		return nil, err
	}
	return self.FetchRaw(url_to_fetch)
}

func (self *API) FetchRaw(url string) (data []byte, err error) {
	res, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if self.user_agent != "" {
		res.Header.Set("User-Agent", self.user_agent)
	}
	res.Header.Set("Accept", "application/json")
	res.Header.Set("x-api-key", self.apiKey)

	client := &http.Client{}

	resp, err := client.Do(res)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
