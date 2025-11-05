package main

import (
	"context"
	"github/Zholdaskali/go-grpc/pkg/api/example"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal()
	}
	client := example.NewExampleClient(conn)

	res, err := client.CreatePost(context.Background(), &example.CreatePostRequest{
		Title:    "asdasd",
		AuthorId: "1",
		Content:  "Musina Azhar",
	})

	if err != nil {
		log.Fatal()
	}

	log.Println(res.GetPostId())

	bytes, err := protojson.Marshal(res)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bytes))

}
