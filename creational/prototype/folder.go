package prototype

import "strings"

type Folder struct {
	Name     string
	Children []INode
}

func (f *Folder) GetName() string {
	builder := &strings.Builder{}
	builder.WriteString(f.Name)

	for _, children := range f.Children {
		builder.WriteString("\n\t\t")
		builder.WriteString(children.GetName())
	}

	return builder.String()
}

func (f *Folder) Clone() INode {
	cloneFolder := &Folder{Name: f.Name + "_clone"}

	tempChildren := make([]INode, 0, len(f.Children))
	for _, children := range f.Children {
		tempChildren = append(tempChildren, children.Clone())
	}

	cloneFolder.Children = tempChildren

	return cloneFolder
}
