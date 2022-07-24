// Package weather got a forecast methos that allows to get the current weather of a city.
package weather

// CurrentCondition represensts the current weather condition.
var CurrentCondition string
// CurrentLocation represensts the current weather location.
var CurrentLocation string


// Forecast got the city and the condition to return the forecast for a specific location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
