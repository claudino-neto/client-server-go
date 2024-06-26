package impl

import (
	"context"
	pb "gRPC/gen"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptrace"
	"net/url"
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

// GET method performs a GET request
func (s *HTTPproc) GET(ctx context.Context, args *pb.Request) (*pb.Response, error) {
	link := args.Link
	method := "GET"
	// var body io.Reader = nil
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

	if s.Jar != nil {
		for _, cookie := range s.Jar.Cookies(req.URL) {
			req.AddCookie(cookie)
		}
	}
	var (
		res *http.Response
	)

	// Sending the request
	res, _ = s.Send(req, s.createTrace(req.Context()))

	defer res.Body.Close()

	bodyBytes, _ := io.ReadAll(res.Body)

	reply := string(bodyBytes)
	return &pb.Response{Body: reply}, nil
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

// if body != nil {
// 	req.Body = io.NopCloser(body)
// 	switch v := body.(type) {
// 	case *bytes.Buffer:
// 		req.ContentLength = int64(v.Len())
// 		buf := v.Bytes()
// 		req.GetBody = func() (io.ReadCloser, error) {
// 			r := bytes.NewReader(buf)
// 			return io.NopCloser(r), nil
// 		}
// 	case *bytes.Reader:
// 		req.ContentLength = int64(v.Len())
// 		snapshot := *v
// 		req.GetBody = func() (io.ReadCloser, error) {
// 			r := snapshot
// 			return io.NopCloser(&r), nil
// 		}
// 	case *strings.Reader:
// 		req.ContentLength = int64(v.Len())
// 		snapshot := *v
// 		req.GetBody = func() (io.ReadCloser, error) {
// 			r := snapshot
// 			return io.NopCloser(&r), nil
// 		}
// 	}
// }
