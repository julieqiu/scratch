package main

import (
	"context"
	"log"

	edgecontainer "cloud.google.com/go/edgecontainer/apiv1"
	edgecontainerpb "cloud.google.com/go/edgecontainer/apiv1/edgecontainerpb"
	"google.golang.org/api/iterator"
)

func main() {
	ctx := context.Background()

	c, err := edgecontainer.NewClient(ctx)
	if err != nil {
	}
	defer c.Close()

	req := &edgecontainerpb.ListClustersRequest{
		Parent: "projects/theviewfromjq/locations/us-central1",
	}

	it := c.ListClusters(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		_ = resp
	}
}
