package engine

type ParserFunc func(contents []byte, url string) ParseResult

type Request struct {
	Url       string
	ParseFunc ParserFunc
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Id      string
	Url     string
	Type    string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
