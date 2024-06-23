package impl

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptrace"
	"net/url"
	"strings"
)

// Args struct to hold arguments for HTTP requests
type Args struct {
	A string
}

// HTTPproc struct to hold the HTTP client and cookie jar
type HTTPproc struct {
	Jar http.CookieJar
}

// NewHTTPproc creates a new instance of HTTPproc with a cookie jar
func NewHTTPproc() *HTTPproc {
	jar, _ := cookiejar.New(nil)
	return &HTTPproc{Jar: jar}
}

// GET method performs a GET request
func (s *HTTPproc) GET(args *Args, reply *string) error {
	link := args.A
	method := "GET"
	var body io.Reader = nil
	u, _ := url.Parse(link)

	// NEW REQUEST
	req := &http.Request{
		Method:     method,
		URL:        u,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Host:       u.Host,
	}

	if body != nil {
		req.Body = io.NopCloser(body)
		switch v := body.(type) {
		case *bytes.Buffer:
			req.ContentLength = int64(v.Len())
			buf := v.Bytes()
			req.GetBody = func() (io.ReadCloser, error) {
				r := bytes.NewReader(buf)
				return io.NopCloser(r), nil
			}
		case *bytes.Reader:
			req.ContentLength = int64(v.Len())
			snapshot := *v
			req.GetBody = func() (io.ReadCloser, error) {
				r := snapshot
				return io.NopCloser(&r), nil
			}
		case *strings.Reader:
			req.ContentLength = int64(v.Len())
			snapshot := *v
			req.GetBody = func() (io.ReadCloser, error) {
				r := snapshot
				return io.NopCloser(&r), nil
			}
		}
	}

	if s.Jar != nil {
		for _, cookie := range s.Jar.Cookies(req.URL) {
			req.AddCookie(cookie)
		}
	}

	var (
		res  *http.Response
		erro error
	)

	// Sending the request
	res, erro = s.Do(req)
	if erro != nil {
		return erro
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	*reply = string(bodyBytes)
	return nil
}

// Do method sends the HTTP request and returns the response
func (s *HTTPproc) Do(req *http.Request) (*http.Response, error) {
	return s.Send(req, s.createTrace(req.Context()))
}

// Send method sends the HTTP request with tracing enabled
func (s *HTTPproc) Send(req *http.Request, trace *httptrace.ClientTrace) (*http.Response, error) {
	client := &http.Client{
		Jar: s.Jar,
	}

	// Assigning the client trace
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	// Sending the request
	res, err := client.Do(req)
	return res, err
}

// createTrace method creates a new ClientTrace for tracing the request
func (s *HTTPproc) createTrace(ctx context.Context) *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) {
			// Called when a connection is obtained
			println("Got Conn:", info.Reused)
		},
		DNSStart: func(info httptrace.DNSStartInfo) {
			// Called when DNS resolution begins
			println("DNS Start:", info.Host)
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			// Called when DNS resolution ends
			println("DNS Done:", info.Err)
		},
	}
}
