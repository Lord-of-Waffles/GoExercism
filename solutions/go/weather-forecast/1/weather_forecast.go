// Package weather contains two variables, CurrentCondition and CurrentLocation.
// it also contains a function Forecast() that takes a city and condition as parameters.
// the function sets CurrentCondition and CurrentLocations values as the passed parameters.
// then returns a string explaining the current weather in the given the location.
package weather

var (
    // CurrentCondition represents weather condition.
	CurrentCondition string
    // CurrentLocation represents the specified city.
	CurrentLocation  string
)

// Forecast function tells the current weather condition in the provided city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
