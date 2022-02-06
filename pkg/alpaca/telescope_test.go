package alpaca

import (
	"math"
	"testing"
	"time"
)

var delay time.Duration = 1

func TestNewTelescopeBaseURL(t *testing.T) {
	telescope := NewTelescope(65535, false, "", "0.0.0.0", 8000, 0, 1)

	var got string = telescope.Alpaca.UrlBase
	var want string = "http://0.0.0.0:8000"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewTelescopeBaseURLForHost(t *testing.T) {
	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got string = telescope.Alpaca.UrlBase
	var want string = "https://virtserver.swaggerhub.com/ASCOMInitiative"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewTelescopeClientID(t *testing.T) {
	telescope := NewTelescope(65535, false, "", "0.0.0.0", 8000, 0, 1)

	var got uint32 = telescope.Alpaca.ClientId
	var want uint32 = 65535

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewTelescopeTransactionID(t *testing.T) {
	telescope := NewTelescope(65535, false, "", "0.0.0.0", 8000, 0, 1)

	var got uint32 = telescope.Alpaca.TransactionId
	var want uint32 = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewTelescopeDeviceNumber(t *testing.T) {
	telescope := NewTelescope(65535, false, "", "0.0.0.0", 8000, 0, 1)

	var got uint = telescope.DeviceNumber
	var want uint = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewTelescopeTrackingMode(t *testing.T) {
	telescope := NewTelescope(65535, false, "", "0.0.0.0", 8000, 0, 1)

	var got = telescope.Tracking
	var want = Alt_Az

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewTelescopeAlignmentMode(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetAlignmentMode()

	var want = AlignmentAltAz

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %d", telescope.Alpaca.ErrorMessage, want)
	}
}

func TestNewTelescopeAltitude(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetAltitude()

	var want float64 = 1.1

	if err != nil {
		t.Errorf("got %q, wanted %f", err, want)
	}

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %f", telescope.Alpaca.ErrorMessage, want)
	}
}

func TestNewTelescopeApertureArea(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetApertureArea()

	var want float64 = 1.1

	if err != nil {
		t.Errorf("got %q, wanted %f", err, want)
	}

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %f", telescope.Alpaca.ErrorMessage, want)
	}
}

func TestNewTelescopeApertureDiameter(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetApertureDiameter()

	var want float64 = 1.1

	if err != nil {
		t.Errorf("got %q, wanted %f", err, want)
	}

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %f", telescope.Alpaca.ErrorMessage, want)
	}
}

func TestNewTelescopeAtHome(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.IsAtHome()

	var want bool = true

	if err != nil {
		t.Errorf("got %q, wanted %t", err, want)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %t", telescope.Alpaca.ErrorMessage, want)
	}
}

func TestNewTelescopeAtPark(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.IsAtPark()

	var want bool = true

	if err != nil {
		t.Errorf("got %q, wanted %t", err, want)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %t", telescope.Alpaca.ErrorMessage, want)
	}
}

func TestNewTelescopeAzimuth(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetAzimuth()

	var want float64 = 1.1

	if err != nil {
		t.Errorf("got %q, wanted %f", err, want)
	}

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %f", telescope.Alpaca.ErrorMessage, want)
	}
}

func TestNewTelescopeCanFindHome(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.CanFindHome()

	var want bool = true

	if err != nil {
		t.Errorf("got %q, wanted %t", err, want)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %t", telescope.Alpaca.ErrorMessage, want)
	}
}
