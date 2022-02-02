package alpaca

import (
	"math"
	"testing"
)

func TestNewAlpacaAPIBaseURL(t *testing.T) {
	client := NewAlpacaAPI(65535, false, "", "0.0.0.0", 8000)

	var got string = client.urlBase
	var want string = "http://0.0.0.0:8000"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewAlpacaAPIBaseURLForHost(t *testing.T) {
	client := NewAlpacaAPI(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1)

	var got string = client.urlBase
	var want string = "https://virtserver.swaggerhub.com/ASCOMInitiative"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewAlpacaAPIClientID(t *testing.T) {
	client := NewAlpacaAPI(65535, false, "", "0.0.0.0", 8000)

	var got uint32 = client.clientId
	var want uint32 = 65535

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewAlpacaAPITransactionID(t *testing.T) {
	client := NewAlpacaAPI(65535, false, "", "0.0.0.0", 8000)

	var got uint32 = client.transactionId
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
	a := NewAlpacaAPI(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1)

	got, err := a.GetStringResponse("telescope", 0, "description")

	var want string = "string"

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if a.errorNumber != 0 {
		t.Errorf("got %q, wanted %q", a.errorMessage, want)
	}
}

func TestNewAlpacaAPIStringListResponse(t *testing.T) {
	a := NewAlpacaAPI(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1)

	got, err := a.GetStringListResponse("telescope", 0, "supportedactions")

	var want string = "string"

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got[0] != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if a.errorNumber != 0 {
		t.Errorf("got %q, wanted %q", a.errorMessage, want)
	}
}

func TestNewAlpacaAPIBooleanResponse(t *testing.T) {
	a := NewAlpacaAPI(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1)

	got, err := a.GetBooleanResponse("telescope", 0, "connected")

	var want bool = true

	if err != nil {
		t.Errorf("got %q, wanted %t", err, want)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if a.errorNumber != 0 {
		t.Errorf("got %q, wanted %t", a.errorMessage, want)
	}
}

func TestNewAlpacaAPIFloat64Response(t *testing.T) {
	a := NewAlpacaAPI(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1)

	got, err := a.GetFloat64Response("telescope", 0, "focallength")

	var want float64 = 1.1

	if err != nil {
		t.Errorf("got %q, wanted %f", err, want)
	}

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}

	if a.errorNumber != 0 {
		t.Errorf("got %q, wanted %f", a.errorMessage, want)
	}
}

func TestNewAlpacaAPIInt32Response(t *testing.T) {
	a := NewAlpacaAPI(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1)

	got, err := a.GetInt32Response("telescope", 0, "alignmentmode")

	var want int32 = 0

	if err != nil {
		t.Errorf("got %q, wanted %d", err, want)
	}

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	if a.errorNumber != 0 {
		t.Errorf("got %q, wanted %d", a.errorMessage, want)
	}
}

func TestNewAlpacaAPIInt32ListResponse(t *testing.T) {
	a := NewAlpacaAPI(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1)

	got, err := a.GetUInt32ListResponse("filterwheel", 0, "focusoffsets")

	var want uint32 = 0

	if err != nil {
		t.Errorf("got %q, wanted %d", err, want)
	}

	if got[0] != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	if a.errorNumber != 0 {
		t.Errorf("got %q, wanted %d", a.errorMessage, want)
	}
}

func TestNewAlpacaAPIConnected(t *testing.T) {
	a := NewAlpacaAPI(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1)

	got, err := a.IsConnected("telescope", 0)

	var want bool = true

	if err != nil {
		t.Errorf("got %q, wanted %t", err, want)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if a.errorNumber != 0 {
		t.Errorf("got %q, wanted %t", a.errorMessage, want)
	}
}

func TestNewAlpacaAPIDescription(t *testing.T) {
	a := NewAlpacaAPI(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1)

	got, err := a.GetDescription("telescope", 0)

	var want string = "string"

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if a.errorNumber != 0 {
		t.Errorf("got %q, wanted %q", a.errorMessage, want)
	}
}

func TestNewAlpacaAPIDriverInfo(t *testing.T) {
	a := NewAlpacaAPI(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1)

	got, err := a.GetDriverInfo("telescope", 0)

	var want string = "string"

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if a.errorNumber != 0 {
		t.Errorf("got %q, wanted %q", a.errorMessage, want)
	}
}
