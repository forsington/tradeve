package config

type Region struct {
	Id   int
	Name string
}

type Regions map[string]int

var (
	regions = Regions{
		"The Forge":   10000002,
		"Domain":      10000043,
		"Sinq Laison": 10000032,
		"Metropolis":  10000042,
		"Heimatar":    10000030,
	}
)
