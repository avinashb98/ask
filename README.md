# ask
CLI-Based HTTP Client in Go

# Requirements

- A user should be able to give a URL for GET requests and a URL + body for POST requests.
- Users should be able to set up custom headers for every request
- The output from the API can be assumed to be JSON
- GET and POST requests
- For JSON requests, the request body is assumed to be JSON

# Usage

```
ask [METHOD] (url) [flags]
```
## Get

```
ask http://httpbin.org/get
ask GET http://httpbin.org/get

// with headers
ask http://httpbin.org/get -H "accept:appplication/json x-user-id: 12345"
```
## Post

```
ask POST http://httpbin.org/post

// with headers
ask POST http://httpbin.org/post -H "accept:appplication/json x-user-id: 12345"

// with body
ask POST http://httpbin.org/post -d "{\"key\": \"value\"}"
```
## Output

```
ask http://httpbin.org/get

Response

Headers
--
--
--

Status: 200

{
  "args": {}, 
  "headers": {
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9", 
    "Accept-Encoding": "gzip, deflate", 
    "Accept-Language": "en-US,en;q=0.9,hi;q=0.8", 
    "Host": "httpbin.org", 
    "Upgrade-Insecure-Requests": "1", 
    "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36", 
    "X-Amzn-Trace-Id": "Root=1-626eac3d-58547338525f9bb0656d89f6"
  }, 
  "origin": "223.236.211.64", 
  "url": "http://httpbin.org/get"
}
```
## Validation

- url is a required parameter
- if the method isn’t passed, it is by default a GET request
- method should be GET and POST - lowercase or uppercase
- url should have `http` or `https` as the protocol
- it should be valid url - verify it via regex
- header should not be malformed - `“key:value key:value”`
- body should be a valid json

# LLD

```
type Request struct {
  URL *net.URL
  Headers map[string]string
  Body interface{}
  Timeout *time.Duration
}

func (r *Reqeust) Print() {}

type Response struct {
  Request Request
  StatusCode int
  TimeTaken float64
  Body interface{}
  Headers map[string]string
}

func (r *Response) Print() {}

func main() {
  cmd.Execute()
}

// root cmd

func Execute() {
  validation()

  parsing()

  execution()
}
```

# Good To Have Features

- accept a verbosity flag `-v -vv -vvv` - levels of information printed
- accept body in formats other than JSON
- prompt to display the request information because making the request, controlled via a flag or based on methods : (POST, DELETE)
- accept the url without a protocol - e.g. `httpbin.org/get`, default would be `http`
- accept body from a file
