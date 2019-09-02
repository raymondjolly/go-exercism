package space

import "math"

type Planet string

//Age takes a planet name and the number of seconds to calculate
//the equivalent age give the respective orbit of the planet.
func Age(seconds float64, name Planet) float64 {
	switch name {
	case "Earth":
		return seconds/31557600
	case "Mercury":
		return seconds/(0.2408467*31557600)
	case "Venus":
		return seconds/(0.61519726*31557600)
	case "Mars":
		return seconds/(1.8808158*31557600)
	case "Jupiter":
		return seconds/(11.862615*31557600)
	case "Saturn":
		return seconds/(29.447498*31557600)
	case "Uranus":
		return seconds/(84.016846*31557600)
	case "Neptune":
		return seconds/(164.79132*31557600)
	default:
		return math.NaN()
	}
}







