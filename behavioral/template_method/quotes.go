package template_method

type QuotesInterface interface {
	Open() string
	Close() string
}

type Quotes struct {
	QuotesInterface
}

func (q *Quotes) Quotes(str string) string {
	return q.Open() + str + q.Close()
}

func NewQuotes(qt QuotesInterface) *Quotes {
	return &Quotes{qt}
}
