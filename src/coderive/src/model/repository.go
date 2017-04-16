package main

import (
	"fmt"
)

type File struct {
	Name string
	Data []byte
}

// Directory
type Directory struct {
	Name string
	Files []*File
	Directories []*Directory
}

// Branch is a representation of a branch within a user's Github repository.
type Branch struct {
	Name string
	Directories []*Directory
}

// Repository is a representation of a user's Github repository.
type Repository struct {
	User string
	Name string
	Branches []*Branch
}

func main() {
	fmt.Println("Hello!")
}