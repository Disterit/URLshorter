package storage

import "errors"

var (
	ErrUrlNotFound = errors.New("URL not found")
	ErrURLExists   = errors.New("URL already exists")
)
