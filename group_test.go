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

func TestGroupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Groups.New(context.TODO(), miruplatform.GroupNewParams{
		Name:     "Robots",
		Expand:   []string{"parent"},
		ParentID: miruplatform.String("grp_456"),
	})
	if err != nil {
		var apierr *miruplatform.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGroupGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Groups.Get(
		context.TODO(),
		"grp_123",
		miruplatform.GroupGetParams{
			Expand: []string{"parent"},
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

func TestGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Groups.Update(
		context.TODO(),
		"grp_123",
		miruplatform.GroupUpdateParams{
			Expand: []string{"parent"},
			Name:   miruplatform.String("Robots"),
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

func TestGroupListWithOptionalParams(t *testing.T) {
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
	_, err := client.Groups.List(context.TODO(), miruplatform.GroupListParams{
		ID:       []string{"grp_123"},
		Expand:   []string{"parent"},
		Limit:    miruplatform.Int(10),
		Name:     []string{"production"},
		Offset:   miruplatform.Int(0),
		OrderBy:  []string{"created_at:desc"},
		ParentID: []string{"grp_456"},
	})
	if err != nil {
		var apierr *miruplatform.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGroupMoveWithOptionalParams(t *testing.T) {
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
	_, err := client.Groups.Move(
		context.TODO(),
		"grp_123",
		miruplatform.GroupMoveParams{
			ParentID: miruplatform.String("grp_456"),
			Expand:   []string{"parent"},
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
