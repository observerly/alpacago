package alpaca

import "testing"

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
