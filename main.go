package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/fatih/color"
)

type UrlFormatter struct {
	UrlColor            *color.Color
	QueryParameterColor *color.Color
	QueryValueColor     *color.Color
}

func NewUrlFormatter() *UrlFormatter {
	return &UrlFormatter{
		UrlColor:            color.New(color.FgMagenta),
		QueryParameterColor: color.New(color.FgWhite),
		QueryValueColor:     color.New(color.FgGreen),
	}
}

func (f *UrlFormatter) Marshal(url *url.URL) string {
	buf := bytes.Buffer{}

	buf.WriteString(f.UrlColor.Sprintf("%s://%s%s\n", url.Scheme, url.Host, url.Path))
	for k, v := range url.Query() {
		buf.WriteString(f.QueryParameterColor.Sprintf("%s: ", k))

		var vs []string
		for _, s := range v {
			buf, err := base64.URLEncoding.DecodeString(s)
			if err != nil {
				vs = append(vs, string(buf))
			}
		}
		buf.WriteString(f.QueryValueColor.Sprintf("%s", strings.Join(v, ",")))
		buf.WriteString("\n")
	}
	return string(buf.Bytes())
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("missing url argument")
	}

	rawurl := os.Args[1]
	url, err := url.Parse(rawurl)
	if err != nil {
		log.Fatal("invalid url", err)
	}

	formatter := NewUrlFormatter()
	fmt.Println(formatter.Marshal(url))
}
