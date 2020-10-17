package parser

import (
	"regexp"
	"spider-zaw/engine"
)

const (
	cityRegex string = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)"[^>]*>([^<]+)</a>`
)

func ParseCityList(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(cityRegex)
	matches := compile.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))

		result.Request = append(result.Request, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
