package main

import (
	"context"
	"github/Zholdaskali/go-grpc/pkg/api/example"
	"log"
	"math/rand"
	"net"
	"sync"

	"buf.build/go/protovalidate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func WithProtoValidator() grpc.UnaryServerInterceptor {
	return nil
}

func main() {
	validator, err := protovalidate.New()
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	service := &ExampleService{
		storage:   make(map[uint64]*Post, 1),
		validator: validator,
	}
	example.RegisterExampleServer(server, service)
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}
	reflection.Register(server)
	log.Println("gRPC server started: 8080")

	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

type Post struct {
	ID       uint64
	Title    string
	Content  string
	AuthorID string
}

type ExampleService struct {
	example.UnimplementedExampleServer

	validator protovalidate.Validator
	storage   map[uint64]*Post
	mx        sync.RWMutex
}

func (s *ExampleService) CreatePost(ctx context.Context, req *example.CreatePostRequest) (*example.CreatePostResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, err
	}

	id := rand.Uint64()
	post := &Post{
		Title:    req.GetTitle(),
		Content:  req.GetContent(),
		AuthorID: req.GetAuthorId(),
	}

	s.mx.Lock()

	s.storage[id] = post
	s.mx.Unlock()

	return &example.CreatePostResponse{
		PostId: id,
	}, nil
}

func (s *ExampleService) ListPosts(ctx context.Context, req *example.ListPostsRequest) (*example.ListPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPosts not implemented")
}
