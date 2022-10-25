package alpacago

import (
	"testing"
)

var monitor = NewSafetyMonitor(65535, false, "100.80.84.116", "", -1, 0)

func TestNewSafetyMonitorBaseURL(t *testing.T) {
	monitor := NewSafetyMonitor(65535, false, "", "0.0.0.0", 8000, 0)

	var got string = monitor.Alpaca.UrlBase
	var want string = "http://0.0.0.0:8000"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewSafetyMonitorBaseURLForHost(t *testing.T) {
	var got string = monitor.Alpaca.UrlBase
	var want string = "http://100.80.84.116"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewSafetyMonitorClientID(t *testing.T) {
	monitor := NewSafetyMonitor(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = monitor.Alpaca.ClientId
	var want uint32 = 65535

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewSafetyMonitorTransactionID(t *testing.T) {
	monitor := NewSafetyMonitor(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = monitor.Alpaca.TransactionId
	var want uint32 = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewSafetyMonitorDeviceNumber(t *testing.T) {
	monitor := NewSafetyMonitor(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint = monitor.DeviceNumber
	var want uint = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewSafetyMonitorIsConnected(t *testing.T) {
	var got, err = monitor.IsConnected()

	var want = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if monitor.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", monitor.Alpaca.ErrorMessage)
	}
}

func TestNewSafetyMonitorSetConnected(t *testing.T) {
	monitor.SetConnected(true)

	var got, err = monitor.IsConnected()

	var want = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if monitor.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", monitor.Alpaca.ErrorMessage)
	}
}

func TestNewSafetyMonitorIsSafe(t *testing.T) {
	var got, err = monitor.IsSafe()

	var want = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if monitor.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", monitor.Alpaca.ErrorMessage)
	}
}
