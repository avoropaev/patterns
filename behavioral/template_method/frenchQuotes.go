package template_method

type FrenchQuotes struct{}

func (q *FrenchQuotes) Open() string {
	return "«"
}

func (q *FrenchQuotes) Close() string {
	return "»"
}
