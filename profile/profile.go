package profile

type Profile struct {
	Label      string
	Default    string
	Collectors []Collector
}

func (p *Profile) GetValue() string {
	for _, col := range p.Collectors {
		if val := col.Collect(); val != "" {
			return val
		}
	}

	return ""
}

type Collector interface {
	Collect() string
}
