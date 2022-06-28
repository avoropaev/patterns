package factory_method

type GunM4A1S struct{}

func (p *GunM4A1S) Fire() string {
	return "tic tic"
}
