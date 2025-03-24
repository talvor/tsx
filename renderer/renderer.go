package renderer

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pkg/errors"
	"github.com/talvor/tsx"
)

type Renderer struct {
	TilesetManager  *tsx.TilesetManager
	TilesetImageMap map[string]*ebiten.Image
}

func NewRenderer(tm *tsx.TilesetManager) *Renderer {
	return &Renderer{
		TilesetManager:  tm,
		TilesetImageMap: make(map[string]*ebiten.Image),
	}
}

func (er *Renderer) MakeSprite(tileset interface{}) SpriteDrawer {
	switch tileset.(type) {
	case string: // single part sprite
		return NewSimpleSprite(tileset.(string), er)
	case []string: // multi part sprite
		return NewCompoundSprite(tileset.([]string), er)
	}
	return nil
}

func (er *Renderer) DrawTilesetByName(name string, screen *ebiten.Image, op *ebiten.DrawImageOptions) error {
	ts, err := er.TilesetManager.GetTilesetByName(name)
	if err != nil {
		return err
	}

	img, err := er.loadTilesetImage(ts)
	if err != nil {
		return err
	}

	screen.DrawImage(img, op)

	return nil
}

func (er *Renderer) loadTilesetImage(ts *tsx.Tileset) (*ebiten.Image, error) {
	if img, ok := er.TilesetImageMap[ts.Name]; ok {
		return img, nil
	}

	img, _, err := ebitenutil.NewImageFromFile(ts.Image.Source)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load tileset image")
	}

	er.TilesetImageMap[ts.Name] = img
	return img, nil
}
