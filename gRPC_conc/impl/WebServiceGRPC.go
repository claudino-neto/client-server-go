package impl

import (
	"context"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptrace"
	"net/url"
	"sync"

	pb "gRPC/gen"
)

// Args struct to hold arguments for HTTP requests
type Args struct {
	A string
}

// HTTPproc struct to hold the HTTP client and cookie jar
type HTTPproc struct {
	pb.UnimplementedHTTPServiceServer
	Jar http.CookieJar
}

// NewHTTPproc creates a new instance of HTTPproc with a cookie jar
func NewHTTPproc() *HTTPproc {
	jar, _ := cookiejar.New(nil)
	return &HTTPproc{Jar: jar}
}

// GET method performs concurrent GET requests
func (s *HTTPproc) GET(ctx context.Context, args *pb.Request) (*pb.Response, error) {
	link := args.Link
	numRequests := 20
	results := make(chan string, numRequests)
	var wg sync.WaitGroup

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go s.performRequest(ctx, link, results, &wg)
	}

	wg.Wait()
	close(results)

	// Concatenate results
	var reply string
	for result := range results {
		reply += result + "\n"
	}

	return &pb.Response{Body: reply}, nil
}

// performRequest performs a single GET request and sends the result to a channel
func (s *HTTPproc) performRequest(ctx context.Context, link string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	u, err := url.Parse(link)
	if err != nil {
		results <- "Error parsing URL: " + err.Error()
		return
	}

	req := s.createRequest("GET", u)
	res, err := s.Send(req, s.createTrace(ctx))
	if err != nil {
		results <- "Error sending request: " + err.Error()
		return
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		results <- "Error reading response body: " + err.Error()
		return
	}

	results <- string(bodyBytes)
}

// createRequest creates a new HTTP request
func (s *HTTPproc) createRequest(method string, u *url.URL) *http.Request {
	return &http.Request{
		Method:     method,
		URL:        u,
		Proto:      "HTTP/2",
		ProtoMajor: 2,
		ProtoMinor: 0,
		Header:     make(http.Header),
		Host:       u.Host,
	}
}

// Send method sends the HTTP request with tracing enabled
func (s *HTTPproc) Send(req *http.Request, trace *httptrace.ClientTrace) (*http.Response, error) {
	client := &http.Client{
		Jar: s.Jar,
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	return client.Do(req)
}

// createTrace method creates a new ClientTrace for tracing the request
func (s *HTTPproc) createTrace(ctx context.Context) *httptrace.ClientTrace {
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
