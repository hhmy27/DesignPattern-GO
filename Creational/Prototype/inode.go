package Prototype

type Inode interface {
	print(string)
	clone() Inode
}
