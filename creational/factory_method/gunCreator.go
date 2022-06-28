package factory_method

import (
	"fmt"
)

type GunCreator interface {
	CreateGun(gunType GunType) (Gun, error)
}

type Gun interface {
	Fire() string
}

type gunCreator struct{}

func NewGunCreator() GunCreator {
	return &gunCreator{}
}

func (p *gunCreator) CreateGun(gunType GunType) (Gun, error) {
	switch gunType {
	case AK47:
		return &GunAK47{}, nil
	case M4A1S:
		return &GunM4A1S{}, nil
	case AWP:
		return &GunAWP{}, nil
	}

	return nil, fmt.Errorf("unexpected gun type: %s", gunType)
}
