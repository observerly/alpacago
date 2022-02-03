package alpaca

import (
	"math"
	"testing"
)

func TestNewAlpacaAPIBaseURL(t *testing.T) {
	client := NewAlpacaAPI(65535, false, "", "0.0.0.0", 8000)

	var got string = client.UrlBase
	var want string = "http://0.0.0.0:8000"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewAlpacaAPIBaseURLForHost(t *testing.T) {
	client := NewAlpacaAPI(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1)

	var got string = client.UrlBase
	var want string = "https://virtserver.swaggerhub.com/ASCOMInitiative"

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
	a := NewAlpacaAPI(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1)

	got, err := a.GetStringResponse("telescope", 0, "description")

	var want string = "string"

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if a.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", a.ErrorMessage, want)
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

	if a.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", a.ErrorMessage, want)
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

	if a.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %t", a.ErrorMessage, want)
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

	if a.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %f", a.ErrorMessage, want)
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

	if a.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %d", a.ErrorMessage, want)
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

	if a.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %d", a.ErrorMessage, want)
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

	if a.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %t", a.ErrorMessage, want)
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

	if a.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", a.ErrorMessage, want)
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

	if a.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", a.ErrorMessage, want)
	}
}

func TestNewAlpacaAPIDriverVersion(t *testing.T) {
	a := NewAlpacaAPI(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1)

	got, err := a.GetDriverVersion("telescope", 0)

	var want string = "string"

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if a.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", a.ErrorMessage, want)
	}
}

func TestNewAlpacaAPIInterfaceVersion(t *testing.T) {
	a := NewAlpacaAPI(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1)

	got, err := a.GetInterfaceVersion("telescope", 0)

	var want int32 = 0

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if a.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", a.ErrorMessage, want)
	}
}

func TestNewAlpacaAPIName(t *testing.T) {
	a := NewAlpacaAPI(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1)

	got, err := a.GetName("telescope", 0)

	var want string = "string"

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if a.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", a.ErrorMessage, want)
	}
}

func TestNewAlpacaAPDSupportedActions(t *testing.T) {
	a := NewAlpacaAPI(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1)

	got, err := a.GetSupportedActions("telescope", 0)

	var want string = "string"

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got[0] != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if a.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", a.ErrorMessage, want)
	}
}
