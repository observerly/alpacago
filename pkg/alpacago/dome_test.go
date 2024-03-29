package alpacago

import (
	"strings"
	"testing"
	"time"
)

var dome = NewDome(65535, false, "100.69.47.32", "", -1, 0)

func TestNewDomeBaseURL(t *testing.T) {
	dome := NewDome(65535, false, "", "0.0.0.0", 8000, 0)

	var got string = dome.Alpaca.UrlBase
	var want string = "http://0.0.0.0:8000"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewDomeBaseURLForHost(t *testing.T) {
	var got string = dome.Alpaca.UrlBase
	var want string = "http://100.69.47.32"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewDomeClientID(t *testing.T) {
	dome := NewDome(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = dome.Alpaca.ClientId
	var want uint32 = 65535

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewDomeTransactionID(t *testing.T) {
	dome := NewDome(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = dome.Alpaca.TransactionId
	var want uint32 = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewDomeDeviceNumber(t *testing.T) {
	dome := NewDome(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint = dome.DeviceNumber
	var want uint = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewDomeIsConnected(t *testing.T) {
	var got, err = dome.IsConnected()

	var want = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeIsConnectedOn(t *testing.T) {
	var err = dome.SetConnected(true)

	var got, _ = dome.IsConnected()

	var want = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeGetAltitude(t *testing.T) {
	var got, err = dome.GetAltitude()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 || got > 90 {
		t.Errorf("got %v, but expected the dome altitude to be between 0° and +90°", got)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeAtHome(t *testing.T) {
	var got, err = dome.IsAtHome()

	var want bool = false

	if err != nil {
		t.Errorf("got %q, wanted %t", err, want)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %t", dome.Alpaca.ErrorMessage, want)
	}
}

func TestNewDomeAtPark(t *testing.T) {
	var got, err = dome.IsAtPark()

	var want bool = false

	if err != nil {
		t.Errorf("got %q, wanted %t", err, want)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %t", dome.Alpaca.ErrorMessage, want)
	}
}

func TestNewDomeAzimuth(t *testing.T) {
	var got, err = dome.GetAzimuth()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 || got > 360 {
		t.Errorf("got %f, but expected the dome azimuth to be between 0° and +360°", got)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeCanFindHome(t *testing.T) {
	var got, err = dome.CanFindHome()

	var want bool = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeCanPark(t *testing.T) {
	var got, err = dome.CanPark()

	var want bool = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeCanSetAltitude(t *testing.T) {
	var got, err = dome.CanSetAltitude()

	var want bool = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeCanSetAzimuth(t *testing.T) {
	var got, err = dome.CanSetAzimuth()

	var want bool = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeCanSetPark(t *testing.T) {
	var got, err = dome.CanSetPark()

	var want bool = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeCanSetShutter(t *testing.T) {
	var got, err = dome.CanSetShutter()

	var want bool = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeCanSlave(t *testing.T) {
	var got, err = dome.CanSlave()

	var want bool = false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeCanSyncAzimuth(t *testing.T) {
	var got, err = dome.CanSyncAzimuth()

	var want bool = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeShutterStatus(t *testing.T) {
	var got, err = dome.GetShutterStatus()

	if err != nil {
		t.Errorf("got %q", err)
	}

	want := []string{"open", "closed", "opening", "closing", "error"}

	if !strings.Contains(strings.Join(want, ","), got) {
		t.Errorf("got %q, wanted %q", got, "the shutter status to represent a valid status value")
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", dome.Alpaca.ErrorMessage, "the shutter status to represent a valid status value")
	}
}

func TestNewDomeShutterStatusToStringRepresentation(t *testing.T) {
	var got, err = dome.GetShutterStatus()

	if got == "" {
		t.Errorf("got %q, wanted %q", got, "the shutter status to represnet an iota in range 0 to 2")
	}

	if err != nil {
		t.Errorf("got %q", err)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", dome.Alpaca.ErrorMessage, "the shutter status to represnet an iota in range 0 to 2")
	}
}

func TestNewDomeIsSlaved(t *testing.T) {
	dome.SetSlaved(false)

	var got, err = dome.IsSlaved()

	var want bool = false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeSetSlaved(t *testing.T) {
	dome.SetSlaved(true)

	var got, err = dome.IsSlaved()

	var want bool = false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeIsSlewing(t *testing.T) {
	var got, err = dome.IsSlewing()

	var want bool = false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeAbortSlew(t *testing.T) {
	var err = dome.AbortSlew()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeOpenShutter(t *testing.T) {
	var err = dome.OpenShutter()

	if err != nil {
		t.Errorf("got %v, wanted %q", err, "the dome to be either open, or opening")
	}

	// Wait for the dome to open the shutter:
	time.Sleep(time.Second * 5)

	got, err := dome.GetShutterStatus()

	if err != nil {
		t.Errorf("got %v, wanted %q", err, "the dome to be either open, or opening")
	}

	want := []string{"open", "opening"}

	if !strings.Contains(strings.Join(want, ","), got) {
		t.Errorf("got %q, wanted %q", got, "the shutter status to represent a valid status value")
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", dome.Alpaca.ErrorMessage, want)
	}
}

func TestNewDomeFindHome(t *testing.T) {
	var err = dome.FindHome()

	if err != nil {
		t.Errorf("got %q", err)
	}

	time.Sleep(time.Second * 10)

	var got, _ = dome.IsAtHome()

	var want = true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomePark(t *testing.T) {
	var err = dome.Park()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeSetAsPark(t *testing.T) {
	var err = dome.SetAsPark()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeSlewToAltitude(t *testing.T) {
	var err = dome.SlewToAltitude(45)

	time.Sleep(time.Second * 8)

	var got, _ = dome.GetAltitude()

	var want float64 = 45

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeSlewToAzimuth(t *testing.T) {
	var err = dome.SlewToAzimuth(60)

	time.Sleep(time.Second * 8)

	var got, _ = dome.GetAzimuth()

	var want float64 = 60

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeSyncoAzimuth(t *testing.T) {
	var err = dome.SyncToAzimuth(80)

	time.Sleep(time.Second * 8)

	var got, _ = dome.GetAzimuth()

	var want float64 = 80

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeCloseShutter(t *testing.T) {
	var err = dome.CloseShutter()

	if err != nil {
		t.Errorf("got %q, wanted %q", err, "the dome to be either close, or closing")
	}

	got, err := dome.GetShutterStatus()

	if err != nil {
		t.Errorf("got %q, wanted %q", err, "the dome to be either close, or closing")
	}

	want := []string{"close", "closing"}

	if !strings.Contains(strings.Join(want, ","), got) {
		t.Errorf("got %q, wanted %q", got, "the shutter status to represent a valid status value")
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", dome.Alpaca.ErrorMessage, want)
	}
}
