package lib

import (
	"fmt"
	"net/http"
)

type Response struct {
	Request    *Request
	StatusCode int
	TimeTaken  int64
	Body       interface{}
	Headers    http.Header
}

func (r *Response) Print() {
	fmt.Println("RESPONSE")
	fmt.Println("Status: ", r.StatusCode)

	fmt.Println("----------------------")
	fmt.Println("Headers")
	for key, value := range r.Headers {
		fmt.Print(key, ": ", value, "\n")
	}
	fmt.Println("----------------------")
	fmt.Println("Body")
	fmt.Println(r.Body.(string))
	fmt.Println("")
}
