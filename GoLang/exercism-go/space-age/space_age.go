package space

type Planet string

const snInWorld float64 = 31557600

func Age(seconds float64, planet Planet) float64{
	var planetYear float64
	switch planet {
	case "Mercury":
		planetYear = seconds / (0.2408467 * snInWorld)
	case "Venus":
		planetYear = seconds / (0.61519726 * snInWorld)
	case "Earth":
		planetYear = seconds / (1 * snInWorld)
	case "Mars":
		planetYear = seconds / (1.8808158 * snInWorld)
	case "Jupiter":
		planetYear = seconds / (11.862615 * snInWorld)
	case "Saturn":
		planetYear = seconds / (29.447498 * snInWorld)
	case "Uranus":
		planetYear = seconds / (84.016846 * snInWorld)
	case "Neptune":
		planetYear = seconds / (164.79132 * snInWorld)
	}
	return planetYear
}
