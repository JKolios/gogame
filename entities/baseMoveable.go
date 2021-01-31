package entities

import (
	"image"
	"math"

	"github.com/JKolios/gogame/consts"

	"github.com/JKolios/gogame/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

// BaseMoveable describes any Entity in the game state that is able to Move
type BaseMoveable struct {
	MoveState
	XPosition int
	YPosition int
	state     map[string]interface{}
}

// MoveState reflects the movement state of an Entity
type MoveState struct {
	Direction        int
	Facing           int
	AnimationCounter int
}

func (baseMoveable *BaseMoveable) Update(otherMoveables []Moveable) {
	baseMoveable.UpdateMoveState()
	baseMoveable.UpdatePosition(otherMoveables)
	baseMoveable.UpdateAnimationState()
}

// updateMoveState adjusts the moveState of a Moveable based on game state and keyboard input
func (baseMoveable *BaseMoveable) UpdateMoveState() {
	// log.Println("Called baseMoveable updateMoveState")
}

// updatePosition adjusts the location of a Moveable based on its moveState
func (baseMoveable *BaseMoveable) UpdatePosition(otherMoveables []Moveable) {

	targetXposition, targetYposition := baseMoveable.XPosition, baseMoveable.YPosition

	switch baseMoveable.MoveState.Direction {
	case consts.MoveDirectionRight:
		targetXposition += consts.MoveSpeed
	case consts.MoveDirectionUp:
		targetYposition -= consts.MoveSpeed
	case consts.MoveDirectionLeft:
		targetXposition -= consts.MoveSpeed
	case consts.MoveDirectionDown:
		targetYposition += consts.MoveSpeed
	}

	if targetXposition >= consts.ScreenWidth-consts.RightPadding {
		targetXposition = consts.ScreenWidth - consts.RightPadding - 1
	}

	if targetXposition <= consts.LeftPadding {
		targetXposition = consts.LeftPadding + 1
	}

	if targetYposition >= consts.ScreenHeight-consts.BottomPadding {
		targetYposition = consts.ScreenHeight - consts.BottomPadding - 1
	}

	if targetYposition <= consts.TopPadding {
		targetYposition = consts.TopPadding + 1
	}

	//Detect collisions with other BaseMoveables, do not move if so
	for _, otherMoveable := range otherMoveables {
		if otherMoveable.CheckCollision(targetXposition, targetYposition) {
			return
		}
	}

	baseMoveable.XPosition = targetXposition
	baseMoveable.YPosition = targetYposition
}

// updateAnimationState adjusts the state of the movement animation of the Moveable
func (baseMoveable *BaseMoveable) UpdateAnimationState() {
	if baseMoveable.MoveState.Direction != consts.MoveDirectionNone {
		baseMoveable.MoveState.AnimationCounter++
	} else {
		baseMoveable.MoveState.AnimationCounter = 0
	}
}

// Draw draws the entity on the main game image
func (baseMoveable *BaseMoveable) Draw(screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}

	if baseMoveable.MoveState.Facing == consts.MoveDirectionLeft {
		drawOptions.GeoM.Scale(-1, 1)
		drawOptions.GeoM.Translate(consts.FrameWidth, 0)
	}
	drawOptions.GeoM.Translate(float64(baseMoveable.XPosition), float64(baseMoveable.YPosition))
	i := (baseMoveable.MoveState.AnimationCounter / 5) % consts.FrameNum
	sx, sy := consts.FrameOX+i*consts.FrameWidth, consts.FrameOY
	entityImage := assets.RunnerImage.SubImage(image.Rect(sx, sy, sx+consts.FrameWidth, sy+consts.FrameHeight)).(*ebiten.Image)
	screen.DrawImage(entityImage, drawOptions)
}

// CheckCollision determines if a basMoveable collides with something in position xPosition, yPosition
func (baseMoveable BaseMoveable) CheckCollision(xPosition, yPosition int) bool {

	if distance(baseMoveable.XPosition, xPosition, baseMoveable.YPosition, yPosition) < consts.MinimumCollisionDistance {
		return true
	}
	return false
}

// Distance finds the length of the hypotenuse between two points.
// Forumula is the square root of (x2 - x1)^2 + (y2 - y1)^2
func distance(x0, x1, y0, y1 int) float64 {
	distance := math.Sqrt(math.Pow(float64(x1-x0), 2) + math.Pow(float64(y1-y0), 2))
	return distance
}
