package space

type Planet string

var planets = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

const EarthYearInSeconds = 31557600

func Age(s float64, p Planet) float64 {

	t := s / (EarthYearInSeconds * planets[p])
	return t
}
