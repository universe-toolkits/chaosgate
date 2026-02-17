package proxy

import (
	"crypto/tls"
	"io"
	"net"
	"net/http"

	"github.com/universe-toolkits/chaosgate/internal/rules"
)

func (p *Proxy) handleHTTPS(w http.ResponseWriter, r *http.Request) {

	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "hijacking not supported", http.StatusInternalServerError)
		return
	}

	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		return
	}

	// acknowledge CONNECT
	_, _ = clientConn.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))

	tlsConfig := &tls.Config{
		GetCertificate: p.getCertificate,
	}

	tlsConn := tls.Server(clientConn, tlsConfig)
	if err := tlsConn.Handshake(); err != nil {
		clientConn.Close()
		return
	}

	go p.serveTLSConnection(tlsConn)
}

func (p *Proxy) serveTLSConnection(conn net.Conn) {

	defer conn.Close()

	server := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := rules.NewContext(w, r, p.transport)
			p.engine.Execute(ctx)
		}),
	}

	server.Serve(&singleUseListener{conn})
}

type singleUseListener struct {
	conn net.Conn
}

func (l *singleUseListener) Accept() (net.Conn, error) {
	if l.conn == nil {
		return nil, io.EOF
	}
	c := l.conn
	l.conn = nil
	return c, nil
}

func (l *singleUseListener) Close() error   { return nil }
func (l *singleUseListener) Addr() net.Addr { return l.conn.LocalAddr() }
