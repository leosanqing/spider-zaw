package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

const (
	cityRegex string = `<a data-v-1573aa7c="" href="http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+"[^>]*>[^<]+</a>`
)

func main() {
	content, err := getWebContent("http://www.zhenai.com/zhenghun")
	if err ==nil{
		printCityList(content)
	}
}


func getWebContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return  nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return nil,err
	}

	e := determineEncoding(resp.Body)

	reader := transform.NewReader(resp.Body, e.NewDecoder())

	all, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	return all,nil
}

func determineEncoding(r io.Reader) encoding.Encoding{
	bytes, err := bufio.NewReader(r).Peek(1024)

	if err != nil{
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte){
	compile := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)"[^>]*>([^<]+)</a>`)
	//<a href="http://www.zhenai.com/zhenghun/aba" data-v-1573aa7c>阿坝</a>
		matches := compile.FindAllSubmatch(contents,-1)

	for _,m := range matches {
		fmt.Printf("City: %s, URL: %s\n",m[2],m[1])
		fmt.Println()
	}
	fmt.Printf("Matches found: %d\n",len(matches))
}

