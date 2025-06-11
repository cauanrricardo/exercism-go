// Package weather provides functionality to forecast weather conditions in Goblinocus.
package weather

// CurrentCondition stores the current weather condition (e.g. sunny, rainy).
var CurrentCondition string

// CurrentLocation stores the name of the city. (cazuza, Belo horioznte).
var CurrentLocation string

// Forecast returns a string with the weather forecast for a given city and condition.
func Forecast() string {
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
