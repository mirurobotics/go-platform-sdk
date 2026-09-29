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

func TestReleaseNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Releases.New(context.TODO(), miruplatform.ReleaseNewParams{
		ConfigSchemaIDs: []string{"cfg_sch_123"},
		Version:         "v1.0.0",
		Expand:          []string{"config_schemas"},
		GitCommitRef: miruplatform.GitCommitRefParam{
			ID:  miruplatform.String("git_cmt_123"),
			Sha: miruplatform.String("1a2b3c4d..."),
		},
	})
	if err != nil {
		var apierr *miruplatform.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReleaseGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Releases.Get(
		context.TODO(),
		"rls_123",
		miruplatform.ReleaseGetParams{
			Expand: []string{"config_schemas"},
		},
	)
	if err != nil {
		var apierr *miruplatform.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReleaseListWithOptionalParams(t *testing.T) {
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
	_, err := client.Releases.List(context.TODO(), miruplatform.ReleaseListParams{
		ID:      []string{"rls_123"},
		Expand:  []string{"config_schemas"},
		Limit:   miruplatform.Int(10),
		Offset:  miruplatform.Int(0),
		OrderBy: []string{"created_at:desc"},
		Version: []string{"v1.0.0"},
	})
	if err != nil {
		var apierr *miruplatform.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
