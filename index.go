package main

import (
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudkms/v1"
	"google.golang.org/api/iterator"
)

/*
import(
	"deck"
)
*/

func main() {
	// d := Deck{}
	// d.reset()
	// deal := d.deal(4)
	// d.reset()
	// d.shuffle()
	// d.shuffle()
	// d.save("first.txt")
	// d.load("first.txt")
	// d.print()
	// fmt.Println("deck:")
	// d.print()
	// fmt.Println("deal cards:")
	// deal.save("first.txt")
	// d.print()
}

// implicit uses Application Default Credentials to authenticate.
func implicit() {
	ctx := context.Background()

	// For API packages whose import path is starting with "cloud.google.com/go",
	// such as cloud.google.com/go/storage in this case, if there are no credentials
	// provided, the client library will look for credentials in the environment.
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	it := storageClient.Buckets(ctx, "project-id")
	for {
		bucketAttrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(bucketAttrs.Name)
	}

	// For packages whose import path is starting with "google.golang.org/api",
	// such as google.golang.org/api/cloudkms/v1, use the
	// golang.org/x/oauth2/google package as shown below.
	oauthClient, err := google.DefaultClient(ctx, cloudkms.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	kmsService, err := cloudkms.New(oauthClient)
	if err != nil {
		log.Fatal(err)
	}

	_ = kmsService

	fmt.Println(kmsService)
}
