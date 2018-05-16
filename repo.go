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
	return nil
}

func mapperIncremental() error {
	return nil
}
