package main

import (
	"os"
	"path/filepath"
)

func mapMake(rootDir string) (*assmap, error) {
	assetmap := &assmap{}

	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if path == rootDir {
			return nil
		}
		if err != nil {
			return nil
		}
		// TODO: encapsulate
		if gParameters.mapperIgnoreHidden {
			base := filepath.Base(path)
			if base[0:1] == "." {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		var asset = &asset{}

		if info.IsDir() {
			return nil
		}

		asset.Path = path
		asset.Version = 1
		asset.Modification = info.ModTime()

		assetmap.append(asset)
		return nil
	})

	return assetmap, nil
}
