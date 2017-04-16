package db

import (
	//"time"
)

type File struct {
	Name string
	Data []byte
}

type Directory struct {
	Name string
	Files []*File
	Directories []*Directory
}

type Branch struct {
	Name string
	Directories []*Directory
}

type Repository struct {
	User string
	Name string
	Branches []*Branch
}