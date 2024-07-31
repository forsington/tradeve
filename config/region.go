package config

type Region struct {
	Id   int
	Name string
}

type Regions []*Region

var (
	regions = Regions{
		{
			Id:   10000002,
			Name: "The Forge",
		},
		{
			Id:   10000043,
			Name: "Domain",
		},
		{
			Id:   10000032,
			Name: "Sinq Laison",
		},
		{
			Id:   10000042,
			Name: "Metropolis",
		},
		{
			Id:   10000030,
			Name: "Heimatar",
		},
	}
)

func (r Regions) GetRegionByName(name string) *Region {
	for i := range r {
		if r[i].Name == name {
			return r[i]
		}
	}
	return nil
}

func (r Regions) GetRegionById(id int) *Region {
	for i := range r {
		if r[i].Id == id {
			return r[i]
		}
	}
	return nil
}
