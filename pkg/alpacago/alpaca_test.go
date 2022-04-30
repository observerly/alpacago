package alpacago

import (
	"math"
	"testing"
)

var client = NewAlpacaAPI(65535, true, "alpaca.observerly.com", "", -1)

func TestNewAlpacaAPIBaseURL(t *testing.T) {
	client := NewAlpacaAPI(65535, false, "", "0.0.0.0", 8000)

	var got string = client.UrlBase
	var want string = "http://0.0.0.0:8000"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewAlpacaAPIBaseURLForHost(t *testing.T) {
	var got string = client.UrlBase
	var want string = "https://alpaca.observerly.com"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewAlpacaAPIClientID(t *testing.T) {
	client := NewAlpacaAPI(65535, false, "", "0.0.0.0", 8000)

	var got uint32 = client.ClientId
	var want uint32 = 65535

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewAlpacaAPITransactionID(t *testing.T) {
	client := NewAlpacaAPI(65535, false, "", "0.0.0.0", 8000)

	var got uint32 = client.TransactionId
	var want uint32 = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewAlpacaAPIQueryString(t *testing.T) {
	client := NewAlpacaAPI(65535, false, "", "0.0.0.0", 8000)

	var got string = client.getQueryString()
	var want string = "ClientID=65535&ClientTransactionID=0"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewAlpacaAPIEndpoint(t *testing.T) {
	client := NewAlpacaAPI(65535, false, "", "0.0.0.0", 8000)

	var got string = client.getEndpoint("telescope", 0, "canslew")
	var want string = "http://0.0.0.0:8000/api/v1/telescope/0/canslew"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewAlpacaAPIStringResponse(t *testing.T) {
	got, err := client.GetStringResponse("telescope", 0, "description")

	var want string = "Software Telescope Simulator for ASCOM"

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if client.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", client.ErrorMessage, want)
	}
}

func TestNewAlpacaAPIStringListResponse(t *testing.T) {
	got, err := client.GetStringListResponse("telescope", 0, "supportedactions")

	var want []string = []string{"AssemblyVersionNumber", "SlewToHA", "AvailableTimeInThisPointingState", "TimeUntilPointingStateCanChange"}

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got[0] != want[0] {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if got[1] != want[1] {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if got[2] != want[2] {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if got[3] != want[3] {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if client.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", client.ErrorMessage, want)
	}
}

func TestNewAlpacaAPIBooleanResponse(t *testing.T) {
	_, err := client.GetBooleanResponse("telescope", 0, "connected")

	if err != nil {
		t.Errorf("got %q, wanted a boolean value", err)
	}

	if client.ErrorNumber != 0 {
		t.Errorf("got %q, wanted a boolean value", client.ErrorMessage)
	}
}

func TestNewAlpacaAPIFloat64Response(t *testing.T) {
	got, err := client.GetFloat64Response("telescope", 0, "focallength")

	var want float64 = 1.260000

	if err != nil {
		t.Errorf("got %q, wanted %f", err, want)
	}

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}

	if client.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %f", client.ErrorMessage, want)
	}
}

func TestNewAlpacaAPIInt32Response(t *testing.T) {
	got, err := client.GetInt32Response("telescope", 0, "alignmentmode")

	var want int32 = 2

	if err != nil {
		t.Errorf("got %q, wanted %d", err, want)
	}

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	if client.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %d", client.ErrorMessage, want)
	}
}

func TestNewAlpacaAPIInt32ListResponse(t *testing.T) {
	got, err := client.GetUInt32ListResponse("filterwheel", 0, "focusoffsets")

	want := [6]uint32{1458, 2070, 1413, 9664, 3593, 826}

	if err != nil {
		t.Errorf("got %q, wanted %d", err, want)
	}

	if got[0] != want[0] {
		t.Errorf("got %d, wanted %d", got, want)
	}

	if got[1] != want[1] {
		t.Errorf("got %d, wanted %d", got, want)
	}

	if got[2] != want[2] {
		t.Errorf("got %d, wanted %d", got, want)
	}

	if got[3] != want[3] {
		t.Errorf("got %d, wanted %d", got, want)
	}

	if got[4] != want[4] {
		t.Errorf("got %d, wanted %d", got, want)
	}

	if got[5] != want[5] {
		t.Errorf("got %d, wanted %d", got, want)
	}

	if client.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %d", client.ErrorMessage, want)
	}
}

func TestNewAlpacaAPIConnected(t *testing.T) {
	_, err := client.IsConnected("telescope", 0)

	if err != nil {
		t.Errorf("got %q, wanted a boolean value", err)
	}

	if client.ErrorNumber != 0 {
		t.Errorf("got %q, wanted a boolean value", client.ErrorMessage)
	}
}

func TestNewAlpacaAPIDescription(t *testing.T) {
	got, err := client.GetDescription("telescope", 0)

	var want string = "Software Telescope Simulator for ASCOM"

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if client.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", client.ErrorMessage, want)
	}
}

func TestNewAlpacaAPIDriverInfo(t *testing.T) {
	got, err := client.GetDriverInfo("telescope", 0)

	var want string = "TelescopeSimulator, Version=0.1.3.0, Culture=neutral, PublicKeyToken=null"

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if client.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", client.ErrorMessage, want)
	}
}

func TestNewAlpacaAPIDriverVersion(t *testing.T) {
	got, err := client.GetDriverVersion("telescope", 0)

	var want string = "0.1"

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if client.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", client.ErrorMessage, want)
	}
}

func TestNewAlpacaAPIInterfaceVersion(t *testing.T) {
	got, err := client.GetInterfaceVersion("telescope", 0)

	var want int32 = 3

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if client.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", client.ErrorMessage, want)
	}
}

func TestNewAlpacaAPIName(t *testing.T) {
	got, err := client.GetName("telescope", 0)

	var want string = "Alpaca Telescope Sim"

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if client.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", client.ErrorMessage, want)
	}
}

func TestNewAlpacaAPDSupportedActions(t *testing.T) {
	got, err := client.GetSupportedActions("telescope", 0)

	var want []string = []string{"AssemblyVersionNumber", "SlewToHA", "AvailableTimeInThisPointingState", "TimeUntilPointingStateCanChange"}

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got[0] != want[0] {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if got[1] != want[1] {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if got[2] != want[2] {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if got[3] != want[3] {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if client.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", client.ErrorMessage, want)
	}
}
