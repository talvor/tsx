package tsx

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	ErrTilesetManagerNotLoaded = errors.New("tsx: tileset manager not loaded")
	ErrTilesetNotFound         = errors.New("tsx: tileset not found")
)

type TilesetManager struct {
	BaseDir  string
	Tilesets map[string]*Tileset
	IsLoaded bool
}

func (tm *TilesetManager) GetTilesetBySource(source string) (*Tileset, error) {
	if !tm.IsLoaded {
		return nil, ErrTilesetManagerNotLoaded
	}
	for _, ts := range tm.Tilesets {
		if ts.Source == source {
			return ts, nil
		}
	}
	return nil, ErrTilesetNotFound
}

func (tm *TilesetManager) GetTilesetByName(name string) (*Tileset, error) {
	if !tm.IsLoaded {
		return nil, ErrTilesetManagerNotLoaded
	}

	if _, ok := tm.Tilesets[name]; !ok {
		return nil, ErrTilesetNotFound
	}

	return tm.Tilesets[name], nil
}

func (tm *TilesetManager) AddTileset(source string) error {
	ts, err := LoadFile(source)
	if err != nil {
		return err
	}

	tm.Tilesets[ts.Name] = ts

	return nil
}

func (tm *TilesetManager) DebugPrintTilesets() {
	for name := range tm.Tilesets {
		fmt.Println(name)
	}
}

func NewTilesetManager(baseDir string) *TilesetManager {
	tm := &TilesetManager{
		BaseDir:  baseDir,
		Tilesets: make(map[string]*Tileset),
		IsLoaded: false,
	}

	loadTilesets(tm)

	return tm
}

func loadTilesets(tm *TilesetManager) {
	tsxFiles, err := findTSXFiles(tm.BaseDir)
	if err != nil {
		return
	}

	for _, tsxFile := range tsxFiles {
		ts, err := LoadFile(tsxFile)
		if err != nil {
			panic(err)
		}

		tm.Tilesets[ts.Name] = ts
	}

	tm.IsLoaded = true
}

func findTSXFiles(dir string) ([]string, error) {
	var tsxFiles []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".tsx" {
			tsxFiles = append(tsxFiles, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return tsxFiles, nil
}
