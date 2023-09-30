package Prototype

import (
	"fmt"
	"testing"
)

func TestClone(t *testing.T) {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")
	//   Folder2
	//    Folder1
	//        File1
	//    File2
	//    File3

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
	//   Folder2_clone
	//    Folder1_clone
	//        File1_clone
	//    File2_clone
	//    File3_clone
}
