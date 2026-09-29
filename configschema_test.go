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

func TestConfigSchemaNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ConfigSchemas.New(context.TODO(), miruplatform.ConfigSchemaNewParams{
		ConfigTypeRef: miruplatform.ConfigSchemaNewParamsConfigTypeRef{
			ID:   miruplatform.String("cfg_123"),
			Slug: miruplatform.String("my-config-type"),
		},
		Documents: []miruplatform.ConfigSchemaNewParamsDocument{{
			Data: "data",
			Name: "motion-control.json",
		}},
		Format:   miruplatform.ConfigSchemaNewParamsFormatJson,
		Language: miruplatform.SchemaLanguageJsonschema,
		Expand:   []string{"documents"},
		GitCommit: miruplatform.ConfigSchemaNewParamsGitCommit{
			CommitRef: miruplatform.GitCommitRefParam{
				ID:  miruplatform.String("git_cmt_123"),
				Sha: miruplatform.String("1a2b3c4d..."),
			},
			SchemaFilepaths: []string{"path/to/config/schema/robot1.cue", "path/to/config/schema/robot2.cue"},
		},
		InstanceFormat: miruplatform.ConfigSchemaNewParamsInstanceFormatJson,
		InstanceSlots: []miruplatform.InstanceSlotParam{{
			Filepath:    "/var/local/forge/configuration/robot_drivers/controller_2/topic_list.yaml",
			Key:         "controller_2",
			Name:        "Controller 2",
			Required:    true,
			Description: miruplatform.String("Topic list for the second motor controller."),
		}},
	})
	if err != nil {
		var apierr *miruplatform.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConfigSchemaGetWithOptionalParams(t *testing.T) {
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
	_, err := client.ConfigSchemas.Get(
		context.TODO(),
		"cfg_sch_123",
		miruplatform.ConfigSchemaGetParams{
			Expand: []string{"documents"},
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

func TestConfigSchemaListWithOptionalParams(t *testing.T) {
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
	_, err := client.ConfigSchemas.List(context.TODO(), miruplatform.ConfigSchemaListParams{
		ID:           []string{"cfg_sch_123"},
		ConfigTypeID: []string{"cfg_typ_123"},
		Expand:       []string{"documents"},
		Limit:        miruplatform.Int(10),
		Offset:       miruplatform.Int(0),
		OrderBy:      []string{"created_at:desc"},
	})
	if err != nil {
		var apierr *miruplatform.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
