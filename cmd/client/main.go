package main

import (
	"context"
	"fmt"
	"github/Zholdaskali/go-grpc/pkg/api/example"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal()
	}

	client := example.NewExampleClient(conn)

	ctx := context.Background()
	var header, trailer metadata.MD
	cctx := metadata.NewOutgoingContext(ctx, metadata.Pairs("client-header-key", "val"))
	res, err := client.CreatePost(cctx, &example.CreatePostRequest{
		Title:    "asdasd",
		AuthorId: "1",
		Content:  "Musina Azhar",
	}, grpc.Header(&header), grpc.Trailer(&trailer))

	if err != nil {
		switch status.Code(err) {
		case codes.InvalidArgument:
			log.Println("некорректный запрос")
		default:
			log.Fatal(err)
		}

		if st, ok := status.FromError(err); ok {
			log.Print("code", st.Code(), "detail", st.Details(), "message", st.Message())
		} else {
			log.Print("not grpc")
		}
	}

	fmt.Println("header:", header)
	fmt.Println("trailer", trailer)

	log.Println(res.GetPostId())

	bytes, err := protojson.Marshal(res)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bytes))

}
