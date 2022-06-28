package prototype

type File struct {
	name string
}

func NewFile(name string) INode {
	return &File{
		name: name,
	}
}

func (p *File) GetName() string {
	return p.name
}

func (p *File) Clone() INode {
	return &File{p.name + "_clone"}
}
