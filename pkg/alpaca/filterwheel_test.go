package alpaca

import "testing"

var filterwheel = NewFilterWheel(65535, true, "alpaca.observerly.com", "", -1, 0)

func TestNewFilterWheelBaseURL(t *testing.T) {
	filterwheel := NewFilterWheel(65535, false, "", "0.0.0.0", 8000, 0)

	var got string = filterwheel.Alpaca.UrlBase
	var want string = "http://0.0.0.0:8000"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewFilterWheelBaseURLForHost(t *testing.T) {
	var got string = filterwheel.Alpaca.UrlBase
	var want string = "https://alpaca.observerly.com"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewFilterWheelClientID(t *testing.T) {
	filterwheel := NewFilterWheel(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = filterwheel.Alpaca.ClientId
	var want uint32 = 65535

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewFilterWheelTransactionID(t *testing.T) {
	filterwheel := NewFilterWheel(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = filterwheel.Alpaca.TransactionId
	var want uint32 = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewFilterWheelDeviceNumber(t *testing.T) {
	filterwheel := NewFilterWheel(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint = filterwheel.DeviceNumber
	var want uint = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
