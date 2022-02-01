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

func NewAlpacaAPI(clientId uint32, ip string, port int32) *ASCOMAlpacaAPIClient {
	client := ASCOMAlpacaAPIClient{
		client:        resty.New(),
		urlBase:       fmt.Sprintf("http://%s:%d", ip, port),
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
