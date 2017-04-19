package main

// File is a representation of a go source code file.
type File struct {
	Name string
	Data string
}

// Directory contains go source code files and subdirectories.
type Directory struct {
	Name           string
	Files          []*File
	Subdirectories []*Directory
}

// Branch is a representation of a branch within a user's GitHub repository.
type Branch struct {
	Name string
	Root *Directory
}

// Repository is a representation of a user's GitHub repository.
type Repository struct {
	Username string
	Name string
	Branches []*Branch
}