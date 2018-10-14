package main

import (
	"engine"
	"parser"
)

func main() {

	engine.Run(parser.ParseRequest{
		URL:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.CityListParseFunc,
	})
}
