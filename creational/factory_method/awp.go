package factory_method

type GunAWP struct{}

func (p *GunAWP) Fire() string {
	return "BOOM"
}
