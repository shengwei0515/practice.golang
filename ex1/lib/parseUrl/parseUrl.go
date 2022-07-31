package parseUrl

import (
	"fmt"
	"net/url"
)

func ParseUrl(theUrl string) {
	fmt.Println(theUrl)
	parsedUrl, err := url.Parse(theUrl)
	if err != nil {
		fmt.Printf("run failed with error: $s\n", err)
	}

	fmt.Printf("Schema: %s\n", parsedUrl.Scheme)
	fmt.Printf("Hostname: %s\n", parsedUrl.Hostname())
	fmt.Printf("Port: %s\n", parsedUrl.Port())
	fmt.Printf("Path: %s\n", parsedUrl.Path)
	fmt.Printf("Query String: %s\n", parsedUrl.RawQuery)
}
