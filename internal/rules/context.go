package rules

import (
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Context struct {
	Writer    http.ResponseWriter
	Request   *http.Request
	transport http.RoundTripper
}

func NewContext(w http.ResponseWriter, r *http.Request, t http.RoundTripper) *Context {
	return &Context{
		Writer:    w,
		Request:   r,
		transport: t,
	}
}

func (c *Context) Forward() {
	target := "http://" + c.Request.Host
	c.ForwardTo(target)
}

func (c *Context) ForwardTo(target string) {
	u, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.Transport = c.transport
	proxy.ServeHTTP(c.Writer, c.Request)
}

func (c *Context) Write(status int, body string, headers map[string]string) {
	for k, v := range headers {
		c.Writer.Header().Set(k, v)
	}
	c.Writer.WriteHeader(status)
	io.WriteString(c.Writer, body)
}
