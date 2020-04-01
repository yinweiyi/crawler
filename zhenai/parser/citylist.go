package parser

import (
	"crawler/engine"
	"regexp"
)

var cityListCompile = regexp.MustCompile(`<a href="(.+www.zhenai.com/zhenghun/\w+)"[^>]*>([^<]+)</a>`)

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	matches := cityListCompile.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
	}

	return result
}
