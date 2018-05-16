package main

import (
	"os"
	"path/filepath"
)

func mapMake(rootDir string) (*assmap, error) {
	assetmap := &assmap{}

	filepath.Walk(rootDir, func(filepath string, info os.FileInfo, err error) error {
		if filepath == rootDir {
			return nil
		}
		if err != nil {
			return nil
		}
		var asset = &asset{}

		asset.Path = filepath
		asset.Version = 1
		asset.Modification = info.ModTime()

		assetmap.append(asset)
		return nil
	})

	return assetmap, nil
}
