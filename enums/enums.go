package enums

import "fmt"

type WeaponType int

const (
	Sword WeaponType = iota
	Axe
	Arrow
	WoodenStick
)

func (w WeaponType) String() string {
	switch w {
	case Sword:
		return "Sword"
	case Axe:
		return "Axe"
	case Arrow:
		return "Arrow"
	case WoodenStick:
		return "Wooden Stick"
	}
	return ""
}

func getDamaged(weaponType WeaponType) int {
	switch weaponType {
	case Sword:
		return 87
	case Axe:
		return 23
	case Arrow:
		return 29
	case WoodenStick:
		return 1
	default:
		panic("The weapon does not exists")
	}
}

func Enums() {
	fmt.Printf("Damage of %s is %d\n", Sword, getDamaged(Sword))
	fmt.Printf("Damage of %s is %d\n", Axe, getDamaged(Axe))
	fmt.Printf("Damage of %s is %d\n", Arrow, getDamaged(Arrow))
	fmt.Printf("Damage of %s is %d\n", WoodenStick, getDamaged(WoodenStick))
}
