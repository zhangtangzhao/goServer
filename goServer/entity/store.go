package entity

import (
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("exist")
)

type Book struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type Store interface {
	Create(*Book) error
	Get(string) (Book, error)
	Delete(string) error
}