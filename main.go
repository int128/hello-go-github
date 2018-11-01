package main

import (
	"context"
	"log"

	"github.com/google/go-github/v18/github"
)

func main() {
	ctx := context.Background()
	c := github.NewClient(nil)

	// Content API
	fc, _, _, err := c.Repositories.GetContents(ctx, "int128", "latest-gradle-wrapper", "gradle/wrapper/gradle-wrapper.properties", nil)
	if err != nil {
		log.Fatalf("Could not get content from GitHub: %s", err)
	}
	content, err := fc.GetContent()
	if err != nil {
		log.Fatalf("Could not decode content: %s", err)
	}
	log.Printf("Content:\n%s", content)

	// Pull Requests API
	pulls, _, err := c.PullRequests.List(ctx, "int128", "gradleupdate",
		&github.PullRequestListOptions{
			State: "close",
		})
	if err != nil {
		log.Fatalf("Could not get pull requests from GitHub: %s", err)
	}
	for _, pull := range pulls {
		log.Printf("#%d %s", pull.GetNumber(), pull.GetTitle())
	}

	// Git Data API
	blobContent := "MjAxOOW5tCAxMeaciCAx5pelIOacqOabnOaXpSAxMeaZgjA55YiGMTbnp5IgSlNUCg=="
	blobEncoding := "base64"
	blob, _, err := c.Git.CreateBlob(ctx, "int128", "gradleupdate",
		&github.Blob{
			Content:  &blobContent,
			Encoding: &blobEncoding,
		})
	if err != nil {
		log.Fatalf("Could not create blob on GitHub: %s", err)
	}
	log.Printf("Blob SHA %s", blob.GetSHA())
}
