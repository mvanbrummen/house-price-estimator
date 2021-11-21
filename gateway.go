package main

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type PropertyGateway struct {
	baseUrl      string
	client       *resty.Client
	accessToken  string
	clientId     string
	clientSecret string
}

type AccessResponse struct {
	AccessToken string `json:"access_token,omitempty"`
}

func NewPropertyGateway(baseUrl string, clientId string, clientSecret string, client *resty.Client) *PropertyGateway {
	return &PropertyGateway{
		baseUrl:      baseUrl,
		client:       client,
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

func (p *PropertyGateway) GetSuggestions(query string) {

}

func (p *PropertyGateway) GetAccessToken() {
	resp, err := p.client.R().SetResult(&AccessResponse{}).
		Get(fmt.Sprintf("%s/access/oauth/token?grant_type=client_credentials&client_id=%s&client_secret=%s", p.baseUrl, p.clientId, p.clientSecret))

	if err != nil {
		panic(err)
	}

	body := resp.Result().(*AccessResponse)

	p.accessToken = body.AccessToken

	log.Println(p.accessToken)
}
