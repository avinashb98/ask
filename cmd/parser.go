package cmd

import (
	"encoding/json"
	"github.com/avinashb98/ask/lib"
	"net/url"
	"strings"
)

func ParseInputToRequest(i Input) (*lib.Request, error) {
	var request lib.Request

	request.Headers = parseHeader(i.Headers)

	body, err := parseBody(i.Body)
	if err != nil {
		return nil, err
	}
	request.Body = body

	request.Method = parseMethod(i.Method)

	parsedUrl, err := parseURL(i.URL)
	if err != nil {
		return nil, err
	}
	request.URL = parsedUrl

	return &request, nil
}

func parseURL(inputURL string) (*url.URL, error) {
	return url.Parse(inputURL)
}

func parseMethod(inputMethod string) string {
	withoutSpace := strings.TrimSpace(inputMethod)
	return strings.ToUpper(withoutSpace)
}

func parseBody(inputBody string) ([]byte, error) {
	b, err := json.Marshal(inputBody)

	if err != nil {
		return nil, err
	}
	return b, nil
}

func parseHeader(inputHeader string) map[string]string {
	segments := strings.Split(inputHeader, " ")

	var trimmedSegments []string
	for _, segment := range segments {
		if segment != "" {
			trimmedSegments = append(trimmedSegments, segment)
		}
	}

	headers := map[string]string{}

	for _, segment := range trimmedSegments {
		tokens := strings.Split(segment, ":")
		headers[tokens[0]] = tokens[1]
	}

	return headers
}
