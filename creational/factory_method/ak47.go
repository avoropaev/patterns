package factory_method

type GunAK47 struct{}

func (p *GunAK47) Fire() string {
	return "tadam tadam"
}
