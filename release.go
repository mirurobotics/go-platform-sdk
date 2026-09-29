// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package miruplatform

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/mirurobotics/go-platform-sdk/internal/apijson"
	"github.com/mirurobotics/go-platform-sdk/internal/apiquery"
	"github.com/mirurobotics/go-platform-sdk/internal/requestconfig"
	"github.com/mirurobotics/go-platform-sdk/option"
	"github.com/mirurobotics/go-platform-sdk/packages/param"
	"github.com/mirurobotics/go-platform-sdk/packages/respjson"
	"github.com/mirurobotics/go-platform-sdk/shared"
	"github.com/mirurobotics/go-platform-sdk/shared/constant"
)

// ReleaseService contains methods and other services that help with interacting
// with the miru API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReleaseService] method instead.
type ReleaseService struct {
	Options []option.RequestOption
}

// NewReleaseService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewReleaseService(opts ...option.RequestOption) (r ReleaseService) {
	r = ReleaseService{}
	r.Options = opts
	return
}

// Create a new release.
func (r *ReleaseService) New(ctx context.Context, params ReleaseNewParams, opts ...option.RequestOption) (res *Release, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "releases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve a release by ID.
func (r *ReleaseService) Get(ctx context.Context, releaseID string, query ReleaseGetParams, opts ...option.RequestOption) (res *Release, err error) {
	opts = slices.Concat(r.Options, opts)
	if releaseID == "" {
		err = errors.New("missing required release_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("releases/%s", releaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List releases.
func (r *ReleaseService) List(ctx context.Context, query ReleaseListParams, opts ...option.RequestOption) (res *ReleaseList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "releases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type Release struct {
	// ID of the release.
	ID string `json:"id" api:"required"`
	// Timestamp of when the release was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The ID of the git commit associated with this release.
	GitCommitID string `json:"git_commit_id" api:"required"`
	// The object type, which is always `release`.
	Object constant.Release `json:"object" default:"release"`
	// An operating system family, from the agent's build-time OS vocabulary. Shared
	// across resources (devices, config schemas, file rules, releases). This is the
	// machine-readable OS kind, distinct from a human-readable version string.
	//
	// - `linux`
	// - `windows`
	//
	// Any of "linux", "windows".
	Os ReleaseOs `json:"os" api:"required"`
	// Timestamp of when the release was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The version of the release.
	Version string `json:"version" api:"required"`
	// Expand the config schemas using 'expand=config_schemas' in the query string.
	ConfigSchemas []ConfigSchema `json:"config_schemas"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		GitCommitID   respjson.Field
		Object        respjson.Field
		Os            respjson.Field
		UpdatedAt     respjson.Field
		Version       respjson.Field
		ConfigSchemas respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Release) RawJSON() string { return r.JSON.raw }
func (r *Release) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An operating system family, from the agent's build-time OS vocabulary. Shared
// across resources (devices, config schemas, file rules, releases). This is the
// machine-readable OS kind, distinct from a human-readable version string.
//
// - `linux`
// - `windows`
type ReleaseOs string

const (
	ReleaseOsLinux   ReleaseOs = "linux"
	ReleaseOsWindows ReleaseOs = "windows"
)

type ReleaseList struct {
	// The list of releases.
	Data []Release `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.PaginatedList
}

// Returns the unmodified JSON received from the API
func (r ReleaseList) RawJSON() string { return r.JSON.raw }
func (r *ReleaseList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReleaseNewParams struct {
	// The IDs of the config schemas included in the release.
	ConfigSchemaIDs []string `json:"config_schema_ids,omitzero" api:"required"`
	// The version of the release.
	Version string `json:"version" api:"required"`
	// Fields to expand on the release resource.
	//
	// Any of "config_schemas".
	Expand []string `query:"expand,omitzero" json:"-"`
	// A reference to a git commit. At least one of `id` or `sha` must be provided.
	// When both are provided, `id` takes precedence and `sha` is ignored.
	GitCommitRef GitCommitRefParam `json:"git_commit_ref,omitzero"`
	paramObj
}

func (r ReleaseNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ReleaseNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReleaseNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ReleaseNewParams]'s query parameters as `url.Values`.
func (r ReleaseNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ReleaseGetParams struct {
	// Fields to expand on the release resource.
	//
	// Any of "config_schemas".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReleaseGetParams]'s query parameters as `url.Values`.
func (r ReleaseGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ReleaseListParams struct {
	// The maximum number of items to return. A limit of 15 with an offset of 0 returns
	// items 1-15.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The offset of the items to return. An offset of 10 with a limit of 10 returns
	// items 11-20.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// The release IDs to filter by.
	ID []string `query:"id,omitzero" json:"-"`
	// Fields to expand on each release in the list.
	//
	// Any of "total_count", "config_schemas".
	Expand []string `query:"expand,omitzero" json:"-"`
	// Sort order for the release results.
	//
	// Any of "id:asc", "id:desc", "created_at:desc", "created_at:asc".
	OrderBy []string `query:"order_by,omitzero" json:"-"`
	// The release versions to filter by.
	Version []string `query:"version,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReleaseListParams]'s query parameters as `url.Values`.
func (r ReleaseListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
