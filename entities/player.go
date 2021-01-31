package entities

import (
	// "log"

	"github.com/JKolios/gogame/consts"
	"github.com/hajimehoshi/ebiten/v2"
)

// Player is an Entity representing the player
type Player struct {
	BaseMoveable
}

func (player *Player) Update(otherEntities []Moveable) {
	// log.Println("Called player Update")
	player.UpdateMoveState()
	player.BaseMoveable.UpdatePosition(otherEntities)
	player.BaseMoveable.UpdateAnimationState()
}

func (player *Player) Draw(screen *ebiten.Image) {
	// log.Println("Called Player Draw")
	player.BaseMoveable.Draw(screen)
}

// UpdateMoveState adjusts the location and state of a Player based on game state and keyboard input
func (player *Player) UpdateMoveState() {
	// log.Println("Called player updateMoveState")
	player.MoveState.Direction = consts.MoveDirectionNone
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		player.MoveState.Direction = consts.MoveDirectionRight
		player.MoveState.Facing = consts.MoveDirectionRight
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		player.MoveState.Direction = consts.MoveDirectionDown
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		player.MoveState.Direction = consts.MoveDirectionLeft
		player.MoveState.Facing = consts.MoveDirectionLeft
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		player.MoveState.Direction = consts.MoveDirectionUp
	}

}
