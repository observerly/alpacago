package alpacago

import (
	"math"
	"testing"
)

var rotator = NewRotator(65535, false, "100.80.84.116", "", -1, 0)

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
	var want string = "http://100.80.84.116"

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

func TestNewRotatorGetDescription(t *testing.T) {
	var got, err = rotator.GetDescription()

	var want = "ASCOM Rotator Driver for RotatorSimulator"

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

func TestNewRotatorIsConnected(t *testing.T) {
	var got, err = rotator.IsConnected()

	var want = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if rotator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", rotator.Alpaca.ErrorMessage)
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

	var want = 0.000000

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}

	if rotator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", rotator.Alpaca.ErrorMessage)
	}
}

func TestNewRotatorGetPosition(t *testing.T) {
	var got, err = rotator.GetPosition()

	var want = 0.000000

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}

	if rotator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", rotator.Alpaca.ErrorMessage)
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

	if rotator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", rotator.Alpaca.ErrorMessage)
	}
}

func TestNewRotatorSetReverse(t *testing.T) {
	var err = rotator.SetReverse(false)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if rotator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", rotator.Alpaca.ErrorMessage)
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

	if rotator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", rotator.Alpaca.ErrorMessage)
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

	if rotator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", rotator.Alpaca.ErrorMessage)
	}
}

func TestNewRotatorSetHalt(t *testing.T) {
	var err = rotator.SetHalt()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if rotator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", rotator.Alpaca.ErrorMessage)
	}
}

func TestNewRotatorSetMove(t *testing.T) {
	var _ = focuser.SetConnected(true)

	var err = rotator.SetMove(0)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if rotator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", rotator.Alpaca.ErrorMessage)
	}
}

func TestNewRotatorSetMoveAbsolute(t *testing.T) {
	var _ = focuser.SetConnected(true)

	var err = rotator.SetMoveAbsolute(0)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if rotator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", rotator.Alpaca.ErrorMessage)
	}
}

func TestNewRotatorSetMoveMechnical(t *testing.T) {
	var _ = rotator.SetConnected(true)

	var err = rotator.SetMoveMechanical(0)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if rotator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", rotator.Alpaca.ErrorMessage)
	}
}

func TestNewRotatorSetSync(t *testing.T) {
	var _ = rotator.SetConnected(true)

	var err = rotator.SetSync(0)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if rotator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", rotator.Alpaca.ErrorMessage)
	}
}
