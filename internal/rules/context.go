package rules

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"sync"
)

type Context struct {
	Writer    http.ResponseWriter
	Request   *http.Request
	transport http.RoundTripper

	mu       sync.RWMutex
	respBody []byte // buffer for response mutation
	status   int
	headers  map[string]string
}

func NewContext(w http.ResponseWriter, r *http.Request, t http.RoundTripper) *Context {
	return &Context{
		Writer:    w,
		Request:   r,
		transport: t,
		headers:   make(map[string]string),
	}
}

// Read request body
func (c *Context) RequestBody() []byte {
	if c.Request.Body == nil {
		return nil
	}
	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewReader(body)) // reset for downstream
	return body
}

// Get response body buffer
func (c *Context) Body() []byte {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.respBody
}

// Set response body (for mutate action)
func (c *Context) SetBody(body []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.respBody = body
}

// Write response
func (c *Context) Write(status int, body string, headers map[string]string) {
	c.mu.Lock()
	c.respBody = []byte(body)
	c.status = status
	if headers != nil {
		for k, v := range headers {
			c.headers[k] = v
		}
	}
	c.mu.Unlock()

	for k, v := range headers {
		c.Writer.Header().Set(k, v)
	}
	c.Writer.WriteHeader(status)
	c.Writer.Write([]byte(body))
}

func (c *Context) Forward() {
	target := "http://" + c.Request.Host
	c.ForwardTo(target)
}

func (c *Context) ForwardTo(target string) {
	targetURL, err := url.Parse(target)
	if err != nil {
		http.Error(c.Writer, "invalid target", http.StatusInternalServerError)
		return
	}

	// Preserve original path & query
	targetURL.Path = c.Request.URL.Path
	targetURL.RawQuery = c.Request.URL.RawQuery

	// Clone body
	var bodyReader io.Reader
	if c.Request.Body != nil {
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		bodyReader = bytes.NewReader(bodyBytes)
		c.Request.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}

	req, err := http.NewRequest(c.Request.Method, targetURL.String(), bodyReader)
	if err != nil {
		http.Error(c.Writer, "forward request failed", http.StatusInternalServerError)
		return
	}

	// Copy headers
	for k, v := range c.Request.Header {
		for _, vv := range v {
			req.Header.Add(k, vv)
		}
	}

	resp, err := c.transport.RoundTrip(req)
	if err != nil {
		http.Error(c.Writer, "upstream error", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	// Store response body for possible mutation
	c.mu.Lock()
	c.respBody = respBody
	c.status = resp.StatusCode
	c.mu.Unlock()

	// Copy response headers
	for k, v := range resp.Header {
		for _, vv := range v {
			c.Writer.Header().Add(k, vv)
		}
	}

	c.Writer.WriteHeader(resp.StatusCode)
	c.Writer.Write(respBody)
}
