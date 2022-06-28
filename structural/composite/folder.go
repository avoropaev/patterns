package composite

type Folder struct {
	name   string
	childs []Component
}

func NewFolder(name string) *Folder {
	return &Folder{
		name: name,
	}
}

func (f *Folder) Add(child Component) {
	f.childs = append(f.childs, child)
}

func (f *Folder) Name() string {
	return f.name
}

func (f *Folder) Child() []Component {
	return f.childs
}

func (f *Folder) Print(prefix string) string {
	result := prefix + "/" + f.Name() + "\n"

	for _, val := range f.Child() {
		result += val.Print(prefix + "/" + f.Name())
	}

	return result
}
