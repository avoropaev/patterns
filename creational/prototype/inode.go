package prototype

type INode interface {
	Clone() INode
	GetName() string
}
