package parser

import (
	"regexp"
	"spider-zaw/engine"
)

const (
	nameReg = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>[^<]([^<]+)</a>`
)

func ParseCity(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(nameReg)
	matches := compile.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User "+string(m[2]))

		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
