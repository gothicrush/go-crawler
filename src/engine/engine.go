package engine

import (
	"parser"
)

func Run(startPreq parser.ParseRequest) {

	var (
		taskqueue []parser.ParseRequest
		req       parser.ParseRequest
	)

	taskqueue = append(taskqueue, startPreq)

	for len(taskqueue) > 0 {

		req = taskqueue[0]
		taskqueue = taskqueue[1:]

		taskqueue = append(taskqueue, parser.Parsing(req)...)
	}
}
