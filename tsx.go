package tsx

import (
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
)

// LoadReader function loads tileset in TSX format from io.Reader
// baseDir is used for loading additional tile data, current directory is used if empty
func tsxReader(source string, r io.Reader) (*Tileset, error) {
	d := xml.NewDecoder(r)

	baseDir := filepath.Dir(source)
	ts := &Tileset{
		baseDir: baseDir,
		Source:  source,
	}
	if err := d.Decode(ts); err != nil {
		return nil, err
	}

	ts.decodeImage()

	return ts, nil
}

// LoadFile function loads tileset in TSX format from file
func LoadFile(fileName string) (*Tileset, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return tsxReader(fileName, f)
}
