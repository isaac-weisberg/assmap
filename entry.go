package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type asset struct {
	Path         string    `json:"path"`
	Version      int       `json:"version"`
	Modification time.Time `json:"mod"`
}

type assmap struct {
	Assets []asset `json:"assets"`
}

func (assmap *assmap) json() ([]byte, error) {
	return json.Marshal(assmap)
}

func (assmap *assmap) unjson(data []byte) error {
	return json.Unmarshal(data, assmap)
}

func (assmap *assmap) write(filepath string) error {
	var error error
	var data []byte
	data, error = assmap.json()
	if error != nil {
		return error
	}
	error = ioutil.WriteFile(filepath, data, 0644)
	return error
}

func (assmap *assmap) read(filepath string) error {
	var error error
	var data []byte
	data, error = ioutil.ReadFile(filepath)
	if error != nil {
		return error
	}
	return assmap.unjson(data)
}
