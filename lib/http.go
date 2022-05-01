package lib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func ExecuteRequest(req *Request) (*Response, error) {
	client := http.Client{}
	requestBody := bytes.NewBuffer(req.Body)
	request, err := http.NewRequest(req.Method, req.URL.String(), requestBody)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body %v", err)
	}
	for key, value := range req.Headers {
		request.Header.Add(key, value)
	}
	startTime := time.Now()
	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error executing request %v", err)
	}

	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)

	var response Response
	response.StatusCode = resp.StatusCode
	response.Request = req
	response.Headers = resp.Header
	response.Body = sb
	timeTaken := time.Since(startTime)
	response.TimeTaken = timeTaken.Milliseconds()

	return &response, nil
}
