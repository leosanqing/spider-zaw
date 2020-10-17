package parser

import (
	"spider-zaw/fetcher"
	"testing"
)

const (
	fileName   = "citylist_test_data.html"
	resultSize = 470
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", contents)

	list := ParseCityList(contents)
	if len(list.Requests) != resultSize {
		t.Errorf("result should have %d requests;"+
			"but had %d", resultSize, len(list.Requests))
	}

	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}

	for i, url := range expectedUrls {
		if list.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, list.Requests[i].Url)
		}
	}

	for i, city := range expectedCities {
		if list.Items[i].(string) != city {
			t.Errorf("expected city #%d: %s; but was %s", i, city, list.Items[i])
		}
	}
}
