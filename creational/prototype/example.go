package prototype

import "fmt"

func Example() {
	file1 := NewFile("file1")
	file2 := NewFile("file2")
	file3 := NewFile("file3")

	folder1 := &Folder{
		Name:     "folder1",
		Children: []INode{file1},
	}

	folder2 := &Folder{
		Name:     "folder2",
		Children: []INode{folder1, file2, file3},
	}

	fmt.Printf("Printing hierarchy for folder 2: \n\t%s\n", folder2.GetName())

	cloneFolder := folder2.Clone()

	fmt.Printf("Printing hierarchy for clone of folder 2: \n\t%s\n", cloneFolder.GetName())
}
