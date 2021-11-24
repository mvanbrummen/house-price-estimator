package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2/clientcredentials"
)

type PropertyGateway struct {
	baseUrl string
	client  *resty.Client
}

func NewPropertyGateway(baseUrl string, clientId string, clientSecret string, client *resty.Client) *PropertyGateway {
	conf := &clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     fmt.Sprintf("%s/access/oauth/token", baseUrl),
	}

	ctx := context.Background()
	oauthClient := conf.Client(ctx)

	client.SetTransport(oauthClient.Transport)

	p := &PropertyGateway{
		baseUrl: baseUrl,
		client:  client,
	}

	return p
}

func (p *PropertyGateway) GetValuation(propertyId int) (*ValuationResponse, error) {
	resp, err := p.client.R().
		SetResult(&ValuationResponse{}).
		SetError(&ErrorResponse{}).
		Get(fmt.Sprintf("%s/avm/au/properties/%d/avm/intellival/consumer/current", p.baseUrl, propertyId))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Valuation returned %d status code: %s", resp.StatusCode(), getErrorMessage(resp)))
	}

	body := resp.Result().(*ValuationResponse)

	return body, nil
}

func (p *PropertyGateway) GetSuggestions(query string) (*SuggestResponse, error) {
	resp, err := p.client.R().
		SetQueryParam("q", query).
		SetQueryParam("suggestionTypes", "address").
		SetResult(&SuggestResponse{}).
		SetError(&ErrorResponse{}).
		Get(fmt.Sprintf("%s/property/au/v2/suggest.json", p.baseUrl))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Suggest returned %d status code: %s", resp.StatusCode(), getErrorMessage(resp)))
	}

	body := resp.Result().(*SuggestResponse)

	return body, nil
}

func (p *PropertyGateway) GetAttributes(id int) (*AttributesResponse, error) {
	resp, err := p.client.R().
		SetResult(&AttributesResponse{}).
		SetError(&ErrorResponse{}).
		Get(fmt.Sprintf("%s/property-details/au/properties/%d/attributes/core", p.baseUrl, id))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Attributes returned %d status code: %s", resp.StatusCode(), getErrorMessage(resp)))
	}

	body := resp.Result().(*AttributesResponse)

	return body, nil
}

func (p *PropertyGateway) GetLastSale(id int) (*LastSaleResponse, error) {
	resp, err := p.client.R().
		SetResult(&LastSaleResponse{}).
		SetError(&ErrorResponse{}).
		Get(fmt.Sprintf("%s/property-details/au/properties/%d/sales/last", p.baseUrl, id))

	if err != nil {
		return nil, err
	}

	log.Println(resp)

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Last Sale returned %d status code: %s", resp.StatusCode(), getErrorMessage(resp)))
	}

	body := resp.Result().(*LastSaleResponse)

	return body, nil
}

func (p *PropertyGateway) GetImagery(id int) (*ImageryResponse, error) {
	resp, err := p.client.R().
		SetResult(&ImageryResponse{}).
		SetError(&ErrorResponse{}).
		Get(fmt.Sprintf("%s/property-details/au/properties/%d/images", p.baseUrl, id))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Imagery returned %d status code: %s", resp.StatusCode(), getErrorMessage(resp)))
	}

	body := resp.Result().(*ImageryResponse)

	return body, nil
}

func getErrorMessage(resp *resty.Response) string {
	body := resp.Error().(*ErrorResponse)

	errorMessage := ""
	for _, e := range body.Errors {
		errorMessage += " " + e.Msg
	}
	return errorMessage
}
