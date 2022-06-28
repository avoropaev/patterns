package composite

import "log"

func Example() {
	rootFolder := NewFolder("root")

	usrDir := NewFolder("usr")
	fileA := NewFile("A")

	rootFolder.Add(usrDir)
	rootFolder.Add(fileA)

	fileB := NewFile("B")

	usrDir.Add(fileB)

	log.Println(rootFolder.Print(""))
}
