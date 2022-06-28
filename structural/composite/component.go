package composite

type Component interface {
	Add(child Component)
	Name() string
	Child() []Component
	Print(prefix string) string
}
