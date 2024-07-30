package impl

import (
	"context"
	pb "gRPC/gen"
	"io"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"sync"
	"os"
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
func (s *HTTPproc) GET(ctx context.Context, req *&pb.HttpRequest) (*pb.Response, error) {
		method := req.GetMethod()
		url := req.GetUrl()
		header := req.GetHeader()
		
		if method != "GET" {
			return nil, fmt.Errorf("unsupported HTTP method: %s", method)
		}
		// Define the path to the local index.html file
		// You can adjust the path as needed
		filePath := fmt.Sprintf("marmotas/%s", url)
	
		// Open the file
		file, err := os.Open(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
		}
		defer file.Close()
	
		// Read the file content
		bodyBytes, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
		}
	
		// Return the file content in the response
		return &pb.Response{Body: string(bodyBytes)}, nil
	}

