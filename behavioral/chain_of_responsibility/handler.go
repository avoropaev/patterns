package chain_of_responsibility

type Handler interface {
	SendRequest(message int) string
}
