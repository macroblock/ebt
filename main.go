package main

// Copyright 2018 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/macroblock/ebt/game"
	"github.com/macroblock/garbage/utils"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

const (
	tileSize   = 16
	tileXNum   = 25
	layerWidth = 15
)

var (
	tiles      []*ebiten.Image
	curTile    int
	tilesImage *ebiten.Image
	bgRect     *ebiten.Image

	normalFont font.Face
	bigFont    font.Face

	scale = 10.0
	quads []*ebiten.Image
	quad0 *ebiten.Image
)

func init() {
	// Decode image from a byte slice instead of a file so that
	// this example works in any working directory.
	// If you want to use a file, there are some options:
	// 1) Use os.Open and pass the file to the image decoder.
	//    This is a very regular way, but doesn't work on browsers.
	// 2) Use ebitenutil.OpenFile and pass the file to the image decoder.
	//    This works even on browsers.
	// 3) Use ebitenutil.NewImageFromFile to create an ebiten.Image directly from a file.
	//    This also works on browsers.
	img, _, err := image.Decode(bytes.NewReader(images.Tiles_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	w, h := tilesImage.Size()
	tilesXNum := w / tileSize
	tilesYNum := h / tileSize
	for y := 0; y < tilesYNum; y++ {
		for x := 0; x < tilesXNum; x++ {
			sx := x * tileSize
			sy := y * tileSize
			tiles = append(tiles, tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image))
		}
	}
	bgRect, _ = ebiten.NewImage(tileSize+2, tileSize+2, ebiten.FilterNearest)
	bgRect.Fill(color.Black)
	fmt.Printf("tiles x-y: %v-%v\n", tilesXNum, tilesYNum)
	fmt.Printf("number of tiles: %v\n", len(tiles))
}

func init() {
	tt, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	normalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	bigFont = truetype.NewFace(tt, &truetype.Options{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

var field *game.Field

func init() {
	quad0, _ = ebiten.NewImage(1, 1, ebiten.FilterNearest)
	quad0.Fill(color.Black)
	for r := 0; r < 10; r++ {
		for g := 0; g < 10; g++ {
			for b := 0; b < 10; b++ {
				img, _ := ebiten.NewImage(1, 1, ebiten.FilterNearest)
				vr := uint8(75 + r*20)
				vg := uint8(75 + g*20)
				vb := uint8(75 + b*20)
				img.Fill(color.RGBA{vr, vg, vb, 255})
				quads = append(quads, img)
			}
		}
	}

	// field = game.NewFieldInt(100, 100)
	field = initField()
	fmt.Printf("field size: %v\n", field.Size())
}

func initField() *game.Field {
	size := game.NewPoint2i(40, 40)
	return game.Generate(size, size, 20, 20)
}

// repeatingKeyPressed return true when key is pressed considering the repeat state.
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

func update(screen *ebiten.Image) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		// ebitenutil.DebugPrint(screen, "You're pressing the 'UP' button.")
		return errors.New("regular termination")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		field = initField()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEqual) {
		scale *= 1.5
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyMinus) {
		scale /= 1.5
	}

	// if ebiten.IsKeyPressed(ebiten.KeyLeft) {
	if repeatingKeyPressed(ebiten.KeyLeft) {
		curTile--
	}
	// if ebiten.IsKeyPressed(ebiten.KeyRight) {
	if repeatingKeyPressed(ebiten.KeyRight) {
		curTile++
	}
	curTile = utils.Max(curTile, 0)
	curTile = utils.Min(curTile, len(tiles)-1)

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// const xNum = screenWidth / tileSize
	// for _, l := range layers {
	// 	for i, t := range l {
	// 		op := &ebiten.DrawImageOptions{}
	// 		op.GeoM.Translate(float64((i%layerWidth)*tileSize), float64((i/layerWidth)*tileSize))
	// 		op.GeoM.Scale(scale, scale)

	// 		// sx := (t % tileXNum) * tileSize
	// 		// sy := (t / tileXNum) * tileSize
	// 		// tile := tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image)
	// 		// screen.DrawImage(tile, op)
	// 		screen.DrawImage(tiles[t], op)
	// 		// r := image.Rect(sx, sy, sx+tileSize, sy+tileSize)
	// 		// op.SourceRect = &r
	// 		// screen.DrawImage(tilesImage, op)
	// 	}
	// }
	screen.Fill(color.RGBA{0, 0, 50, 255})
	for y, line := range field.Grid() {
		for x, val := range line {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			op.GeoM.Scale(scale, scale)
			t := quad0
			if val > -1 && val < len(quads) {
				t = quads[val]
			}
			screen.DrawImage(t, op)
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))

	// msg := fmt.Sprintf("Tile: %v", curTile)
	// text.Draw(screen, msg, normalFont, 0, 75, color.White)
	// // ebitenutil.DebugPrint(screen, fmt.Sprintf("\nTile: %v", curTile))
	// op := &ebiten.DrawImageOptions{}
	// scale := 6.0
	// x := 5.0
	// y := 40.0
	// op.GeoM.Scale(scale, scale)
	// op.GeoM.Translate(x-1, y-1)
	// screen.DrawImage(bgRect, op)
	// op.GeoM.Reset()
	// op.GeoM.Scale(scale, scale)
	// op.GeoM.Translate(x, y)
	// screen.DrawImage(tiles[curTile], op)

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Tiles (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
