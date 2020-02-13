package space

// Planet contains planet's name to
// identify orbit system
type Planet string

var earthOrbitTime float64 = 31557600

var orbitTimes = map[Planet]float64{
	Planet("Earth"):   earthOrbitTime,
	Planet("Mercury"): earthOrbitTime * 0.2408467,
	Planet("Venus"):   earthOrbitTime * 0.61519726,
	Planet("Mars"):    earthOrbitTime * 1.8808158,
	Planet("Jupiter"): earthOrbitTime * 11.862615,
	Planet("Saturn"):  earthOrbitTime * 29.447498,
	Planet("Uranus"):  earthOrbitTime * 84.016846,
	Planet("Neptune"): earthOrbitTime * 164.79132,
}

// Age calculates current age according to time lived
// in seconds evaluated on different orbit systems
func Age(secondsSinceBirth float64, planetName Planet) float64 {
	var ageInPlanet float64
	orbitTime := orbitTimes[planetName]
	ageInPlanet = secondsSinceBirth / orbitTime
	return ageInPlanet
}
