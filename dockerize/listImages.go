package main

import (
	"context"
	"log"
	"fmt"

	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
)

const VERSION  = "1.39"

func main() {
	cli, err := client.NewClientWithOpts(client.WithVersion(VERSION))
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()
	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	for _, image := range images {
		fmt.Printf("[%s] Labels: %v\n", image.ID, image.RepoTags)
	}
}
