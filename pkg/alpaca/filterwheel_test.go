package alpaca

import (
	"testing"
)

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

func TestNewFilterWheelGetFocusOffsets(t *testing.T) {
	var got, err = filterwheel.GetFocusOffsets()

	want := [6]uint32{8280, 8079, 9234, 790, 6553, 4650}

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got[0] != want[0] {
		t.Errorf("got %d, wanted %q", got[0], want[0])
	}

	if got[1] != want[1] {
		t.Errorf("got %d, wanted %q", got[1], want[1])
	}

	if got[2] != want[2] {
		t.Errorf("got %d, wanted %q", got[2], want[2])
	}

	if got[3] != want[3] {
		t.Errorf("got %d, wanted %q", got[3], want[3])
	}

	if got[4] != want[4] {
		t.Errorf("got %d, wanted %q", got[4], want[4])
	}

	if got[5] != want[5] {
		t.Errorf("got %d, wanted %q", got[5], want[5])
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %d", telescope.Alpaca.ErrorMessage, want)
	}
}