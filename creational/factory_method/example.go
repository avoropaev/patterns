package factory_method

import "fmt"

func Example() {
	factory := NewGunCreator()

	ak47, err := factory.CreateGun(AK47)
	if err != nil {
		fmt.Printf("failed to create ak47, error: %v", err)
	}

	fmt.Printf("AK47 FIRE: %s", ak47.Fire())
	fmt.Println()

	m4a1s, err := factory.CreateGun(M4A1S)
	if err != nil {
		fmt.Printf("failed to create m4a1s, error: %v", err)
	}

	fmt.Printf("M4A1S FIRE: %s", m4a1s.Fire())
	fmt.Println()

	awp, err := factory.CreateGun(AWP)
	if err != nil {
		fmt.Printf("failed to create awp, error: %v", err)
	}

	fmt.Printf("AWP FIRE: %s", awp.Fire())
	fmt.Println()
}
