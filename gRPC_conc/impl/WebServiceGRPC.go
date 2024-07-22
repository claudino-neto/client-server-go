package impl

import (
	"context"
	pb "gRPC/gen"
	"io"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"sync"
)

// HTTPproc struct to hold the HTTP client and cookie jar
type HTTPproc struct {
	pb.UnimplementedHTTPServiceServer
}

// NewHTTPproc creates a new instance of HTTPproc
func NewHTTPproc() *HTTPproc {
	return &HTTPproc{}
}

// GET method performs a GET request
func (s *HTTPproc) GET(ctx context.Context, args *pb.Request) (*pb.Response, error) {
	var (
		client      = &http.Client{}
		link        = args.Link
		u, _        = url.Parse(link)
		reqContext  context.Context
		clientTrace *httptrace.ClientTrace
		wg          sync.WaitGroup
	)

	//Creating the request
	req := &http.Request{
		Method:     "GET",
		URL:        u,
		Proto:      "HTTP/2",
		ProtoMajor: 2,
		ProtoMinor: 0,
		Header:     make(http.Header),
		Host:       u.Host,
	}

	// Creating Request Context and Client Trace concurrently
	wg.Add(2)
	go func() {
		defer wg.Done()
		reqContext = req.Context()
	}()
	go func() {
		defer wg.Done()
		clientTrace = s.createTrace()
	}()
	wg.Wait()

	req = req.WithContext(httptrace.WithClientTrace(reqContext, clientTrace))
	res, _ := client.Do(req)

	defer res.Body.Close()
	bodyBytes, _ := io.ReadAll(res.Body)
	reply := string(bodyBytes)
	return &pb.Response{Body: reply}, nil
}

// createTrace method creates a new ClientTrace for tracing the request
func (s *HTTPproc) createTrace() *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) {
			// Called when a connection is obtained
		},
		DNSStart: func(info httptrace.DNSStartInfo) {
			// Called when DNS resolution begins
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			// Called when DNS resolution ends
		},
	}
}
