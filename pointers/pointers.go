package pointers

import "fmt"

type Player struct {
	HP int
}

/*
func TakeDamage(player *Player, amount int) {
	player.HP -= amount
	fmt.Println("Player is taking damage. Current HP is =>", player.HP)
} */

// TakeDamage function receiver
func (p *Player) TakeDamage(amount int) {
	p.HP -= amount
	fmt.Println("Player is taking damage. Current HP is =>", p.HP)
}

func Pointer() {
	player := &Player{100}
	player.TakeDamage(10)
	fmt.Printf("current player HP:%d\n", player.HP)
}
