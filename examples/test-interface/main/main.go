package main


import (
	"log"
	"fmt"
	"google.golang.org/grpc"
	//pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"github.com/jameszhangcn/post"
)

const (
	address = "localhost:9027"
	defaultNmae = "world"
)
func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println(conn)
	defer conn.Close()
	fmt.Println("Test interface")
	
	svc := post.NewService(conn)
	posts, err := svc.ListPosts()
	if err != nil {
		panic(err)
	}
	fmt.Println(posts)
} 