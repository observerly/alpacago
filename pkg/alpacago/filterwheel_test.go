package alpacago

import (
	"testing"
)

var filterwheel = NewFilterWheel(65535, false, "100.80.84.116", "", -1, 0)

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
	var want string = "http://100.80.84.116"

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

func TestNewFilterWheelSetConnected(t *testing.T) {
	var err = filterwheel.SetConnected(true)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if filterwheel.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", filterwheel.Alpaca.ErrorMessage)
	}
}

func TestNewFilterWheelGetFocusOffsets(t *testing.T) {
	var got, err = filterwheel.GetFocusOffsets()

	want := [6]uint32{2295, 9177, 3053, 5430, 2965, 4952}

	if err != nil {
		t.Errorf("got %, wanted %q", err, want)
	}

	if got[0] != want[0] {
		t.Errorf("got %v, wanted %v", got[0], want[0])
	}

	if got[1] != want[1] {
		t.Errorf("got %v, wanted %v", got[1], want[1])
	}

	if got[2] != want[2] {
		t.Errorf("got %v, wanted %v", got[2], want[2])
	}

	if got[3] != want[3] {
		t.Errorf("got %v, wanted %v", got[3], want[3])
	}

	if got[4] != want[4] {
		t.Errorf("got %v, wanted %v", got[4], want[4])
	}

	if got[5] != want[5] {
		t.Errorf("got %v, wanted %v", got[5], want[5])
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %d", telescope.Alpaca.ErrorMessage, want)
	}
}

func TestNewFilterWheelGetNames(t *testing.T) {
	var got, err = filterwheel.GetNames()

	want := [6]string{"Red", "Green", "Blue", "Clear", "Ha", "OIII"}

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got[0] != want[0] {
		t.Errorf("got %q, wanted %q", got[0], want[0])
	}

	if got[1] != want[1] {
		t.Errorf("got %q, wanted %q", got[1], want[1])
	}

	if got[2] != want[2] {
		t.Errorf("got %q, wanted %q", got[2], want[2])
	}

	if got[3] != want[3] {
		t.Errorf("got %q, wanted %q", got[3], want[3])
	}

	if got[4] != want[4] {
		t.Errorf("got %q, wanted %q", got[4], want[4])
	}

	if got[5] != want[5] {
		t.Errorf("got %q, wanted %q", got[5], want[5])
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", telescope.Alpaca.ErrorMessage, want)
	}
}

func TestNewFilterWheelGetPosition(t *testing.T) {
	var got, err = filterwheel.GetPosition()

	var want int32 = 0

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", telescope.Alpaca.ErrorMessage, want)
	}
}

func TestNewFilterWheelSetPosition(t *testing.T) {
	var err = filterwheel.SetPosition(0)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}
