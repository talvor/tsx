package tsx

import (
	"errors"
	"image"
	"path"
)

var (
	ErrTileTypeNotFound  = errors.New("tsx: tile type not found")
	ErrTileIDOutOfBounds = errors.New("tsx: tile id out of bounds")
)

type Tileset struct {
	baseDir    string
	Source     string
	Name       string `xml:"name,attr"`
	TileWidth  int    `xml:"tilewidth,attr"`
	TileHeight int    `xml:"tileheight,attr"`
	TileCount  int    `xml:"tilecount,attr"`
	Spacing    int    `xml:"spacing,attr"`
	Margin     int    `xml:"margin,attr"`
	Columns    int    `xml:"columns,attr"`
	Image      Image  `xml:"image"`
	Tiles      []Tile `xml:"tile"`
}

func (ts *Tileset) GetTileRect(tileID uint32) (image.Rectangle, error) {
	if tileID >= uint32(ts.TileCount) {
		return image.Rectangle{}, ErrTileIDOutOfBounds
	}

	tilesetColumns := ts.Columns

	if tilesetColumns == 0 {
		tilesetColumns = ts.Image.Width / (ts.TileWidth + ts.Spacing)
	}

	x := int(tileID) % tilesetColumns
	y := int(tileID) / tilesetColumns

	xOffset := int(x)*ts.Spacing + ts.Margin
	yOffset := int(y)*ts.Spacing + ts.Margin

	rect := image.Rect(x*ts.TileWidth+xOffset,
		y*ts.TileHeight+yOffset,
		(x+1)*ts.TileWidth+xOffset,
		(y+1)*ts.TileHeight+yOffset)
	return rect, nil
}

func (ts *Tileset) GetTileByType(tileType string) (*Tile, error) {
	for _, t := range ts.Tiles {
		if t.Type == tileType {
			return &t, nil
		}
	}

	return nil, ErrTileTypeNotFound
}

func (ts *Tileset) decodeImage() {
	if ts.Image.Source == "" {
		return
	}

	ts.Image.Source = path.Join(ts.baseDir, ts.Image.Source)
}

type Image struct {
	Source string `xml:"source,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
}

type Tile struct {
	ID        uint32    `xml:"id,attr"`
	Type      string    `xml:"type,attr"`
	Animation Animation `xml:"animation"`
}

type Animation struct {
	Frames []Frame `xml:"frame"`
}

type Frame struct {
	ID       uint32 `xml:"tileid,attr"`
	Duration string `xml:"duration,attr"`
}
