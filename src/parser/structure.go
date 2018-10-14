package parser

type ParseRequest struct {
	URL       string
	ParseFunc func([]byte) []ParseRequest
}
