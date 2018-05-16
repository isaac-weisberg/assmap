package main

const (
	mapPath = "./.assmap"
)

func mapperMain() error {
	var error error
	var exists bool
	exists, error = localMapExists()
	if error != nil {
		return error
	}
	if !exists {
		return mapperInitial()
	}
	return mapperIncremental()
}

func mapperInitial() error {
	var error error
	var assmap *assmap
	assmap, error = mapMake(".")
	if error != nil {
		return nil
	}
	return assmap.write(mapPath)
}

func mapperIncremental() error {
	var error error
	var localMap *assmap
	error = localMap.read(mapPath)
	if error != nil {
		return error
	}

	return nil
}
