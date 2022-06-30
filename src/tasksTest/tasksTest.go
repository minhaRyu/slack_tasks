package tasksTest

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/cloudkms/v1"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/tasks/v1"
)

func Insert() {
	ctx := context.Background()
	tasksService, err := tasks.NewService(ctx, option.WithCredentialsFile("../task-sample-354616-de943d9002f1.json"))
	// tasksService, err := tasks.NewService(ctx, option.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))

	log.Println(tasksService, err)

	// insert 테스트
	// tempTaskList := tasks.TaskList{Title: "Minha task list test"}
	// tasksService.Tasklists.Insert(&tempTaskList).Do()

	list, _ := tasksService.Tasklists.List().MaxResults(30).Do()

	log.Println(list.Items[0])
}

// os 환경변수에 따른 기본 인증
// https://github.com/GoogleCloudPlatform/golang-samples/blob/main/auth/snippets.go
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
	defer storageClient.Close()

	it := storageClient.Buckets(ctx, os.Getenv("GOOGLE_STORAGE_BUCKET_PROJECTID"))
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
	// such as google.golang.org/api/cloudkms/v1, use NewService to create the client.
	kmsService, err := cloudkms.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}

	_ = kmsService
}
