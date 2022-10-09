package alpacago

import (
	"math"
	"testing"
	"time"
)

var latitude float64 = 19.820667

var longitude float64 = -155.468167

var telescope = NewTelescope(65535, false, "100.80.84.116", "", -1, 0, 1)

func TestNewTelescopeBaseURL(t *testing.T) {
	telescope := NewTelescope(65535, false, "", "0.0.0.0", 8000, 0, 1)

	var got string = telescope.Alpaca.UrlBase
	var want string = "http://0.0.0.0:8000"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewTelescopeBaseURLForHost(t *testing.T) {
	var got string = telescope.Alpaca.UrlBase
	var want string = "http://100.80.84.116"

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

func TestNewTelescopeIsConnected(t *testing.T) {
	var got, err = telescope.IsConnected()

	var want = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeSetConnected(t *testing.T) {
	telescope.SetConnected(true)

	var got, err = telescope.IsConnected()

	var want = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeSetAbortSlew(t *testing.T) {
	var err = telescope.SetAbortSlew()

	if err != nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeAlignmentMode(t *testing.T) {
	var got, err = telescope.GetAlignmentMode()

	var want = AlignmentGermanPolar

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
	var got, err = telescope.GetAltitude()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < -90 || got > 90 {
		t.Errorf("got %f", got)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeApertureArea(t *testing.T) {
	var got, err = telescope.GetApertureArea()

	var want float64 = 0.026900

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
	var got, err = telescope.GetApertureDiameter()

	var want float64 = 0.026900

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
	var got, err = telescope.IsAtHome()

	var want bool = false

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
	var got, err = telescope.IsAtPark()

	var want bool = false

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

func TestNewTelescopeAxisRates(t *testing.T) {
	var got, err = telescope.GetAxisRates(1)

	var want = make(map[string]float64)

	want["Maximum"] = 6.666667
	want["Minimum"] = 0.000000

	if err != nil {
		t.Errorf("got %q", err)
	}

	if math.Abs(got["Maximum"]-want["Maximum"]) > 0.00001 {
		t.Errorf("got %f, wanted %f", got["Maximum"], want["Maximum"])
	}

	if math.Abs(got["Minimum"]-want["Minimum"]) > 0.00001 {
		t.Errorf("got %f, wanted %f", got["Minimum"], want["Minimum"])
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeAzimuth(t *testing.T) {
	var got, err = telescope.GetAzimuth()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 || got > 360 {
		t.Errorf("got %f", got)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeCanFindHome(t *testing.T) {
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
	var got, err = telescope.GetDeclination()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < -90 || got > 90 {
		t.Errorf("got %f", got)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeDeclinationRate(t *testing.T) {
	var got, err = telescope.GetDeclinationRate()

	var want float64 = 5.000000

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
	var err = telescope.SetDeclinationRate(5)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeDoesRefraction(t *testing.T) {
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
	var err = telescope.SetDoesRefraction(true)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeEquatorialSystem(t *testing.T) {
	var got, err = telescope.GetEquatorialSystem()

	var want = J2000

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
	var got, err = telescope.GetFocalLength()

	var want float64 = 1.260000

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
	var got, err = telescope.IsPulseGuiding()

	var want bool = false

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
	var got, err = telescope.GetRightAscension()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 || got > 360 {
		t.Errorf("got %f, an unrelastic right ascension value", got)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeRightAscensionRate(t *testing.T) {
	var got, err = telescope.GetRightAscensionRate()

	var want float64 = 5.000000

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
	var err = telescope.SetRightAscensionRate(5)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeSideOfPier(t *testing.T) {
	var got, err = telescope.GetSideOfPier()

	var want = PierWest | PierEast

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got != PierWest {
		if got != PierEast {
			t.Errorf("got %q, wanted %q", got, PierEast)
		}
	}

	if got != PierEast {
		if got != PierWest {
			t.Errorf("got %q, wanted %q", got, PierWest)
		}
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %d", telescope.Alpaca.ErrorMessage, want)
	}
}

func TestNewTelescopeSetSideOfPierEastPut(t *testing.T) {
	var err = telescope.SetSideOfPier(0)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeSetSideOfPierWestPut(t *testing.T) {
	var err = telescope.SetSideOfPier(1)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeSetSideOfPierInvalidPut(t *testing.T) {
	var err = telescope.SetSideOfPier(-1)

	if err == nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeSiderealTime(t *testing.T) {
	var got, err = telescope.GetSiderealTime()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 || got > 24 {
		t.Errorf("got %f, an unrelastic sidereal value", got)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeSiteElevation(t *testing.T) {
	var got, err = telescope.GetSiteElevation()

	var want float64 = 4207.000000

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

func TestNewTelescopeSiteElevationPut(t *testing.T) {
	// Elevation at Mauna Kea, Hawaii:
	var err = telescope.SetSiteElevation(4207)

	if err != nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeSiteElevationPutUnrealisticallyHigh(t *testing.T) {
	var err = telescope.SetSiteElevation(20000)

	if err == nil {
		t.Errorf("got %q", err)
	}
}
func TestNewTelescopeSiteElevationPutUnrealisticallyLow(t *testing.T) {
	var err = telescope.SetSiteElevation(-1001)

	if err == nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeSiteLatitude(t *testing.T) {
	var got, err = telescope.GetSiteLatitude()

	var want float64 = 19.820667

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

func TestNewTelescopeSiteLatitudePutInvalid90Minus(t *testing.T) {
	var err = telescope.SetSiteLatitude(-91)

	if err == nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeSiteLatitudePutInvalid90Plus(t *testing.T) {
	var err = telescope.SetSiteLatitude(91)

	if err == nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeSiteLatitudePut(t *testing.T) {
	var err = telescope.SetSiteLatitude(latitude)

	if err != nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeSiteLongitude(t *testing.T) {
	var got, err = telescope.GetSiteLongitude()

	var want float64 = -155.468167

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

func TestNewTelescopeSiteLongitudePutInvalid180Minus(t *testing.T) {
	var err = telescope.SetSiteLongitude(-181)

	if err == nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeSiteLongitudePutInvalid180Plus(t *testing.T) {
	var err = telescope.SetSiteLongitude(181)

	if err == nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeSiteLongitudePut(t *testing.T) {
	var err = telescope.SetSiteLongitude(longitude)

	if err != nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeSlewing(t *testing.T) {
	var _, err = telescope.IsSlewing()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeGetSlewSettleTime(t *testing.T) {
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

func TestNewTelescopeSetSlewSettleTime(t *testing.T) {
	var err = telescope.SetSlewSettleTime(0)

	if err != nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeGetTargetDeclination(t *testing.T) {
	var got, err = telescope.GetTargetDeclination()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < -90 || got > 90 {
		t.Errorf("got %f", got)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeSetTargetDeclination(t *testing.T) {
	var err = telescope.SetTargetDeclination(7.4070639)

	if err != nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeTargetRightAscension(t *testing.T) {
	var got, err = telescope.GetTargetRightAscension()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 || got > 24 {
		t.Errorf("got %f", got)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeSetTargetRightAscension(t *testing.T) {
	var err = telescope.SetTargetRightAscension(88.7929583 / 24)

	if err != nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeTracking(t *testing.T) {
	var _, err = telescope.IsTracking()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewTelescopeSetTracking(t *testing.T) {
	var err = telescope.SetTracking(false)

	if err != nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeGetTrackingRate(t *testing.T) {
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
	var got, err = telescope.GetUTCDate()

	var want time.Time = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want.String())
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want.String())
	}

	if telescope.Alpaca.ErrorNumber != 0 && telescope.Alpaca.ErrorMessage != "" {
		t.Errorf("got %q, wanted %q", telescope.Alpaca.ErrorMessage, want.String())
	}
}

func TestNewTelescopeSetUTCDate(t *testing.T) {
	var UTCDate time.Time = time.Now().UTC()

	var err = telescope.SetUTCDate(UTCDate)

	if err != nil {
		t.Errorf("got %q", err)
	}
}

// Live Telescope Controls

// Slew to Altitude / Azimuth
func TestNewTelescopeSetSlewToAltAz(t *testing.T) {
	var err = telescope.SetSlewToAltAz(45.0, 90.0)

	if err != nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeSetSlewToAltAzInValidAltitude(t *testing.T) {
	var err = telescope.SetSlewToAltAz(-91, 90.0)

	if err == nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeSetSlewToAltAzInValidAzimuth(t *testing.T) {
	var err = telescope.SetSlewToAltAz(45.0, 361)

	if err == nil {
		t.Errorf("got %q", err)
	}
}

// Slew To Equatorial Coordinates

func TestNewTelescopeSetSlewToCoordinates(t *testing.T) {
	var err = telescope.SetSlewToCoordinates(45.0, 90.0)

	if err != nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeSetSlewToCoordinatesInValiDeclination(t *testing.T) {
	var err = telescope.SetSlewToCoordinates(-91, 90.0)

	if err == nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeSetSlewToCoordinatesInValidRightAscension(t *testing.T) {
	var err = telescope.SetSlewToCoordinates(45.0, 361)

	if err == nil {
		t.Errorf("got %q", err)
	}
}

// Slew To Equatorial Coordinates Target

func TestNewTelescopeSetSlewToTarget(t *testing.T) {
	var err = telescope.SetSlewToTarget()

	if err != nil {
		t.Errorf("got %q", err)
	}
}

// Slew to Altitude / Azimuth Async

func TestNewTelescopeSetSlewToAltAzAsync(t *testing.T) {
	var err = telescope.SetSlewToAltAzAsync(45.0, 90.0)

	if err != nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeSetSlewToAltAzAsyncInValidAltitude(t *testing.T) {
	var err = telescope.SetSlewToAltAzAsync(-91, 90.0)

	if err == nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeSetSlewToAltAzAsyncInValidAzimuth(t *testing.T) {
	var err = telescope.SetSlewToAltAzAsync(45.0, 361)

	if err == nil {
		t.Errorf("got %q", err)
	}
}

// Slew To Equatorial Coordinates Async

func TestNewTelescopeSetSlewToCoordinatesAsync(t *testing.T) {
	var err = telescope.SetSlewToCoordinatesAsync(45.0, 90.0)

	if err != nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeSetSlewToCoordinatesAsyncInValiDeclination(t *testing.T) {
	var err = telescope.SetSlewToCoordinatesAsync(-91, 90.0)

	if err == nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeSetSlewToCoordinatesAsyncInValidRightAscension(t *testing.T) {
	var err = telescope.SetSlewToCoordinatesAsync(45.0, 361)

	if err == nil {
		t.Errorf("got %q", err)
	}
}

// Slew To Equatorial Coordinates Target Async

func TestNewTelescopeSetSlewToTargetAsync(t *testing.T) {
	var err = telescope.SetSlewToTargetAsync()

	if err != nil {
		t.Errorf("got %q", err)
	}
}

func TestNewTelescopeSetUnPark(t *testing.T) {
	var err = telescope.SetUnPark()

	if err != nil {
		t.Errorf("got %q", err)
	}
}
