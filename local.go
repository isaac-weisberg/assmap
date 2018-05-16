package main

import "os"

func localMapExists() (bool, error) {
	var error error
	var info os.FileInfo
	info, error = os.Stat(mapPath)
	if error != nil {
		if os.IsNotExist(error) {
			return false, nil
		}
		return false, error
	}
	if info.IsDir() {
		return false, newAssmapError("Asset is somehow a directory")
	}
	return true, nil
}

func localMapRead() (*assmap, error) {
	var assmap *assmap
	error := assmap.read(mapPath)
	return assmap, error
}
