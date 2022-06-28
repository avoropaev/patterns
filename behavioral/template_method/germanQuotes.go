package template_method

type GermanQuotes struct{}

func (q *GermanQuotes) Open() string {
	return "„"
}

func (q *GermanQuotes) Close() string {
	return "“"
}
