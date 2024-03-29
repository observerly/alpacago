package alpacago

import (
	"math"
	"testing"
)

var conditions = NewObservingConditions(65535, false, "100.69.47.32", "", -1, 0)

func TestNewObservingConditionsBaseURL(t *testing.T) {
	conditions := NewObservingConditions(65535, false, "", "0.0.0.0", 8000, 0)

	var got string = conditions.Alpaca.UrlBase
	var want string = "http://0.0.0.0:8000"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewObservingConditionsBaseURLForHost(t *testing.T) {
	var got string = conditions.Alpaca.UrlBase
	var want string = "http://100.69.47.32"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewObservingConditionsClientID(t *testing.T) {
	conditions := NewObservingConditions(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = conditions.Alpaca.ClientId
	var want uint32 = 65535

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewObservingConditionsTransactionID(t *testing.T) {
	conditions := NewObservingConditions(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = conditions.Alpaca.TransactionId
	var want uint32 = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewObservingConditionsDeviceNumber(t *testing.T) {
	conditions := NewObservingConditions(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint = conditions.DeviceNumber
	var want uint = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewObservingConditionsSetConnectedOn(t *testing.T) {
	var err = conditions.SetConnected(true)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsIsConnectedOn(t *testing.T) {
	var got, err = conditions.IsConnected()

	var want = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsSetConnectedOff(t *testing.T) {
	var err = conditions.SetConnected(false)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsIsConnectedOff(t *testing.T) {
	var got, err = conditions.IsConnected()

	var want = false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t wanted %t", got, want)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsGetAveragePeriod(t *testing.T) {
	conditions.SetConnected(true)

	var got, err = conditions.GetAveragePeriod()

	var want float64 = 0.0

	if err != nil {
		t.Errorf("got %q, wanted %f", err, want)
	}

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %f", conditions.Alpaca.ErrorMessage, want)
	}
}

func TestNewObservingConditionsGetCloudCoverage(t *testing.T) {
	conditions.SetConnected(true)

	var got, err = conditions.GetCloudCover()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 || got > 100 {
		t.Errorf("got %f", got)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsGetDewPoint(t *testing.T) {
	conditions.SetConnected(true)

	var got, err = conditions.GetDewPoint()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < -273.15 || got > 1000 {
		t.Errorf("got %f", got)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsGetHumidity(t *testing.T) {
	conditions.SetConnected(true)

	var got, err = conditions.GetHumidity()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 || got > 100 {
		t.Errorf("got %f", got)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsGetPressure(t *testing.T) {
	conditions.SetConnected(true)

	var got, err = conditions.GetPressure()

	// The average pressure at mean sea-level (MSL) in the International
	// Standard Atmosphere (ISA) is 1013.25 hPa, or 1 atmosphere (atm),
	// or 29.92 inches of mercury.
	var want = 1020.500

	if err != nil {
		t.Errorf("got %q", err)
	}

	// This will test for realistic values of atmospheric pressure:
	if got < 0 || got > 10000 {
		t.Errorf("got %f", got)
	}

	// This will test for pressures near to the accepted value of 1013.25hPa:
	if math.Abs(got-want) > 1000 {
		t.Errorf("got %f, wanted %f", got, want)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsGetRainRate(t *testing.T) {
	conditions.SetConnected(true)

	var got, err = conditions.GetRainRate()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 || got > 100 {
		t.Errorf("got %f", got)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsGetSkyBrightness(t *testing.T) {
	conditions.SetConnected(true)

	var got, err = conditions.GetSkyBrightness()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 {
		t.Errorf("got %f", got)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsGetSkyQuality(t *testing.T) {
	conditions.SetConnected(true)

	var got, err = conditions.GetSkyQuality()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < -20 || got > 30 {
		t.Errorf("got %f", got)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsGetSkyTemperature(t *testing.T) {
	conditions.SetConnected(true)

	var got, err = conditions.GetSkyTemperature()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < -273.15 || got > 120 {
		t.Errorf("got %f", got)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsGetSeeingStarFWHM(t *testing.T) {
	conditions.SetConnected(true)

	var got, err = conditions.GetSeeingStarFWHM()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < -1000 || got > 1000 {
		t.Errorf("got %f", got)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsGetTemperature(t *testing.T) {
	conditions.SetConnected(true)

	var got, err = conditions.GetTemperature()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < -273.15 || got > 120 {
		t.Errorf("got %f", got)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsGetWindDirection(t *testing.T) {
	conditions.SetConnected(true)

	var got, err = conditions.GetWindDirection()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 || got > 360 {
		t.Errorf("got %f", got)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsGetWindGust(t *testing.T) {
	conditions.SetConnected(true)

	var got, err = conditions.GetWindGust()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 {
		t.Errorf("got %f", got)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsGetWindSpeed(t *testing.T) {
	conditions.SetConnected(true)

	var got, err = conditions.GetWindSpeed()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 {
		t.Errorf("got %f", got)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsGetSensorDescription(t *testing.T) {
	conditions.SetConnected(true)

	var _, err = conditions.GetSensorDescription("Pressure")

	if err != nil {
		t.Errorf("got %q", err)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsGetTimeSinceLastUpdate(t *testing.T) {
	conditions.SetConnected(true)

	var _, err = conditions.GetTimeSinceLastUpdate("Pressure")

	if err != nil {
		t.Errorf("got %q", err)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsSetRefresh(t *testing.T) {
	conditions.SetConnected(true)

	var err = conditions.SetRefresh()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if conditions.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}
