package cmd

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Input struct {
	URL     string
	Method  string
	Headers string
	Body    string
}

func (i Input) Validate() error {
	var err error
	err = i.validateMethod()
	if err != nil {
		return err
	}

	err = i.validateURL()
	if err != nil {
		return err
	}

	err = i.validateBody()
	if err != nil {
		return err
	}

	err = i.validateHeaders()
	if err != nil {
		return err
	}
	return nil
}

func (i Input) validateURL() error {
	validURLProtocols := []string{"https", "http"}
	if i.URL == "" {
		return fmt.Errorf("url cannot be empty")
	}

	isValidProtocol := false
	for _, protocol := range validURLProtocols {
		if strings.HasPrefix(i.URL, protocol) {
			isValidProtocol = true
			break
		}
	}

	if !isValidProtocol {
		return fmt.Errorf("only %s url protocols are supported", strings.Join(validURLProtocols, ","))
	}

	return nil
}

func (i Input) validateBody() error {
	if i.Body == "" {
		return nil
	}
	if !json.Valid([]byte(i.Body)) {
		return fmt.Errorf("body is an invalid json")
	}

	return nil
}

func (i Input) validateMethod() error {
	validMethods := []string{"GET", "POST"}
	if i.Method == "" {
		return nil
	}

	isValidmethod := false
	for _, method := range validMethods {
		if strings.ToUpper(i.Method) == method {
			isValidmethod = true
			break
		}
	}

	if !isValidmethod {
		return fmt.Errorf("only %s http methods are supported", strings.Join(validMethods, ","))
	}

	return nil
}

func (i Input) validateHeaders() error {
	segments := strings.Split(i.Headers, " ")

	var trimmedSegments []string
	for _, segment := range segments {
		if segment != "" {
			trimmedSegments = append(trimmedSegments, segment)
		}
	}

	for _, segment := range trimmedSegments {
		err := validateHeaderSegment(segment)
		if err != nil {
			return err
		}
	}

	return nil
}

func validateHeaderSegment(segment string) error {
	subSegments := strings.Split(segment, ":")
	if len(subSegments) != 2 || subSegments[0] == "" || subSegments[1] == "" {
		return fmt.Errorf("header malformed")
	}
	return nil
}
