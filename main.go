package main

import (
	"spider-zaw/engine"
	"spider-zaw/zhenai/parser"
)

const (
	webUrl = "http://www.zhenai.com/zhenghun"
)

func main() {
	engine.Run(engine.Request{
		Url:        webUrl,
		ParserFunc: parser.ParseCityList,
	})
}
