package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	FileName string
}

func createStorage[T any](filename string) *Storage[T] {
	return &Storage[T]{
		FileName: filename,
	}
}

func (store *Storage[T]) save(data T) error {
	file, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		return err
	}

	return os.WriteFile(store.FileName, file, 0644)
}

func (store *Storage[T]) load(data *T) error {
	file, err := os.ReadFile(store.FileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(file, data)

}
