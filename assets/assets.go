package assets

import (
	"bytes"
	"image"
	_ "image/png" //Imported to handle PNG images
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

var (
	// RunnerImage is an animated PNG of a running man from ebiten examples
	RunnerImage *ebiten.Image
	TilesImage *ebiten.Image
)

// LoadAssets creates ebiten.Images from image assets
func LoadAssets() {
	runnerImg, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}
	RunnerImage = ebiten.NewImageFromImage(runnerImg)

	tilesImg, _, err := image.Decode(bytes.NewReader(images.Tiles_png))
	if err != nil {
		log.Fatal(err)
	}
	TilesImage = ebiten.NewImageFromImage(tilesImg)
}
