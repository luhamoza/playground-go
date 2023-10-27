package types

import "fmt"

type Position struct {
	posX, posY int
}

type Entity struct {
	name    string
	id      string
	version string
	Position
}

type SpecialEntity struct {
	Entity
	specialField float64
}

type PlayerPos struct {
	x, y int
}

type Player struct {
	PlayerPos
}

func (a Player) Move(val, val2 int) {
	fmt.Println("The player is moving to position", val, "and position", val2)
	fmt.Println(PlayerPos{val, val2})
}

func Types() {
	e := SpecialEntity{
		specialField: 1.56,
		Entity: Entity{
			name: "luqman",
			id:   "345",
			Position: Position{
				posX: 3}}}
	e.posY, e.version = 2, "3.4"
	fmt.Printf("%+v\n", e)

	p := Player{}
	p.Move(3, 5)

}
