package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/JKolios/gogame/assets"
	"github.com/JKolios/gogame/consts"
	"github.com/JKolios/gogame/entities"
)

// Game implements ebiten.Game interface.
type Game struct {
	gameState int
	player    *entities.Player
	entities  []entities.Moveable
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {

	if err := g.HandleQuit(); err != nil {
		return err
	}

	g.HandlePauseUnpause()

	if g.gameState == consts.MainGameState {
		for entityIndex := range g.entities {
			// log.Printf("entityIndex: %+v", entityIndex)
			// log.Printf("updated entity: %+v", g.entities[entityIndex])
			// log.Printf("entities: %+v", g.entities)
			var entityCopy []entities.Moveable = make([]entities.Moveable, len(g.entities))
			copy(entityCopy, g.entities)
			// log.Printf("copied: %+v", copied)
			// log.Printf("entityCopy: %+v", entityCopy)
			otherEntities := append(entityCopy[:entityIndex], entityCopy[entityIndex+1:]...)
			// log.Printf("entities: %+v", g.entities)
			// log.Printf("otherEntities: %+v", otherEntities)
			g.entities[entityIndex].Update(otherEntities)
		}
	}

	// log.Fatal("bye")

	if ebiten.IsKeyPressed(ebiten.KeyN) {
		g.entities = append(g.entities, &entities.NPC{BaseMoveable: entities.BaseMoveable{XPosition: rand.Intn(1024), YPosition: rand.Intn(768)}})
	}
	return nil
}

// HandlePauseUnpause detects if the game should move between the main state and the paused state
func (g *Game) HandlePauseUnpause() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if g.gameState == consts.MainGameState {
			g.gameState = consts.PausedGameState
		} else if g.gameState == consts.PausedGameState {
			g.gameState = consts.MainGameState
		}
	}
}

// HandleQuit detects if the game should quit
func (g *Game) HandleQuit() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("game ended by player")
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS:%.0f", math.Round(ebiten.CurrentFPS())))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS:%.0f", math.Round(ebiten.CurrentTPS())), 0, 20)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Game State:%v", g.gameState), 0, 40)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Player Move Direction:%v", g.player.MoveState.Direction), 0, 60)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Player Facing:%v", g.player.MoveState.Facing), 0, 80)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("PlayerAnimation Frame Counter:%v", g.player.MoveState.AnimationCounter), 0, 100)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Entity 0:%+v", g.entities[0]), 0, 120)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Entity 1:%+v", g.entities[1]), 0, 140)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Entity Count: %v", len(g.entities)), 0, 160)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	op.GeoM.Scale(float64(consts.ScreenWidth)/16, float64(consts.ScreenHeight)/16)
	screen.DrawImage(assets.TilesImage, op)

	for _, entity := range g.entities {
		entity.Draw(screen)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {

	// Init the RNG
	rand.Seed(time.Now().Unix())

	player := &entities.Player{BaseMoveable: entities.BaseMoveable{XPosition: 100, YPosition: 100}}
	npc := &entities.NPC{BaseMoveable: entities.BaseMoveable{XPosition: 400, YPosition: 400}}
	npc2 := &entities.NPC{BaseMoveable: entities.BaseMoveable{XPosition: 600, YPosition: 600}}
	game := &Game{
		gameState: consts.MainGameState,
		player:    player,
		entities:  []entities.Moveable{player, npc, npc2},
	}

	assets.LoadAssets()

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(consts.ScreenWidth, consts.ScreenHeight)
	ebiten.SetWindowTitle("A Game")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
