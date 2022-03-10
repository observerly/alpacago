package alpaca

import (
	"math"
	"testing"
)

var rotator = NewRotator(65535, true, "alpaca.observerly.com", "", -1, 0)

func TestNewRotatorBaseURL(t *testing.T) {
	rotator := NewRotator(65535, false, "", "0.0.0.0", 8000, 0)

	var got string = rotator.Alpaca.UrlBase
	var want string = "http://0.0.0.0:8000"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewRotatorBaseURLForHost(t *testing.T) {
	var got string = rotator.Alpaca.UrlBase
	var want string = "https://alpaca.observerly.com"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewRotatorClientID(t *testing.T) {
	rotator := NewRotator(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = rotator.Alpaca.ClientId
	var want uint32 = 65535

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewRotatorTransactionID(t *testing.T) {
	rotator := NewRotator(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = rotator.Alpaca.TransactionId
	var want uint32 = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewRotatorDeviceNumber(t *testing.T) {
	rotator := NewRotator(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint = rotator.DeviceNumber
	var want uint = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewRotatorSetConnected(t *testing.T) {
	var err = rotator.SetConnected(true)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if rotator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", rotator.Alpaca.ErrorMessage)
	}
}

func TestNewRotatorCanReverse(t *testing.T) {
	var got, err = rotator.CanReverse()

	var want = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}

	if rotator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", rotator.Alpaca.ErrorMessage)
	}
}

func TestNewRotatorIsMoving(t *testing.T) {
	var got, err = rotator.IsMoving()

	var want = false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}

	if rotator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", rotator.Alpaca.ErrorMessage)
	}
}

func TestNewRotatorGetMechanicalPosition(t *testing.T) {
	var got, err = rotator.GetMechanicalPosition()

	var want = 0.0

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}
}

func TestNewRotatorGetPosition(t *testing.T) {
	var got, err = rotator.GetPosition()

	var want = 0.0

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}
}

func TestNewRotatorGetReverse(t *testing.T) {
	var got, err = rotator.GetReverse()

	var want = false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}
}

func TestNewRotatorSetReverse(t *testing.T) {
	var err = rotator.SetReverse(false)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}
}

func TestNewRotatorGetStepSize(t *testing.T) {
	var got, err = rotator.GetStepSize()

	var want = 0.75

	if err != nil {
		t.Errorf("got %q", err)
	}

	if math.Abs(got-want) > 50 {
		t.Errorf("got %f, wanted %f", got, want)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}
}

func TestNewRotatorGetTargetPosition(t *testing.T) {
	var got, err = rotator.GetTargetPosition()

	var want = 0.75

	if err != nil {
		t.Errorf("got %q", err)
	}

	if math.Abs(got-want) > 50 {
		t.Errorf("got %f, wanted %f", got, want)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}
}

func TestNewRotatorSetHalt(t *testing.T) {
	var err = rotator.SetHalt()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}
}
