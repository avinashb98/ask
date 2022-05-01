package lib

import (
	"fmt"
	"net/url"
)

type Request struct {
	URL     *url.URL
	Method  string
	Headers map[string]string
	Body    []byte
}

func (r *Request) Print() {
	fmt.Println("REQUEST")
	fmt.Println("----------------------")
	fmt.Println("URL")
	fmt.Println(r.URL.String())
	fmt.Println("----------------------")
	fmt.Println("Method")
	fmt.Println(r.Method)

	fmt.Println("----------------------")
	fmt.Println("Headers")
	for key, value := range r.Headers {
		fmt.Print(key, ": ", value, "\n")
	}
	fmt.Println("----------------------")
	fmt.Println("Body")
	fmt.Println(string(r.Body))
	fmt.Println("")
}
