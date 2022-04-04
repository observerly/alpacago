package alpacago

import "testing"

var message = NewDiscoveryMessage(ALPACA_DISCOVERY_VERSION)

func TestNewDiscoveryMessageFixed(t *testing.T) {
	var got string = string(message.Fixed)

	var want string = "alpacadiscovery"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewDiscoveryMessageReserved(t *testing.T) {
	var got []byte = message.Reserved

	var want []byte = make([]byte, 48)

	if len(got) != len(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestNewDiscoveryMessageVersion(t *testing.T) {
	var got string = string(message.Version)

	var want string = "1"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
