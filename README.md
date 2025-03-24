# tsx

A go module to read TSX tilesets from tiled map editor (https://www.mapeditor.org/)

## Reading individual tilesets

To read in a tileset use the `tsx.LoadFile` function

```golang
package main

import (
    "encoding/json"
    "fmt"

    "github.com/talvor/tsx"
)

func main() {
  tileset, err := tsx.LoadFile("~/Documents/tilesets/player.tsx")
  if err != nil {
    panic(err)
  }

  tsJSON, _ := json.Marshal(tileset)
  fmt.Println(string(tsJSON))
}
```

## Managing multiple tilesets using TilesetManager

To read in bulk tilesets use the `TilesetManager` struct.

```golang
package main

import "github.com/talvor/tsx"

func main() {
	// Create a new tileset manager and load all tilesets from the directory
	tsm := tsx.NewTilesetManager("/home/phillip/Documents/tilesets")
	tsm.AddTileset("/home/phillip/Documents/tilesets/player.tsx")

	tileset, _ := tsm.GetTilesetBySource("/home/phillip/Documents/tilesets/player.tsx")
	tileset, _ = tsm.GetTilesetByName("player")
}
```

## Using the renderer

The `tsx.renderer` works with the `TilesetManager` and the [ebitenengine](https://ebitengine.org/) 2D game engine to provide convenient methods for rendering
tilesets into the ebiten screen.

See `renderer/examples/main.go` for an example of using the renderer
