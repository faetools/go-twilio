package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/google/go-github/v43/github"
)

const org, repo, branch, folder = "twilio", "twilio-oai", "main", "spec/yaml"

var reName = regexp.MustCompile("twilio_(.+)_(v[^_]+).yaml")

//go:generate go run gen-openapi.go

func main() {
	ctx := context.Background()
	cl := github.NewClient(http.DefaultClient)

	_, contents, _, err := cl.Repositories.GetContents(ctx, org, repo, folder, &github.RepositoryContentGetOptions{
		Ref: branch,
	})
	if err != nil {
		log.Fatal(err)
	}

	apis := map[string]api{}
	for _, content := range contents {
		match := reName.FindStringSubmatch(content.GetName())

		if len(match) != 3 {
			log.Fatalf("name %q does not match", content.GetName())
		}

		name, version := match[1], match[2]

		if apis[name].version > version {
			continue // ignore lower versions
		}

		apis[name] = api{
			name:              name,
			version:           version,
			RepositoryContent: content,
		}
	}

	for _, api := range apis {
		dir := filepath.Join("api", api.name)

		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, api.GetDownloadURL(), nil)
		if err != nil {
			log.Fatal(err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		path := filepath.Join(dir, "openapi.yaml")

		f, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0o644)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := io.Copy(f, res.Body); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Done.")
}

type api struct {
	name    string
	version string
	*github.RepositoryContent
}
