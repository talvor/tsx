package renderer

import (
	"fmt"

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

func (er *Renderer) DrawTile(ts *tsx.Tileset, tileId uint32, opts *DrawOptions) error {
	img, err := er.loadTilesetImage(ts)
	if err != nil {
		return err
	}

	rect, err := ts.GetTileRect(tileId)
	if err != nil {
		return fmt.Errorf("failed to get tile rect for tile %d in tileset %s: %w", tileId, ts.Name, err)
	}

	img = transformImage(img.SubImage(rect).(*ebiten.Image), opts)

	opts.Screen.DrawImage(img, opts.Op)

	return nil
}

func (er *Renderer) DrawTileWithSource(tilesetSource string, tileId uint32, opts *DrawOptions) error {
	ts, err := er.TilesetManager.GetTilesetBySource(tilesetSource)
	if err != nil {
		return err
	}

	return er.DrawTile(ts, tileId, opts)
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
