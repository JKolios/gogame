package entities

import "github.com/hajimehoshi/ebiten/v2"

//Moveable represents any Entity that can Move and Animate
type Moveable interface {
	Update(otherEntities []Moveable)
	Draw(screen *ebiten.Image)
	CheckCollision(int, int) bool
}
