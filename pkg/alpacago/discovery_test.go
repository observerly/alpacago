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

func TestNewDiscoveryServerLocalHostAlias(t *testing.T) {
	var server = NewDiscoveryServer(65535, "udp4", "", "localhost", 8080)

	var got string = server.UrlBase

	var want string = "localhost:8080"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewDiscoveryServerPort32227(t *testing.T) {
	var server = NewDiscoveryServer(65535, "udp4", "", "0.0.0.0", 32227)

	var got string = server.UrlBase

	var want string = "0.0.0.0:32227"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewDiscoveryServerPort30999(t *testing.T) {
	var server = NewDiscoveryServer(65535, "udp4", "", "127.0.0.1", 30999)

	var got string = server.UrlBase

	var want string = "127.0.0.1:30999"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewDiscoveryServerMalformedIPError(t *testing.T) {
	var server = NewDiscoveryServer(65535, "udp4", "", "127.0.0.1", 30999)

	var got string = server.UrlBase

	var want string = "127.0.0.1:30999"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewDiscoveryServerOutOfRangePortError(t *testing.T) {
	var server = NewDiscoveryServer(65535, "udp6", "", "127.0.0.1", 67900)

	var got string = server.UrlBase

	var want string = "127.0.0.1:11111"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewDiscoveryServerOutOfRangePortErrorAlt(t *testing.T) {
	var server = NewDiscoveryServer(65535, "udp6", "", "localhost", 67900)

	var got string = server.UrlBase

	var want string = "localhost:11111"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
