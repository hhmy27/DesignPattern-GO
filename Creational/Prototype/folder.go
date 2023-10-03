package Prototype

import (
	"fmt"
)

type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, child := range f.children {
		child.print(indentation + indentation)
	}
}

func (f *Folder) clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	for _, child := range f.children {
		cloneFolder.children = append(cloneFolder.children, child.clone())
	}
	return cloneFolder
}
