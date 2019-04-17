package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {

	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()
	image := "golang:alpine"
	reader, err := cli.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	io.Copy(os.Stdout, reader)
}
