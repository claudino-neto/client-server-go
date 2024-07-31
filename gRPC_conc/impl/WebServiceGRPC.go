package impl

import (
	"context"
	"fmt"
	pb "gRPC_conc/gen"
	"io"
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
func (s *HTTPproc) GET(ctx context.Context, req *pb.HttpRequest) (*pb.Response, error) {
	method := req.GetMethod()

	if method != "GET" {
		return nil, fmt.Errorf("unsupported HTTP method: %s", method)
	}

	filePath := "C:/Users/labou/OneDrive/Documentos/UFPE/CODING/codesGO/marmotas-3/gRPC_conc/index.html"

	// Open the file
	file, _ := os.Open(filePath)

	defer file.Close()

	// Read the file content
	bodyBytes, _ := io.ReadAll(file)

	// Return the file content in the response
	return &pb.Response{Body: string(bodyBytes)}, nil
}
