// Copyright (c) 2021-2022 observerly Ltd (michael@observerly.com), All rights reserved.

package alpaca

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ASCOMAlpacaAPIClient struct {
	client        *resty.Client
	urlBase       string
	clientId      uint32
	transactionId uint32
	errorNumber   int
	errorMessage  string
}

func NewAlpacaAPI(clientId uint32, secure bool, domain string, ip string, port int32) *ASCOMAlpacaAPIClient {
	var protocol string = "https"

	if !secure {
		protocol = "http"
	}

	var urlBase string = fmt.Sprintf("%s://%s:%d", protocol, ip, port)

	if port == -1 && len(domain) > 0 {
		urlBase = fmt.Sprintf("%s://%s", protocol, domain)
	}

	client := ASCOMAlpacaAPIClient{
		client:        resty.New(),
		urlBase:       urlBase,
		clientId:      clientId,
		transactionId: 0,
	}

	return &client
}

/*
	getQueryString()

	Many ASCOM methods require parameter values. All methods
	that use the HTTP GET verb should include parameters as
	query string name-value pairs.
*/
func (a *ASCOMAlpacaAPIClient) getQueryString() string {
	return fmt.Sprintf("ClientID=%d&ClientTransactionID=%d", a.clientId, a.transactionId)
}

/*
	getEndpoint()

	Alpaca Device API URLs are of the form http(s)://host:port/path
	where path comprises "/api/v1/" followed by one of the ASCOM
	method from https://ascom-standards.org/api/
*/
func (a *ASCOMAlpacaAPIClient) getEndpoint(deviceType string, deviceNumber uint, method string) string {
	return fmt.Sprintf("%s/api/v1/%s/%d/%s", a.urlBase, deviceType, deviceNumber, method)
}

type stringResponse struct {
	Value               string `json:"Value"`
	ClientTransactionID uint32 `json:"ClientTransactionID"`
	ServerTransactionID uint32 `json:"ServerTransactionID"`
	ErrorNumber         int32  `json:"ErrorNumber"`
	ErrorMessage        string `json:"ErrorMessage"`
}

/*
	GetStringResponse()

	Global public method to work with calls returning stringResponse
*/
func (a *ASCOMAlpacaAPIClient) GetStringResponse(deviceType string, deviceNumber uint, method string) (string, error) {
	// Build the ASCOM endpoint:
	url := a.getEndpoint(deviceType, deviceNumber, method)

	// Setup the resty client:
	resp, err := a.client.R().SetResult(&stringResponse{}).SetQueryString(a.getQueryString()).SetHeader("Accept", "application/json").Get(url)

	if err != nil {
		return "", err
	}

	// If the response object has a REST error:
	if resp.IsError() {
		a.errorNumber = resp.StatusCode()
		a.errorMessage = resp.String()
	}

	// Return the result:
	result := (resp.Result().(*stringResponse))

	return result.Value, nil
}

type stringlistResponse struct {
	Value               []string `json:"Value"`
	ClientTransactionID uint32   `json:"ClientTransactionID"`
	ServerTransactionID uint32   `json:"ServerTransactionID"`
	ErrorNumber         int32    `json:"ErrorNumber"`
	ErrorMessage        string   `json:"ErrorMessage"`
}

/*
	GetStringListResponse()

	Global public method to work with calls returning stringListResponse
*/
func (a *ASCOMAlpacaAPIClient) GetStringListResponse(deviceType string, deviceNumber uint, method string) ([]string, error) {
	// Build the ASCOM endpoint:
	url := a.getEndpoint(deviceType, deviceNumber, method)

	resp, err := a.client.R().SetResult(&stringlistResponse{}).SetQueryString(a.getQueryString()).SetHeader("Accept", "application/json").Get(url)

	if err != nil {
		return []string{""}, err
	}

	// If the response object has a REST error:
	if resp.IsError() {
		a.errorNumber = resp.StatusCode()
		a.errorMessage = resp.String()
	}

	// Return the result:
	result := (resp.Result().(*stringlistResponse))

	return result.Value, nil
}
