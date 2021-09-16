package main

import (
        "log"  
        "os"  

        "cloud.google.com/go/storage"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
        "google.golang.org/api/iterator"
)

func main() {
	ctx := context.Background()

	credentials, err := google.FindDefaultCredentials(ctx,compute.ComputeScope)
	if err != nil {
	    log.Fatal("error getting ADC: ", err)
	}

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
	    projectID = credentials.ProjectID
	}

	if projectID == "" {
		log.Fatal("Unable to determine Project ID")
	}

	storeageClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Unable to acquire storage Client: %v", err)
	}

	it := storeageClient.Buckets(ctx, projectID)
	for {
		bucketAttrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Unable to acquire storage Client: %v", err)
		}
		log.Printf(bucketAttrs.Name)
	}
}
