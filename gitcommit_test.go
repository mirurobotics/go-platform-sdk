// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package miruplatform_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/mirurobotics/go-platform-sdk"
	"github.com/mirurobotics/go-platform-sdk/internal/testutil"
	"github.com/mirurobotics/go-platform-sdk/option"
)

func TestGitCommitNew(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := miruplatform.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.GitCommits.New(context.TODO(), miruplatform.GitCommitNewParams{
		Message: "Update config schema for robot1",
		Origin:  "https://github.com/robot1/robot1.git",
		Sha:     "a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0",
	})
	if err != nil {
		var apierr *miruplatform.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGitCommitGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := miruplatform.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.GitCommits.Get(context.TODO(), "git_commit_123")
	if err != nil {
		var apierr *miruplatform.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGitCommitListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := miruplatform.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.GitCommits.List(context.TODO(), miruplatform.GitCommitListParams{
		ID:      []string{"git_cmt_123"},
		Expand:  []string{"total_count"},
		Limit:   miruplatform.Int(10),
		Offset:  miruplatform.Int(0),
		OrderBy: []string{"created_at:desc"},
		Sha:     []string{"a1b2c3d4e5f6"},
	})
	if err != nil {
		var apierr *miruplatform.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
