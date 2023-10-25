package play

import "fmt"

type Player struct {
	name   string
	health int
	hp     float64
}

type Weapon string

func (player Player) getHealth() int {
	return player.health
}

func getWeapon(w Weapon) Weapon {
	w += " Berapi"
	return w
}

func Play() {
	player := Player{"Luhamoza", 100, 55.6}
	fmt.Printf("This is a player's health: %d\n", player.getHealth())
	fmt.Printf("This is a player's weapon: %s\n", getWeapon("Senapang"))
}
