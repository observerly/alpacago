package alpaca

import "testing"

func TestNewAlpacaAPIBaseURL(t *testing.T) {
	client := NewAlpacaAPI(65535, "0.0.0.0", 8000)

	var got string = client.urlBase
	var want string = "http://0.0.0.0:8000"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewAlpacaAPIClientID(t *testing.T) {
	client := NewAlpacaAPI(65535, "0.0.0.0", 8000)

	var got uint32 = client.clientId
	var want uint32 = 65535

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewAlpacaAPITransactionID(t *testing.T) {
	client := NewAlpacaAPI(65535, "0.0.0.0", 8000)

	var got uint32 = client.transactionId
	var want uint32 = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
