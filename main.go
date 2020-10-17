package main

import (
	"spider-zaw/engine"
	"spider-zaw/zhenai/parser"
)

const (
	WebUrl = "http://www.zhenai.com/zhenghun"
)

func main() {
	engine.Run(engine.Request{
		Url:        WebUrl,
		ParserFunc: parser.ParseCityList,
	})
}
