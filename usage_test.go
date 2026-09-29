// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package miruplatform_test

import (
	"context"
	"os"
	"testing"

	"github.com/mirurobotics/go-platform-sdk"
	"github.com/mirurobotics/go-platform-sdk/internal/testutil"
	"github.com/mirurobotics/go-platform-sdk/option"
)

func TestUsage(t *testing.T) {
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
	deploymentList, err := client.Deployments.List(context.TODO(), miruplatform.DeploymentListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", deploymentList)
}
