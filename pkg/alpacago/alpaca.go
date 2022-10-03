// Copyright (c) 2021-2022 observerly Ltd (michael@observerly.com), All rights reserved.
package alpacago

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/go-resty/resty/v2"
)

type Direction int32

const (
	North Direction = iota
	South
	East
	West
)

type ASCOMAlpacaAPIClient struct {
	Client        *resty.Client
	UrlBase       string
	ClientId      uint32
	TransactionId uint32
	ErrorNumber   int
	ErrorMessage  string
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
		Client:        resty.New(),
		UrlBase:       urlBase,
		ClientId:      clientId,
		TransactionId: 0,
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
	return fmt.Sprintf("ClientID=%d&ClientTransactionID=%d", a.ClientId, a.TransactionId)
}

/*
	getEndpoint()

	Alpaca Device API URLs are of the form http(s)://host:port/path
	where path comprises "/api/v1/" followed by one of the ASCOM
	method from https://ascom-standards.org/api/
*/
func (a *ASCOMAlpacaAPIClient) getEndpoint(deviceType string, deviceNumber uint, method string) string {
	return fmt.Sprintf("%s/api/v1/%s/%d/%s", a.UrlBase, deviceType, deviceNumber, method)
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
	resp, err := a.Client.R().SetResult(&stringResponse{}).SetQueryString(a.getQueryString()).SetHeader("Accept", "application/json").Get(url)

	if err != nil {
		return "", err
	}

	// If the response object has a REST error:
	if resp.IsError() {
		a.ErrorNumber = resp.StatusCode()
		a.ErrorMessage = resp.String()
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

	resp, err := a.Client.R().SetResult(&stringlistResponse{}).SetQueryString(a.getQueryString()).SetHeader("Accept", "application/json").Get(url)

	if err != nil {
		return []string{""}, err
	}

	// If the response object has a REST error:
	if resp.IsError() {
		a.ErrorNumber = resp.StatusCode()
		a.ErrorMessage = resp.String()
	}

	// Return the result:
	result := (resp.Result().(*stringlistResponse))

	return result.Value, nil
}

type booleanResponse struct {
	Value               bool   `json:"Value"`
	ClientTransactionID uint32 `json:"ClientTransactionID"`
	ServerTransactionID uint32 `json:"ServerTransactionID"`
	ErrorNumber         int32  `json:"ErrorNumber"`
	ErrorMessage        string `json:"ErrorMessage"`
}

/*
	GetBooleanResponse()

	Global public method to work with calls returning booleanResponse
*/
func (a *ASCOMAlpacaAPIClient) GetBooleanResponse(deviceType string, deviceNumber uint, method string) (bool, error) {
	// Build the ASCOM endpoint:
	url := a.getEndpoint(deviceType, deviceNumber, method)

	resp, err := a.Client.R().SetResult(&booleanResponse{}).SetQueryString(a.getQueryString()).SetHeader("Accept", "application/json").Get(url)

	if err != nil {
		return false, err
	}

	// If the response object has a REST error:
	if resp.IsError() {
		a.ErrorNumber = resp.StatusCode()
		a.ErrorMessage = resp.String()
	}

	// Return the result:
	result := (resp.Result().(*booleanResponse))

	return result.Value, nil
}

type float64Response struct {
	Value               float64 `json:"Value"`
	ClientTransactionID int32   `json:"ClientTransactionID"`
	ServerTransactionID int32   `json:"ServerTransactionID"`
	ErrorNumber         int32   `json:"ErrorNumber"`
	ErrorMessage        string  `json:"ErrorMessage"`
}

/*
	GetFloat64Response()

	Global public method to work with calls returning float64Response
*/
func (a *ASCOMAlpacaAPIClient) GetFloat64Response(deviceType string, deviceNumber uint, method string) (float64, error) {
	// Build the ASCOM endpoint:
	url := a.getEndpoint(deviceType, deviceNumber, method)

	resp, err := a.Client.R().SetResult(&float64Response{}).SetQueryString(a.getQueryString()).SetHeader("Accept", "application/json").Get(url)

	if err != nil {
		return 0, err
	}

	// If the response object has a REST error:
	if resp.IsError() {
		a.ErrorNumber = resp.StatusCode()
		a.ErrorMessage = resp.String()
	}

	// Return the result:
	result := (resp.Result().(*float64Response))

	return result.Value, nil
}

type int32Response struct {
	Value               int32  `json:"Value"`
	ClientTransactionID uint32 `json:"ClientTransactionID"`
	ServerTransactionID uint32 `json:"ServerTransactionID"`
	ErrorNumber         int32  `json:"ErrorNumber"`
	ErrorMessage        string `json:"ErrorMessage"`
}

/*
	GetInt32Response()

	Global public method to work with calls returning int32Response
*/
func (a *ASCOMAlpacaAPIClient) GetInt32Response(deviceType string, deviceNumber uint, method string) (int32, error) {
	// Build the ASCOM endpoint:
	url := a.getEndpoint(deviceType, deviceNumber, method)

	resp, err := a.Client.R().SetResult(&int32Response{}).SetQueryString(a.getQueryString()).SetHeader("Accept", "application/json").Get(url)

	if err != nil {
		return 0, err
	}

	// If the response object has a REST error:
	if resp.IsError() {
		a.ErrorNumber = resp.StatusCode()
		a.ErrorMessage = resp.String()
	}

	// Return the result:
	result := (resp.Result().(*int32Response))

	return result.Value, nil
}

type uint32listResponse struct {
	Value               []uint32 `json:"Value"`
	ClientTransactionID int32    `json:"ClientTransactionID"`
	ServerTransactionID int32    `json:"ServerTransactionID"`
	ErrorNumber         int32    `json:"ErrorNumber"`
	ErrorMessage        string   `json:"ErrorMessage"`
}

/*
	GetInt32ListResponse()

	Global public method to work with calls returning int32listResponse
*/
func (a *ASCOMAlpacaAPIClient) GetUInt32ListResponse(deviceType string, deviceNumber uint, method string) ([]uint32, error) {
	// Build the ASCOM endpoint:
	url := a.getEndpoint(deviceType, deviceNumber, method)

	resp, err := a.Client.R().SetResult(&uint32listResponse{}).SetQueryString(a.getQueryString()).SetHeader("Accept", "application/json").Get(url)

	if err != nil {
		return []uint32{}, err
	}

	// If the response object has a REST error:
	if resp.IsError() {
		a.ErrorNumber = resp.StatusCode()
		a.ErrorMessage = resp.String()
	}

	// Return the result:
	result := (resp.Result().(*uint32listResponse))

	return result.Value, nil
}

type uint32Rank2ArrayResponse struct {
	Value               [][]uint32 `json:"Value"`
	Rank                uint32     `json:"Rank"`
	ClientTransactionID int32      `json:"ClientTransactionID"`
	ServerTransactionID int32      `json:"ServerTransactionID"`
	ErrorNumber         int32      `json:"ErrorNumber"`
	ErrorMessage        string     `json:"ErrorMessage"`
}

/*
	GetUInt32RankArrayResponse()

	Global public method to work with calls returning a uint32Rank2ArrayResponse
*/
func (a *ASCOMAlpacaAPIClient) GetUInt32Rank2ArrayResponse(deviceType string, deviceNumber uint, method string) ([][]uint32, uint32, error) {
	// Build the ASCOM endpoint:
	url := a.getEndpoint(deviceType, deviceNumber, method)

	resp, err := a.Client.R().SetResult(&uint32Rank2ArrayResponse{}).SetQueryString(a.getQueryString()).SetHeader("Accept", "application/json").Get(url)

	if err != nil {
		return [][]uint32{}, 0, err
	}

	// If the response object has a REST error:
	if resp.IsError() {
		a.ErrorNumber = resp.StatusCode()
		a.ErrorMessage = resp.String()
	}

	// Return the result:
	result := (resp.Result().(*uint32Rank2ArrayResponse))

	return result.Value, result.Rank, nil
}

type putResponse struct {
	ClientTransactionID uint32 `json:"ClientTransactionID"`
	ServerTransactionID uint32 `json:"ServerTransactionID"`
	ErrorNumber         int32  `json:"ErrorNumber"`
	ErrorMessage        string `json:"ErrorMessage"`
}

func (a *ASCOMAlpacaAPIClient) Put(deviceType string, deviceNumber uint, method string, form map[string]string) error {
	// Build the ASCOM endpoint:
	url := a.getEndpoint(deviceType, deviceNumber, method)

	resp, err := a.Client.R().SetHeader("Content-Type", "application/x-www-form-urlencoded").SetResult(&putResponse{}).SetHeader("Accept", "application/json").SetFormData(form).Put(url)

	if err != nil {
		return err
	}

	// If the response object has a REST error:
	if resp.IsError() {
		a.ErrorNumber = resp.StatusCode()
		a.ErrorMessage = resp.String()
	}

	// Return the result:
	result := (resp.Result().(*putResponse))

	if result.ErrorNumber != 0 {
		return fmt.Errorf("%d: %s", result.ErrorNumber, result.ErrorMessage)
	}

	log.Debugf("%v", result)

	return nil
}

/*
	IsConnected() common method to all ASCOM Alpaca compliant devices

	@returns the connection status of the device
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__connected
*/
func (a *ASCOMAlpacaAPIClient) IsConnected(deviceType string, deviceNumber uint) (bool, error) {
	return a.GetBooleanResponse(deviceType, deviceNumber, "connected")
}

/*
	GetDescription() common method to all ASCOM Alpaca compliant devices

	@returns the description of the device
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__description
*/
func (a *ASCOMAlpacaAPIClient) GetDescription(deviceType string, deviceNumber uint) (string, error) {
	return a.GetStringResponse(deviceType, deviceNumber, "description")
}

/*
	GetDriverInfo() common method to all ASCOM Alpaca compliant devices

	@returns the description of the driver
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__driverinfo
*/
func (a *ASCOMAlpacaAPIClient) GetDriverInfo(deviceType string, deviceNumber uint) (string, error) {
	return a.GetStringResponse(deviceType, deviceNumber, "driverinfo")
}

/*
	GetDriverVersion() common method to all ASCOM Alpaca compliant devices

	@returns a string containing only the major and minor version of the driver.
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__driverversion
*/
func (a *ASCOMAlpacaAPIClient) GetDriverVersion(deviceType string, deviceNumber uint) (string, error) {
	return a.GetStringResponse(deviceType, deviceNumber, "driverversion")
}

/*
	GetInterfaceVersion() common method to all ASCOM Alpaca compliant devices

	@returns the version of the ASCOM device interface contract to which
	this device complies. Only one interface version is current at a
	moment in time and all new devices should be built to the latest
	interface version. Applications can choose which device interface
	versions they support and it is in their interest to support previous
	versions as well as the current version to ensure thay can use the
	largest number of devices.
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__interfaceversion
*/
func (a *ASCOMAlpacaAPIClient) GetInterfaceVersion(deviceType string, deviceNumber uint) (int32, error) {
	return a.GetInt32Response(deviceType, deviceNumber, "interfaceversion")
}

/*
	GetName() common method to all ASCOM Alpaca compliant devices

	@returns the name of the device
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__name
*/
func (a *ASCOMAlpacaAPIClient) GetName(deviceType string, deviceNumber uint) (string, error) {
	return a.GetStringResponse(deviceType, deviceNumber, "name")
}

/*
	GetSupportedActions() common method to all ASCOM Alpaca compliant devices

	@returns the list of action names supported by this driver.
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__supportedactions
*/
func (a *ASCOMAlpacaAPIClient) GetSupportedActions(deviceType string, deviceNumber uint) ([]string, error) {
	return a.GetStringListResponse(deviceType, deviceNumber, "supportedactions")
}
