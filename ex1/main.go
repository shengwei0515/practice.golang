package main

import (
	"flag"

	"github.com/shengwei0515/parctice.golang/ex1/lib/parseUrl"
)

func main() {
	theUrl := flag.String("u", "https://example.com:443/path?arg=value", "url to parse")
	flag.Parse()
	parseUrl.ParseUrl(*theUrl)
}
