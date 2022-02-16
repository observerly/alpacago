package alpaca

import (
	"math"
	"testing"
	"time"
)

var delay time.Duration = 9

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

func TestNewTelescopeCanMoveAxis(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.CanMoveAxis(0)

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

func TestNewTelescopeCanPark(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.CanPark()

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

func TestNewTelescopeCanPulseGuide(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.CanPulseGuide()

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

func TestNewTelescopeCanSetDeclinationRate(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.CanSetDeclinationRate()

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

func TestNewTelescopeCanSetGuideRates(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.CanSetGuideRates()

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

func TestNewTelescopeCanSetPark(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.CanSetPark()

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

func TestNewTelescopeCanSetPierSide(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.CanSetPierSide()

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

func TestNewTelescopeCanSetRightAscensionRate(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.CanSetRightAscensionRate()

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

func TestNewTelescopeCanSetTracking(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.CanSetTracking()

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

func TestNewTelescopeCanSlew(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.CanSlew()

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

func TestNewTelescopeCanSlewAltAz(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.CanSlewAltAz()

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

func TestNewTelescopeCanSlewAltAzAsync(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.CanSlewAltAzAsync()

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

func TestNewTelescopeCanSlewAsync(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.CanSlewAsync()

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

func TestNewTelescopeCanSync(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.CanSync()

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

func TestNewTelescopeCanSyncAltAz(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.CanSyncAltAz()

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

func TestNewTelescopeCanUnPark(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.CanUnPark()

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

func TestNewTelescopeDeclination(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetDeclination()

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

func TestNewTelescopeDeclinationRate(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetDeclinationRate()

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

func TestNewTelescopeDeclinationRatePut(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var err = telescope.SetDeclinationRate(5)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeDoesRefraction(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.DoesRefraction()

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

func TestNewTelescopeSetDoesRefractionPut(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var err = telescope.SetDoesRefraction(true)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeEquatorialSystem(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetEquatorialSystem()

	var want = Topocentric

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

func TestNewTelescopeFocalLength(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetFocalLength()

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

func TestNewTelescopeIsPulseGuiding(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.IsPulseGuiding()

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

func TestNewTelescopeRightAscension(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetRightAscension()

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

func TestNewTelescopeRightAscensionRate(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetRightAscensionRate()

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

func TestNewTelescopeSetRightAscensionRatePut(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var err = telescope.SetRightAscensionRate(5)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeSideOfPier(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetSideOfPier()

	var want = PierEast

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

func TestNewTelescopeSetSideOfPierEastPut(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var err = telescope.SetSideOfPier(0)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeSetSideOfPierWestPut(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var err = telescope.SetSideOfPier(1)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeSetSideOfPierInvalidPut(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var err = telescope.SetSideOfPier(-1)

	if err == nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeSiderealTime(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetSiderealTime()

	var want float64 = 1.100000023841858

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

func TestNewTelescopeSiteElevation(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetSiteElevation()

	var want float64 = 1.100000023841858

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

func TestNewTelescopeSiteLatitude(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetSiteLatitude()

	var want float64 = 1.100000023841858

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

func TestNewTelescopeSiteLongitude(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetSiteLongitude()

	var want float64 = 1.100000023841858

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

func TestNewTelescopeSlewing(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.IsSlewing()

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

func TestNewTelescopeGetSlewSettleTime(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetSlewSettleTime()

	var want int32 = 0

	if err != nil {
		t.Errorf("got %q, wanted %d", err, want)
	}

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %d", telescope.Alpaca.ErrorMessage, want)
	}
}

func TestNewTelescopeTargetDeclination(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetTargetDeclination()

	var want float64 = 1.100000023841858

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

func TestNewTelescopeTargetRightAscension(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetTargetRightAscension()

	var want float64 = 1.100000023841858

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

func TestNewTelescopeTracking(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.IsTracking()

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

func TestNewTelescopeGetTrackingRate(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetTrackingRate()

	var want int32 = 0

	if err != nil {
		t.Errorf("got %q, wanted %d", err, want)
	}

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %d", telescope.Alpaca.ErrorMessage, want)
	}
}

func TestNewTelescopeGetUTCDate(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetUTCDate()

	var want time.Time = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want.String())
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want.String())
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", telescope.Alpaca.ErrorMessage, want.String())
	}
}

func TestNewTelescopeGetAxisRates(t *testing.T) {
	time.Sleep(delay * time.Second)

	telescope := NewTelescope(65535, true, "virtserver.swaggerhub.com/ASCOMInitiative", "", -1, 0, 1)

	var got, err = telescope.GetAxisRates(0)

	var want = make(map[string]float64)

	want["Maximum"] = 1.5
	want["Minimum"] = 1.5

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got["Maximum"] != want["Maximum"] {
		t.Errorf("got %f, wanted %f", got["Maximum"], want["Maximum"])
	}

	if got["Minimum"] != want["Minimum"] {
		t.Errorf("got %f, wanted %f", got["Minimum"], want["Minimum"])
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}
