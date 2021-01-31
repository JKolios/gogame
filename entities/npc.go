package entities

import (
	// "log"
	"math/rand"

	"github.com/JKolios/gogame/consts"
	"github.com/hajimehoshi/ebiten/v2"
)

// NPC is an Entity representing a character other than the player
type NPC struct {
	BaseMoveable
	updateIntentTicks int
}

func (npc *NPC) Draw(screen *ebiten.Image) {
	// log.Println("Called NPC Draw")
	npc.BaseMoveable.Draw(screen)
}

func (npc *NPC) Update(otherEntities []Moveable) {
	// log.Println("Called NPC Update")
	npc.UpdateMoveState()
	npc.BaseMoveable.UpdatePosition(otherEntities)
	npc.BaseMoveable.UpdateAnimationState()
}

// UpdateMoveState adjusts the location and state of a Player based on game state and keyboard input
func (npc *NPC) UpdateMoveState() {
	// log.Println("Called NPC updateMoveState")
	if npc.updateIntentTicks == consts.UpdateNPCIntentInterval {
		npc.MoveState.Direction = rand.Intn(5)
		npc.MoveState.Facing = rand.Intn(5)
		npc.updateIntentTicks = 0
	} else {
		npc.updateIntentTicks++
	}	
}
