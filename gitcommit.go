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

// GitCommitService contains methods and other services that help with interacting
// with the miru API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGitCommitService] method instead.
type GitCommitService struct {
	Options []option.RequestOption
}

// NewGitCommitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGitCommitService(opts ...option.RequestOption) (r GitCommitService) {
	r = GitCommitService{}
	r.Options = opts
	return
}

// Register a git commit.
func (r *GitCommitService) New(ctx context.Context, body GitCommitNewParams, opts ...option.RequestOption) (res *GitCommit, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "git_commits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a git commit by ID.
func (r *GitCommitService) Get(ctx context.Context, gitCommitID string, opts ...option.RequestOption) (res *GitCommit, err error) {
	opts = slices.Concat(r.Options, opts)
	if gitCommitID == "" {
		err = errors.New("missing required git_commit_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("git_commits/%s", gitCommitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List git commits.
func (r *GitCommitService) List(ctx context.Context, query GitCommitListParams, opts ...option.RequestOption) (res *GitCommitList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "git_commits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type GitCommit struct {
	// ID of the git commit.
	ID string `json:"id" api:"required"`
	// The URL of the git commit.
	CommitURL string `json:"commit_url" api:"required"`
	// Timestamp of when the git commit was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The commit message.
	Message string `json:"message" api:"required"`
	// The object type, which is always `git_commit`.
	Object constant.GitCommit `json:"object" default:"git_commit"`
	// The name of the git repository.
	RepositoryName string `json:"repository_name" api:"required"`
	// The owner of the git repository.
	RepositoryOwner string `json:"repository_owner" api:"required"`
	// Any of "github", "gitlab", "bitbucket".
	RepositoryType GitCommitRepositoryType `json:"repository_type" api:"required"`
	// The URL of the git repository.
	RepositoryURL string `json:"repository_url" api:"required"`
	// The SHA hash of the git commit.
	Sha string `json:"sha" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CommitURL       respjson.Field
		CreatedAt       respjson.Field
		Message         respjson.Field
		Object          respjson.Field
		RepositoryName  respjson.Field
		RepositoryOwner respjson.Field
		RepositoryType  respjson.Field
		RepositoryURL   respjson.Field
		Sha             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GitCommit) RawJSON() string { return r.JSON.raw }
func (r *GitCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GitCommitRepositoryType string

const (
	GitCommitRepositoryTypeGitHub    GitCommitRepositoryType = "github"
	GitCommitRepositoryTypeGitlab    GitCommitRepositoryType = "gitlab"
	GitCommitRepositoryTypeBitbucket GitCommitRepositoryType = "bitbucket"
)

type GitCommitList struct {
	// The list of git commits.
	Data []GitCommit `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.PaginatedList
}

// Returns the unmodified JSON received from the API
func (r GitCommitList) RawJSON() string { return r.JSON.raw }
func (r *GitCommitList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A reference to a git commit. At least one of `id` or `sha` must be provided.
// When both are provided, `id` takes precedence and `sha` is ignored.
type GitCommitRefParam struct {
	// ID of the git commit. Takes precedence over `sha` when both are provided.
	ID param.Opt[string] `json:"id,omitzero"`
	// The SHA hash of the git commit. Used only when `id` is not provided.
	Sha param.Opt[string] `json:"sha,omitzero"`
	paramObj
}

func (r GitCommitRefParam) MarshalJSON() (data []byte, err error) {
	type shadow GitCommitRefParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GitCommitRefParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GitCommitNewParams struct {
	// The commit message.
	Message string `json:"message" api:"required"`
	// The URL of the git repository origin.
	Origin string `json:"origin" api:"required"`
	// The SHA hash of the git commit.
	Sha string `json:"sha" api:"required"`
	paramObj
}

func (r GitCommitNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GitCommitNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GitCommitNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GitCommitListParams struct {
	// The maximum number of items to return. A limit of 15 with an offset of 0 returns
	// items 1-15.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The offset of the items to return. An offset of 10 with a limit of 10 returns
	// items 11-20.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// The git commit IDs to filter by.
	ID []string `query:"id,omitzero" json:"-"`
	// Fields to expand on each git commit in the list.
	//
	// Any of "total_count".
	Expand []string `query:"expand,omitzero" json:"-"`
	// Sort order for the git commit results.
	//
	// Any of "id:asc", "id:desc", "created_at:desc", "created_at:asc".
	OrderBy []string `query:"order_by,omitzero" json:"-"`
	// The git commit SHAs to filter by.
	Sha []string `query:"sha,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GitCommitListParams]'s query parameters as `url.Values`.
func (r GitCommitListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
