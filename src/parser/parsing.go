package parser

import (
	"fetcher"
)

func Parsing(preq ParseRequest) []ParseRequest {

	var (
		set  []ParseRequest
		data []byte
	)

	data = fetcher.Fetching(preq.URL)

	set = preq.ParseFunc(data)

	return set
}
